package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)


// helper function for various tests
func getBigRandomString() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.FormatInt(rand.Int63(), 10)
}

/* Helper function tests */

func Test_encodeAndDecode(t *testing.T) {
	tables := []struct {
		x int
		s string
	}{
		{0,"a"},
		{1,"b"},
		{9, "j"},
		{61, "9"},
		{62, "ba"},
		{63, "bb"},
		{123, "b9"},
		{124, "ca"},
		{125, "cb"},
		{162,"cM"},
	}

	for _, table := range tables {
		tryS := encodeBase62(table.x)
		tryX := decodeBase62(table.s)
		if tryS != table.s {
			t.Errorf("Encoding of %d was incorrect, got: %s, want: %s.", table.x, tryS, table.s)
		}
		if tryX != table.x {
			t.Errorf("Decoding of %s was incorrect, got: %d, want: %d.", table.s, tryX, table.x)
		}
	}
}

func Test_checkUrl(t *testing.T) {
	tables := []struct {
		url string
		valid bool
	}{
		{"https://www.google.com/",true},
		{"www.google.com/",false},
		{"google.com", false},
		{"https://www.google.com/search?source=hp&ei=0uiRW6OaIKWs_Qb9o63AAQ&q=gfdgfdgd&btnK=Google+Search&oq=gfdgfdgd&gs_l=psy-ab.3..0j0i10l5.146389.147136..147407...2.0..1.115.809.8j2....1..0....1..gws-wiz.....6..35i39j0i131.jJUy1nynSic", true},
		{"https://duckduckgo.com/", true},
		{"https://duckduckgo.com/?q=gfdgfgd&t=hi&atb=v41-2&ia=videos", true},

	}

	for _, table := range tables {
		tryValid := checkUrl(table.url)
		if tryValid != table.valid {
			t.Errorf("Checking of url %s was incorrect, got: %t, want: %t.", table.url, tryValid, table.valid)
		}
	}
}

func Test_reverse(t *testing.T) {
	tables := []struct {
		s string
		r string
	}{
		{"racecar","racecar"},
		{"word","drow"},
		{"123abc", "cba321"},
		{"boot", "toob"},
	}

	for _, table := range tables {
		tryR := reverse(table.s)
		if tryR != table.r {
			t.Errorf("Reversing of string %s was incorrect, got: %s, want: %s.", table.s, tryR, table.r)
		}
	}
}

/* Database tests */

func Test_DBQueries(t *testing.T) {

	db, err := initializeDB()
	if err != nil {
		t.Errorf("error connecting to db: %s.", err)
	}

	url := "https://randomurl.com/" + getBigRandomString()

	// test that addLongUrlGetId does not produce an error
	id, err := addLongUrlGetId(db, url)
	if err != nil {
		t.Errorf("calling addLongUrlGetId for url %s had the following error: %s.", url, err)
		return
	}

	// test that getLongUrlFromId does not produce an error and that the url returned is the same as the one inserted
	tryUrl, err := getLongUrlFromId(db, id)
	if err != nil {
		t.Errorf("calling getLongUrlFromId for id %d had the following error: %s.", id, err)
	}
	if tryUrl != url {
		t.Errorf("calling getLongUrlFromId for id %d, wrong url, got: %s, want: %s.", id, tryUrl, url)
	}

	// test that getIdOfLongUrl does not produce an error and that the id returned is the same as where the url was inserted
	tryId, err := getIdOfLongUrl(db, url)
	if err != nil {
		t.Errorf("calling getIdOfLongUrl for url %s had the following error: %s.", url, err)
	}
	if tryId != id {
		t.Errorf("calling getIdOfLongUrl for url %s, wrong id, got: %d, want: %d.", url, tryId, id)
	}

	// test that the id of the next url inserted is consecutive to the last one
	url2 := "https://randomurl2.com/" + getBigRandomString()
	id2, err := addLongUrlGetId(db, url)
	if err != nil {
		t.Errorf("calling addLongUrlGetId for url %s had the following error: %s.", url, err)
		return
	}
	if id2 != id + 1 {
		t.Errorf("calling addLongUrlGetId for url %s, wrong id, got: %d, want: %d.", url2, id2, id+1)
	}
	db.Close()

}

/* Endpoint request and response testing */

// page representation
type page struct {
	method string
	url string
	desc string
	inPhrases []string
	outPhrases []string
	status int
	values url.Values
}

// test for any generic page in the web app
func genericPageTest(t *testing.T, p page ) {

	req, err := http.NewRequest(p.method, p.url, strings.NewReader(p.values.Encode()))
	if err != nil {
		t.Errorf("Request for %s had the following error: %s.", p.desc, err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res := httptest.NewRecorder()
	initializeRouter().ServeHTTP(res, req)
	if res.Code != p.status {
		t.Errorf("Wrong status code, got: %d, want: %d", res.Code, p.status)
	}
	//defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Reading body response had the following error: %s.", err)
	}
	for _, phrase := range p.inPhrases {
		if !strings.Contains(string(body), phrase) {
			t.Errorf("Expected phrase '%s' on %s", phrase, p.desc)
		}
	}
	for _, phrase := range p.outPhrases {
		if strings.Contains(string(body), phrase) {
			t.Errorf("Unexpected phrase '%s' on %s", phrase, p.desc)
		}
	}

}

func Test_homePage(t *testing.T) {
	inPhrases := []string{"Get a shortened url:", "Enter the url...", "Submit"}
	outPhrases := []string{"Original URL:", "Shortened URL:", "localhost", INVALID_URL, GEN_ERROR, BAD_SHORT_URL}
	p := page{"GET", "/", "home page", inPhrases, outPhrases, http.StatusOK, url.Values{}}
	genericPageTest(t, p)
}

func Test_shortUrlPage(t *testing.T) {
	longurl := "https://longurl.com/" + getBigRandomString()
	inPhrases := []string{"Get a shortened url:", "Enter the url...", "Submit", "Original URL:", "Shortened URL:", "localhost"}
	outPhrases := []string{INVALID_URL, GEN_ERROR, BAD_SHORT_URL}
	p := page{"POST", "/", "short url page", inPhrases, outPhrases, http.StatusOK, url.Values{"longurl": {longurl}}}
	genericPageTest(t, p)
}

func Test_invalidUrlPage(t *testing.T) {
	longurl := "invalid.url"
	inPhrases := []string{"Get a shortened url:", "Enter the url...", "Submit", INVALID_URL}
	outPhrases := []string{"Original URL:", "Shortened URL:", "localhost", GEN_ERROR, BAD_SHORT_URL}
	p := page{"POST", "/", "invalid url page", inPhrases, outPhrases, http.StatusOK, url.Values{"longurl": {longurl}}}
	genericPageTest(t, p)
}

func Test_shortUrlRedirect(t *testing.T) {
	shorturl := "e"
	inPhrases := []string{"github", "go-sql-driver", "mysql"}
	outPhrases := []string{"Enter the url...", "Get a shortened url:", "Original URL:", "Shortened URL:", "localhost", GEN_ERROR, BAD_SHORT_URL, INVALID_URL}
	p := page{"GET", SHORT_URL_ROUTE + shorturl, "short url redirect page", inPhrases, outPhrases, http.StatusSeeOther, url.Values{}}
	genericPageTest(t, p)
}

func test_invalidShortUrlRedirect(t *testing.T) {
	shorturl := "a"
	inPhrases := []string{BAD_SHORT_URL, "Enter the url...", "Get a shortened url:", "Submit"}
	outPhrases := []string{"Original URL:", "Shortened URL:", "localhost", GEN_ERROR, INVALID_URL}
	p := page{"GET",  SHORT_URL_ROUTE + shorturl, "short url redirect page", inPhrases, outPhrases, http.StatusOK, url.Values{}}
	genericPageTest(t, p)

}
