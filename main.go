package main

import (
	"fmt"
	"tvguide/getdata"
	"tvguide/prettier"
	"tvguide/readtext"
)

func main() {
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	urls := readtext.OpenTextFile("url.txt")
	searches := readtext.OpenTextFile("channel.txt")
	channels := getdata.FetchJSONData(urls[0])
	searched := getdata.SearchForChannels(channels, searches)

	prettier.Print(searched)
}
