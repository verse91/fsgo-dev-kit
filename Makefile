.PHONY: build clean
build:
	go build -o build/fsgo .
	sudo cp build/fsgo /usr/local/bin/
	@echo "DONE"

clean:
	rm -f fsgo
