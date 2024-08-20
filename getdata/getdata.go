package getdata

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
	"tvguide/model"
	"tvguide/mytime"
)

// Retrieves the JSON data from the API endpoint
func FetchScheduleData(url string) []model.Channel {
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
	schedule := &guide.Data.Channels
	// Iterate over the slice with range
	for i := range *schedule {
		ch := &(*schedule)[i] // Get pointer to the i-th Channel
		// Iterate over the Schedule slice within the Channel
		for j := range ch.Schedule {
			pg := &ch.Schedule[j] // Get pointer to the j-th Program
			// Update UTCStartTime
			pg.UTCStartTime = mytime.GetUTCTimeFromEpoch(pg.StartTime).Format(time.RFC3339)
		}
	}
	return *schedule
}

func SearchChannel(channels []model.Channel, channelName string, onlyOneMatch bool) []model.Channel {
	results := []model.Channel{}

	// Split multiple words in search
	words := strings.Split(channelName, " ")
	for _, ch := range channels {
		lowerChannelName := strings.ToLower(ch.Channel.Name)
		if strings.Contains(lowerChannelName, words[0]) {
			if wordlen := len(words); wordlen > 1 {
				matchAllWords := true
				// Try and find channel with multiple words in search
				for i := 1; i < wordlen; i++ {
					if !strings.Contains(lowerChannelName, words[i]) {
						matchAllWords = false
						break
					}
				}

				// If the search name doesn't match all words, skip to the next channel
				if !matchAllWords {
					continue
				}
			}
			results = append(results, ch)
			// Only get the first channel found
			if onlyOneMatch {
				break
			}
		}
	}
	return results
}

func SearchForChannels(channels []model.Channel, channelList []string) []model.Channel {
	searched := []model.Channel{}
	for _, sh := range channelList {
		result := SearchChannel(channels, sh, true)
		searched = append(searched, result...)
	}

	return searched
}
