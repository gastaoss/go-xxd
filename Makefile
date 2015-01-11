BINARY = xxd

all: $(BINARY)

**/*.go:
	go build ./...

$(BINARY): **/*.go src/*.go

deps:
	go get .

build: $(BINARY)

clean:
	rm $(BINARY)

run: $(BINARY)
	./$(BINARY) -vv

debug: $(BINARY)
	./$(BINARY) --pprof 6060 -vv

test:
	go test .
	golint

