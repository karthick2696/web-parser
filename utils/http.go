package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"web-parser/constants"
)

// MakeHttpCall ...
// fetches the provided URL and returns the response body or an error
func MakeHttpCall(url string) (data []byte, err error) {
	for {
		// call for get content using go http
		var resp *http.Response
		if resp, err = http.Get(url); err != nil {
			break
		}
		defer resp.Body.Close()
		// check for request is success or not
		if resp.StatusCode != http.StatusOK {
			err = errors.New(constants.HttpError)
			break
		}
		// converting as byte
		data, err = ioutil.ReadAll(resp.Body)
		break
	}
	return
}
