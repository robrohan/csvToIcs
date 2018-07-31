hash = $(shell git log --pretty=format:'%h' -n 1)
os = macOS10.13

run:
	go run main.go ./example.csv ./plan.ics

clean:
	rm -rf dist

build: clean
	mkdir dist
	go build -o dist/csvToIcs main.go
	cp example.csv dist/
	cp example.sh dist/
	cp example.numbers dist/
	cp readme.md dist/

distrib: build
	mv dist csvToIcs
	zip csvToIcs-$(os)-$(hash).zip ./csvToIcs/*
	rm -rf csvToIcs