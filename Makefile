
run:
	go run main.go ./example.csv ./plan.ics

clean:
	rm -rf dist

build: clean
	mkdir dist
	go build -o dist/csvToIcs main.go
	cp example.csv dist/
	cp example.sh dist/
	cp readme.md dist/