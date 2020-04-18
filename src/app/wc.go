package main

import (
	"flag"
	"fmt"

	"github.com/jake-hansen/wc/src/app/utilities"
)

func main() {
	fmt.Println("Welcome to Go WC")
	fileptr := flag.String("filepath", "test.txt", "file to read from")
	flag.Parse()
	utilities.Readfile(*fileptr)
}
