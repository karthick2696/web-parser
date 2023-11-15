package logic

import (
	"io"
	"log"
	"strings"
	"web-parser/constants"

	"web-parser/utils"

	"golang.org/x/net/html"
)

// ParseURL ...
// Will receive urls and metadata flag as input based that will fetch
func ParseURL(urls []string, metadata bool) {
	// iterating through the urls
	for index, url := range urls {
		var err error
		// call for fetch html content from url
		var htmlContent []byte
		if htmlContent, err = utils.MakeHttpCall(url); err != nil {
			log.Println(err, url)
			continue
		}
		// check for meta data required or not
		if metadata {
			// parsing html and getting element count
			anchorCount, imageCount := parseHTML(htmlContent)
			// call for print parsed data
			utils.PrintMetaData(url, anchorCount, imageCount, index)
		}
		// archive html data into local file system
		if err = utils.ArchiveContent(url, htmlContent); err != nil {
			log.Println(err)
		}
	}
}

// parseHTML ...
// will parse the provided html content and return the count of tag elements
func parseHTML(inputData []byte) (anchorCount int, imageCount int) {
	// converting response as html tokenizer
	content := io.NopCloser(strings.NewReader(string(inputData)))
	htmlTokens := html.NewTokenizer(content)
	// itrating through html content and finding count of tag element
loop:
	for {
		tokenType := htmlTokens.Next()
		switch tokenType {
		case html.ErrorToken:
			break loop
		case html.StartTagToken, html.SelfClosingTagToken:
			tag := htmlTokens.Token()
			// check for anchor
			if tag.Data == constants.AnchorTag {
				anchorCount++
			} else if tag.Data == constants.ImgTag { // check for image
				imageCount++
			}
		}
	}
	return
}
