FROM golang:1.12.9-alpine
RUN apk add --no-cache git make gcc musl-dev linux-headers
# install govendor and compile daemon
RUN go get github.com/kardianos/govendor && \
	go get github.com/githubnemo/CompileDaemon
# create a working directory
WORKDIR /go/src/github.com/anielski/download-url-in-go
# add vendor files
ADD vendor vendor/
# install packages
RUN govendor sync
# add source code
ADD . .
# run main.go
CMD CompileDaemon -log-prefix=false -build="go build -a -installsuffix cgo" -command="./download-url-in-go"
