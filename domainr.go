package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	apiURL = "https://domai.nr/api/json/search?client_id=domainr_command_line_app&q="
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
		fmt.Println("Missing search query. Specify a string to search domainr for.")
		os.Exit(1)
	}

	var query string = os.Args[1]

	httpResponse, err := http.Get(apiURL + query)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var sr SearchResults

	// Decode json string into custom structs.
	json.Unmarshal(body, &sr)

	buf := bufio.NewWriter(os.Stdout)

	// Print results to stdout
	fmt.Fprintf(buf, "\n Results for \"%s\"\n\n", sr.Query)
	for _, result := range sr.Results {
		var available string
		switch result.Availability {
		case "available":
			available = "✔"
		default:
			available = "✘"
		}
		fmt.Fprintf(buf, "     %s %s\n", available, result.Domain)
	}
	fmt.Fprintf(buf, "\n")
	buf.Flush()
}
