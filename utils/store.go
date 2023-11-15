package utils

import (
	"io/ioutil"
	"strings"
	"web-parser/constants"
)

// ArchiveContent ...
// Will store html content of url in local system
func ArchiveContent(url string, content []byte) (err error) {
	// Extract the domain from the URL
	domain := strings.TrimPrefix(url, constants.Http)
	domain = strings.TrimPrefix(domain, constants.Https)
	domain = strings.Replace(domain, "/", "_", -1)
	filePath := domain + constants.Html
	// Write file on system
	err = ioutil.WriteFile(filePath, content, 0644)
	return err
}
