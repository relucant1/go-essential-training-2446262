// Writing a function that return Content-Type header
package main

import (
	"fmt"
	"net/http"
)

func main() {
	ctype, bodyd, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
		fmt.Println(bodyd)
	}
}

// contentType will return the value of Content-Type header returned by making an
// HTTP GET request to url
func contentType(url string) (string, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close() // Make sure we close the body

	ctype := resp.Header.Get("Content-Type")
	bodyd, err := resp.Body.Read([]byte{255})
	if ctype == "" { // Return error if Content-Type header not found
		return "", "", fmt.Errorf("can't find Content-Type header")
	}

	return ctype, string(bodyd), nil
}
