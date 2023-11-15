package main

import (
	"flag"
	"log"
	"web-parser/constants"
	"web-parser/logic"
)

// main ...
// Entry point for webpage parser script
func main() {
	// Reading metadata required or not flah
	withMetadata := flag.Bool(constants.Metadata, false, "Record metadata for the fetched web pages")
	flag.Parse()
	// Reading url from arguments
	urls := flag.Args()
	// Check for url found or not
	if len(urls) == constants.Zero {
		log.Println("Altest one webpage url required")
		return
	}
	// Call for parse html content of url
	logic.ParseURL(urls, *withMetadata)
}
