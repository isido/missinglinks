package missinglinks

import (
	"fmt"
	"io"
//	"io/ioutil"
	"net/http"
	"os"
	"strings"
	
	"golang.org/x/net/html"
)

// Get reader from addr
func ReaderFromUrl(addr string) io.Reader {

	resp, err := http.Get(addr)
	if err != nil {
		fmt.Printf("Cannot get url: %s\n", err)
		os.Exit(1)
	}
	return resp.Body
	// TODO: ensure UTF-8 and make sure to close the reader
}

// Get HTTP response code
func HTTPResponseCode(addr string) int {

	resp, err := http.Get(addr) // Use GET instead of HEAD just in case
	if err != nil {
		fmt.Printf("Cannot get url: %s\n", err)
		os.Exit(1)
	}
	return resp.StatusCode
}

// Find all links in a page, page is represented as a reader
func Links(r io.Reader) []string {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Printf("Cannot parse html: %s\n", err)
		os.Exit(1)
	}

	links := []string{}
	
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					//fmt.Println(a.Val)
					links = append(links, a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links
}

// Add domain prefix, if missing
func AddPrefix(prefix string, url string) string {

	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	} else {
		return concatUrlParts(prefix, url)
	}
}

// Concatenate to two url parts
func concatUrlParts(prefix string, suffix string) string {
	if strings.HasSuffix(prefix, "/") {
		prefix = prefix[:len(prefix) - 1]
	}

	if !strings.HasPrefix(suffix, "/") {
		suffix = "/" + suffix
	}

	return prefix + suffix
}

