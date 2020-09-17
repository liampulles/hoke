# Keep test at the top so that it is default when `make` is called.
# This is used by Travis CI.
ci: clean coverage.txt
coverage.txt: install
	go test -race -coverprofile=coverage.txt -covermode=atomic -coverpkg=./pkg/...,./ ./...
view-cover: clean coverage.txt
	go tool cover -html=coverage.txt
build:
	go build ./...
install: build
	go install ./...
run:
	hoke https://github.com
inspect: build
	golint ./...
update:
	go get -u ./...
pre-commit: update clean coverage.txt inspect
	go mod tidy
clean:
	rm -f ${GOPATH}/bin/hoke
	rm -f coverage.txt