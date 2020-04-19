package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"

	"github.com/jake-hansen/wordcount/src/utilities"
)

func main() {
	fileptr := flag.String("filepath", "test.txt", "file to read from")
	flag.Parse()
	filecontents := utilities.Readfile(*fileptr)

	scanner := bufio.NewScanner(bytes.NewReader(filecontents))

	// Print line count
	scanner.Split(bufio.ScanLines)
	linecount := 0
	for scanner.Scan() {
		linecount++
	}
	fmt.Print(linecount)
	scanner = bufio.NewScanner(bytes.NewReader(filecontents))

	// Print word count
	scanner.Split(bufio.ScanWords)
	wordcount := 0
	for scanner.Scan() {
		wordcount++
	}
	fmt.Print(" ", wordcount)
	scanner = bufio.NewScanner(bytes.NewReader(filecontents))

	// Print character count
	scanner.Split(bufio.ScanRunes)
	charactercount := 0
	for scanner.Scan() {
		charactercount++
	}
	fmt.Print(" ", charactercount)

	fmt.Println(" ", *fileptr)
	
}
