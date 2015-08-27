package main

import (
	"fmt"
	"os"
	
	ml "github.com/isido/missinglinks"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s target-url\n", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1]

	r := ml.ReaderFromUrl(url)
	links := ml.Links(r)

	for _, link := range links {
		link = ml.AddPrefix(url, link)
		status := ml.HTTPResponseCode(link)
		if status == 404 {
			fmt.Printf("%s\n", link)
		}
	}
}
