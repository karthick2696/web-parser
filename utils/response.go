package utils

import (
	"fmt"
	"time"
	"web-parser/constants"
)

// PrintMetaData ..
// function will print meta data in required format
func PrintMetaData(url string, anchorCount, imageCount, index int) {
	// check for index greater than zero
	if index > constants.Zero {
		fmt.Println("--------------------------------------------")
	}
	fmt.Println(constants.Site, url)
	fmt.Println(constants.NumLinks, anchorCount)
	fmt.Println(constants.Images, imageCount)
	fmt.Println(constants.LastFetch, time.Now().UTC().Format(constants.TimeFormat))
}
