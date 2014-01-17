package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Result struct {
	Domain       string
	Availability string
}

type SearchResults struct {
	Query   string
	Results []Result
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Missing search query. Specify a string to search domainr for.")
	}

	var query string = os.Args[1]

	httpResponse, _ := http.Get("https://domai.nr/api/json/search?client_id=domainr_command_line_app&q=" + query)

	defer httpResponse.Body.Close()
	body, _ := ioutil.ReadAll(httpResponse.Body)

	var sr SearchResults

	// Decode json string into custom structs.
	json.Unmarshal(body, &sr)

	// Print results to stdout
	fmt.Printf("\n Results for \"%s\"\n\n", sr.Query)
	for _, result := range sr.Results {
		var available string
		switch result.Availability {
		case "available":
			available = "✔"
		default:
			available = "✘"
		}
		fmt.Printf("     %s %s\n", available, result.Domain)
	}
	fmt.Printf("\n")
}
