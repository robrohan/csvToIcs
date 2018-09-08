package main

import (
	"bytes"
	"crypto/rand"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	subject = iota
	startDate
	startTime
	endDate
	endTime
	allDayEvent
	description
	location
	private
)
const (
	calendar = "Plan"
	prodID   = "//Rob Rohan//Made up go code//EN"
	timeZone = "NZDT"
)

type recordReader func([]string)

func usage() {
	fmt.Printf("\nUsage:\n\t%v <csv path> <output file>\n\n", os.Args[0])
	fmt.Printf("Example:\n\tcsvToICS /home/myfile.csv /home/plan.ics\n\n")
}

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}
	inFile := os.Args[1]
	outFile := os.Args[2]

	log.Printf("Using: %v\n", inFile)

	var ics bytes.Buffer

	log.Printf("Creating prolog")
	prolog(&ics)

	log.Printf("Parsing file")
	foo := func(record []string) {
		err := formatRecord(record, &ics)
		if err != nil {
			log.Printf("Couldn't format record: %v", err)
		}
	}
	err := parseFile(inFile, foo)
	if err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}

	log.Printf("Writing epilog")
	epilog(&ics)

	log.Printf("Creating ics")
	writeICS(outFile, &ics)

	os.Exit(0)
}

func formatRecord(record []string, event *bytes.Buffer) error {
	uuid, err := newid()
	if err != nil {
		panic("Bad id gen")
	}

	if len(record)-1 < private {
		fmt.Printf("%v %v", len(record), private)
		return errors.New("Bad record length")
	}

	formatDate := strings.Replace(record[startDate], "-", "", -1)
	if formatDate == "" {
		return errors.New("Event missing date")
	}

	event.WriteString("BEGIN:VEVENT\n")
	fmt.Fprintf(event, "DTSTAMP:%vT000000Z\n", formatDate)
	fmt.Fprintf(event, "UID:ROHAN-%v\n", uuid)
	fmt.Fprintf(event, "DTSTART;VALUE=DATE:%v\n", formatDate)
	fmt.Fprintf(event, "DTEND;VALUE=DATE:%v\n", formatDate)
	fmt.Fprintf(event, "SUMMARY:%v\n", record[subject])
	fmt.Fprintf(event, "DESCRIPTION:%v\n", record[description])
	fmt.Fprintf(event, "CATEGORIES:%v\n", calendar)
	event.WriteString("END:VEVENT\n")

	return nil
}

func prolog(prolog *bytes.Buffer) {
	prolog.WriteString("BEGIN:VCALENDAR\n")
	prolog.WriteString("VERSION:2.0\n")
	fmt.Fprintf(prolog, "X-WR-CALNAME:%v\n", calendar)
	fmt.Fprintf(prolog, "PRODID:%v\n", prodID)
	fmt.Fprintf(prolog, "X-WR-TIMEZONE:%v\n", timeZone)
	fmt.Fprintf(prolog, "X-WR-CALDESC:%v\n", calendar)
	prolog.WriteString("CALSCALE:GREGORIAN\n")
}

func epilog(epilog *bytes.Buffer) {
	epilog.WriteString("END:VCALENDAR\n")
}

func writeICS(path string, buffer *bytes.Buffer) {
	err := ioutil.WriteFile(path, buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func parseFile(path string, fn recordReader) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		// ignore header row
		if record[subject] != "Subject" {
			fn(record)
		}
	}

	return nil
}

// Thanks internet
// https://stackoverflow.com/questions/15130321/is-there-a-method-to-generate-a-uuid-with-go-language#15134490
// Close enough for jazz.
func newid() (string, error) {
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		return "", err
	}

	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F

	return hex.EncodeToString(u), nil
}
