package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: go run scraper.go 'URL' \n")
		os.Exit(1)
	}

	doc, err := htmlquery.LoadURL(os.Args[1])
	if err != nil {
		panic(err)
	}

	list, err := htmlquery.QueryAll(doc, "//ol/li")
	if err != nil {
		panic(err)
	}
	for i, n := range list {
		a := htmlquery.FindOne(n, "//a")
		fmt.Printf("%d %s(%s)\n", i, htmlquery.InnerText(a), htmlquery.SelectAttr(a, "href"))
	}
}
