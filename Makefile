.SILENT: build install clean

build:
	go build -o bin/lignaj main.go

install: build
	cp bin/lignaj /usr/local/bin/lignaj

clean:
	rm -rf ./bin