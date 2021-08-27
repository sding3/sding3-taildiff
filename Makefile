.PHONY: build install clean

test:
	go test -cover -race ./...

build:
	go build -o ./bin/ ./...

install: build
	cp ./bin/* ~/go/bin/

clean:
	rm -f ./bin
