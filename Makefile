.PHONY=build test test-race clean
GOLDFLAGS += -s -w -extldflags "-static"
GOLDFLAGS += -X main.Version=$(shell git describe)
GOLDFLAGS += -X main.GitCommit=$(shell git rev-parse HEAD)
GOLDFLAGS += -X main.BuildTime=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
GOFLAGS = -mod=vendor -ldflags "$(GOLDFLAGS)"

%: bin/ayaya

bin/ayaya:
	CGO_ENABLED=0 go build $(GOFLAGS) -o bin/ayaya main.go

test:
	CGO_ENABLED=0 go vet -mod=vendor ./...
	CGO_ENABLED=0 go test -mod=vendor -cover ./...

test-race:
	CGO_ENABLED=0 go vet -mod=vendor ./...
	go test -mod=vendor -race -cover ./...

clean:
	$(RM) -rf bin
