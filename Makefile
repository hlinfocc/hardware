export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on
LDFLAGS := -s -w

all: fmt build

build: hardware


fmt:
	go fmt ./...

fmt-more:
	gofumpt -l -w .

vet:
	go vet ./...


hardware:
	env CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o bin/hlinfo-hardware ./

test: gotest

gotest:
	go test -v --cover ./pkg/...

	
clean:
	rm -f ./bin/hlinfo-hardware
