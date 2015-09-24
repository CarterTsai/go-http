package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
)

var (
	uri    = flag.String("uri", "https://www.google.com", "The URI")
	header = flag.Bool("header", false, "Show Header")
	html   = flag.Bool("html", true, "Show HTML")
)

func main() {
	flag.Parse()
	yellow := color.New(color.FgYellow).SprintFunc()
	res, err := http.Get(*uri)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if *html {
		body, _ := ioutil.ReadAll(res.Body)
		r := string(body)
		fmt.Printf("%c", r)
	}

	if *header {
		for key, value := range res.Header {
			fmt.Printf("\n %s = %s", yellow(key), value)
		}
		fmt.Println("\n")
	}
}
