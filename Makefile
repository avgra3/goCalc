intro:
	@echo "Available commands:"
	@echo "make test: runs unit tests"
	@echo "make build: builds the project into the build directory"
	@echo "make install: installs the go app"
test:
	go test -v ./...
run: test
	go run main.go
build: test
	go build -o build/goCalc
install: test
	go install
