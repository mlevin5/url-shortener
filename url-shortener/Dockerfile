FROM golang:1.8.5-jessie
# install glide
RUN go get github.com/golang/dep/cmd/dep
# create a working directory
WORKDIR /go/src/app
# add glide.yaml and glide.lock
ADD Gopkg.toml Gopkg.toml
ADD Gopkg.lock Gopkg.lock
# install packages
RUN dep ensure --vendor-only
# add source code
ADD src src
# run main.go
CMD ["go", "run", "src/main.go"]
EXPOSE 8080
