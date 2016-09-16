// Package ParseHTML parses html of http website and returns root of html tree
package ParseHTML

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

/* FetchAndParse fetches a response from url server and parses response in html format, returning root of html tree */
func FetchAndParse(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetchAndParse(%s) : %v\n", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetchAndParse(%s) : Bad response\n", url)
	}
	htmlroot, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fetchAndParse(%s) : %v\n", url, err)
	}
	return htmlroot, nil
}
