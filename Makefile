.PHONY: clean build

all: clean build

build: 
	@echo "compiling:"
	@$(MAKE) --no-print-directory  build-host
	@$(MAKE) --no-print-directory  build-linux-arm64
	@$(MAKE) --no-print-directory  build-linux-arm
	@$(MAKE) --no-print-directory  build-mac-arm64
	@$(MAKE) --no-print-directory  build-mac-amd64
	@echo "build targets created:"
	@ls -1 ./bin/


build-host:
	@echo "building target: host"
	@go build -o bin/main-host-arch main.go

build-linux-arm64:
	@echo "building target: linux-arm64"
	@GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go

build-linux-arm:
	@echo "building target: linux-arm"
	@GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go

build-mac-arm64:
	@echo "building target: darwin-arm64"
	@GOOS=darwin GOARCH=arm64 go build -o bin/main-mac-arm64 main.go

build-mac-amd64:
	@echo "building target: darwin-amd64"
	@GOOS=darwin GOARCH=amd64 go build -o bin/main-mac-amd64 main.go

run:
	go run main.go

clean:
	@echo "cleaning..."
	@rm -rf ./bin/*
