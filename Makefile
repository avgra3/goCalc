intro:
	@echo "Available commands:"
	@echo "make test: runs unit tests"
	@echo "make build: builds the project into the build directory"
	@echo "make install: installs the go app"
test:
	@echo "Running tests..."
	@go test -v ./...
	@echo "Running gosec..."
	gosec -enable-audit ./...
run: test
	@go run main.go
build: test
	@go build -o build/gocalc
	@echo "Build finished and located in build/gocalc"
install: test
	@go install
format:
	@echo "Running go fmt..."
	go fmt ./...
