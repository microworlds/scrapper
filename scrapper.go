package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	data "github.com/microworlds/scrapper/packages"
)

func remoteDocument(addr string) *goquery.Document {
	var doc *goquery.Document

	// Make a GET request
	resp, err := http.Get(addr)
	if err != nil {
		fmt.Println("Error: Looks like there's a problem loading", addr, ". Try again.")
		os.Exit(1)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	html, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		os.Exit(1)
	}
	doc = html

	return doc
}

func file() *goquery.Document {
	// Loading file from file reader stream
	html := data.Data()

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func getLinks(doc *goquery.Document) []data.LinkData {
	var Fields []data.LinkData

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		link := item
		linkTitle := link.Text()
		linkHref, exist := link.Attr("href")
		if exist {
			linkvalue := data.LinkData{
				Index:    index,
				LinkName: linkTitle,
				URL:      linkHref,
			}

			//fmt.Println(linkvalue)
			//println(index, linkTitle, "-", linkHref)
			Fields = append(Fields, linkvalue)

		}
	})

	return Fields
}

func scrapeLinks() {
	pURL := flag.String("u", "http://www.bbc.com", "Enter the website you want to harvest the mock data")
	pfilename := flag.String("f", "data", "Enter the filename to save the data")
	flag.Parse()

	url := *pURL
	filename := *pfilename + ".json"

	doc := remoteDocument(url)

	links := getLinks(doc)
	fmt.Println(links)

	// Save the links to disk and convert to JSON API Server response
	json, err := json.Marshal(links)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(json))
	if err != nil {
		fmt.Println("Error: There was a problem writing the data to local disk. Copy the data on your terminal and save it to a file manually")
		os.Exit(1)
	}
	fmt.Println(links)
}

func main() {
	fmt.Println("Fetching data...")
	scrapeLinks()
}
