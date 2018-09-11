# url-shortener

## Start mysql database "db" docker container
- Run: `docker run -d --name db -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=urldb -e MYSQL_USER=user -e MYSQL_PASSWORD=password mysql`
- To get my table of urls, use the database `dump url-db/urldb_dump.sql` (some unit tests are dependent on urls in my table).

## Create a network and connect the db
- Run: `docker network create --driver=bridge`
- Run: `docker network connect db`

## Run the server docker container
- Find the the ip address of the db docker container and make sure it is correct in the constants of the golang web app.
- Build the container: `docker build -t myapp .`
- Run: `docker run --name myapp --rm -it -p 8080:8080 myapp`
- Connect to network: `docker network connect myapp`

## Run the test suite
- Run: `docker exec -i -t myapp /bin/bash`
- Go into `src` and run: `go test -v -cover`

## Usage after running containers
- Go to http://localhost:8080/.
