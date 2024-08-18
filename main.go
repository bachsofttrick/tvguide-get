package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tvguide/model"
	"tvguide/readtext"
)

// fetchJSONData retrieves the JSON data from the API endpoint
func fetchJSONData(url string, printToTerm bool) []model.Channel {
	// Create a custom request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}

	// Start getting the JSON
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Decode the body to an object
	var guide model.TVGuide
	err = json.Unmarshal(body, &guide)
	if err != nil {
		panic(err)
	}

	// Get the schedule and print to terminal
	schedule := guide.Data.Channels
	if printToTerm {
		prettyJSON, err := json.MarshalIndent(schedule, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(prettyJSON))
	}

	return schedule
}

func main() {
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	urls := readtext.OpenTextFile("url.txt")
	fetchJSONData(urls[0], true)
}
