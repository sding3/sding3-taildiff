.PHONY: build install clean

BINARY := "taildiff"

build:
	go build -o $(BINARY)

install: build
	sudo cp $(BINARY) /usr/local/bin

clean:
	rm -f $(BINARY)
