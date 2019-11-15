default: run

depends:
	govendor sync

test:
	go test -cover -covermode=count -coverprofile=coverage.out ./...

cover: test
	go tool cover -html=coverage-all.out

run:
	go run server.go

build: clean
	go build -a -o server server.go

clean:
	rm -rf server coverage.out coverage-all.out

dockerrun:
	sudo chmod -R 777 ./data && docker-compose up --build
