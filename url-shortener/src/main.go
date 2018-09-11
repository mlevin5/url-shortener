package main

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"net/url"
	"path"
)

/* Globals */

// database constants
var USER = "user"
var PASSWORD = "password"
var PROTOCAL = "tcp"
var DB_IP = "172.18.0.2"
var DB_PORT = "3306"
var DB_NAME = "urldb"
var TABLE_NAME = "long_urls"
var URL_COL_NAME = "url"
var ID_COL_NAME = "id"

// file location constants
var FILE_DIR = "/go/src/app/src"
var BASE_PATH = path.Join(FILE_DIR, "templates/base.html")
var MSG_PATH = path.Join(FILE_DIR, "templates/msg.html")
var URL_PATH = path.Join(FILE_DIR, "templates/url.html")

// hosting and routing constants
var PORT = "8080"
var SHORT_URL_ROUTE = "/p/"
var SHORT_URL = "http://localhost:" + PORT + SHORT_URL_ROUTE

// error message constants
var GEN_ERROR = "Sorry, there was an internal error."
var INVALID_URL = "Invalid URL: Must have a scheme, host, and path."
var BAD_SHORT_URL = "The short URL you tried to navigate to has not been made yet!"

// encoder constant
var ENCODER = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

/* Helper functions */

// reverse a string
func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// database id to short url
func encodeBase62(id int) string {
	var buffer bytes.Buffer
	for ok := true; ok; ok = (id > 0) {
		buffer.WriteRune(ENCODER[id%62])
		id = id/62
	}
	shortUrl := reverse(buffer.String())
	return shortUrl
}

// short url to database id
// returns error if character in string is outside of base 62 characters
func decodeBase62(shortURLString string) int {
	var id int32 = 0
	for _, char := range shortURLString {
		if 'a' <= char && char <= 'z' {
			id = id*62 + char - 'a'
		} else if 'A' <= char && char <= 'Z' {
			id = id*62 + char - 'A' + 26
		} else if '0' <= char && char <= '9' {
			id = id*62 + char - '0' + 52
		}
	}
	return int(id)
}

// makes sure long url is a valid url (has a scheme, host, and path)
// returns true if valid or false if not
func checkUrl(s string) bool {
	u, err := url.Parse(s)
	if err != nil { return false }
	if u.Scheme == "" || u.Host == "" || u.Path == "" {
		return false
	}
	return true
}

// error handling helper
func handleError(w http.ResponseWriter, actualErr error, displayErr string) {
	fmt.Println(actualErr)
	msgPageHandler(w, displayErr)
}


/* SQL functions */

// return id of long url if it is already in the database, else returns error
func getIdOfLongUrl(db *sql.DB, longurl string) (int, error) {
	var id int
	var url string
	dbQuery := "SELECT * FROM " + TABLE_NAME + " WHERE " + URL_COL_NAME + " = ?"
	err := db.QueryRow(dbQuery, longurl).Scan(&id, &url)
	if err != nil { return -1, err }
	return id, nil
}

// return long url in database given its id
func getLongUrlFromId(db *sql.DB, id int) (string, error) {
	var url string
	dbQuery := "SELECT " + URL_COL_NAME + " FROM " + TABLE_NAME + " WHERE " + ID_COL_NAME + " = ?"
	err := db.QueryRow(dbQuery, id).Scan(&url)
	if err != nil { return "", err }
	return url, nil
}

// insert long url into database, return id
func addLongUrlGetId(db *sql.DB, longurl string) (int, error) {
	dbQuery := "INSERT " + TABLE_NAME + " SET " + ID_COL_NAME + "=DEFAULT, " + URL_COL_NAME + "=?"
	stmt, err := db.Prepare(dbQuery)
	if err != nil { return -1, err }
	res, err := stmt.Exec(longurl)
	if err != nil { return -1, err }
	id, err := res.LastInsertId()
	if err != nil { return -1, err }
	return int(id), nil
}

/* Page Handlers */

// base page with a single message
func msgPageHandler(w http.ResponseWriter, msg string) {
	t, err := template.ParseFiles(MSG_PATH, BASE_PATH)
	if err != nil { fmt.Println(err) }
	varmap := map[string]interface{}{
		"msg": msg,
	}
	err = t.ExecuteTemplate(w, "base", varmap)
}

// home page with no messages yet
func mainHandler(w http.ResponseWriter, r *http.Request) {
	msgPageHandler(w, "")
}

// handles when the user submits a long url and responds with
// the long and short urls on the page or an error message
func makeShortURLHandler(w http.ResponseWriter, r *http.Request) {
	db, err := initializeDB()
	if err != nil { handleError(w, err, GEN_ERROR)}

	err = r.ParseForm()
	if err != nil { handleError(w, err, GEN_ERROR) }

	longUrl := r.Form.Get("longurl")
	if !checkUrl(longUrl) {
		msgPageHandler(w, INVALID_URL)
		return
	}

	id, err := getIdOfLongUrl(db, longUrl)
	if err != nil || id < 0 {
		id, err = addLongUrlGetId(db, longUrl)
		if err != nil { handleError(w, err, GEN_ERROR) }
	}
	shortUrl := SHORT_URL + encodeBase62(id)
	t, err := template.ParseFiles(URL_PATH, BASE_PATH)
	if err != nil { handleError(w, err, GEN_ERROR) }
	varmap := map[string]interface{}{
		"shortUrl": shortUrl,
		"longUrl": longUrl,
	}
	err = t.ExecuteTemplate(w, "base", varmap)
	if err != nil { handleError(w, err, GEN_ERROR) }
	db.Close()
}

// given a short url, redirects user to appropriate long url
// else brings back to home page with error message
func goToShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	db, err := initializeDB()
	if err != nil { handleError(w, err, GEN_ERROR)}

	shortUrl := mux.Vars(r)["shortUrl"]
	id := decodeBase62(shortUrl)
	newUrl, err := getLongUrlFromId(db, id)
	if err != nil { handleError(w, err, GEN_ERROR) }

	http.Redirect(w, r, newUrl, http.StatusSeeOther)
	db.Close()
}

/* Initialize DB and router functions */

// connects to the mysql database
func initializeDB() (*sql.DB, error) {
	var err error
	dbString := USER+":"+PASSWORD+"@"+PROTOCAL+"("+DB_IP+":"+DB_PORT+")/"+DB_NAME
	db, err := sql.Open("mysql", dbString)
	if err != nil { return nil, err }

	return db, nil
}

// initializes a router for endpoints
func initializeRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", mainHandler).Methods("GET")
	router.HandleFunc("/", makeShortURLHandler).Methods("POST")
	router.HandleFunc(SHORT_URL_ROUTE + "{shortUrl}", goToShortUrlHandler).Methods("GET")
	return router
}

/* Main function */

func main() {
	http.Handle("/", initializeRouter())
	http.ListenAndServe(":" + PORT, nil)
}






