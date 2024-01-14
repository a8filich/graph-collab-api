MAIN_PATH := ./cmd/main/.
BIN_PATH := ./bin/$(basename $(go list -m))
VERSION := 0.1.0
BUILD_REF := `git rev-parse HEAD`
LDFLAGS = -ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD_REF)"
EXT_PKG_NAME =

.PHONY: all get tidy test vet check run fmt simplify install build clean

all: tidy check simplify build

get:
	@go get $(EXT_PKG_NAME)

tidy:
	@go mod tidy

test:
	#@go clean -cache
	@go test -v ./...

vet:
	#@go clean -cache
	@go vet ./...

check: test vet

run:
	@go run $(MAIN_PATH)

fmt:
	@gofmt -l -w .

simplify:
	@gofmt -s -l -w .

install:
	@go install $(LDFLAGS) ./...

build:
	@go build $(LDFLAGS) -v -o $(BIN_PATH) ./...

clean:
	@go clean -r
	@rm -f $(BIN_PATH)
