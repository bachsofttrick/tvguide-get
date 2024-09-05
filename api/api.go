package api

import (
	"fmt"
	"net/http"
	"time"
	"tvguide/getdata"
	"tvguide/model"
	"tvguide/readtext"

	"github.com/gin-gonic/gin"
)

var urls = readtext.OpenTextFile("url.txt")
var apiKey = readtext.OpenTextFile("ak.txt")[0]
var searches = readtext.OpenTextFile("channel.txt")
var channels []model.Channel

func AttachApi(r *gin.Engine) {
	// Routes
	r.GET("/schedule", getSchedule)
}

func getSchedule(c *gin.Context) {
	name := c.Query("name")
	if len(name) == 0 {
		c.JSON(http.StatusOK, channels)
		return
	}
	result := getdata.SearchChannel(channels, name, false)
	c.JSON(http.StatusOK, result)
}

func getScheduleJob() {
	channels = getdata.SearchForChannels(getdata.FetchScheduleData(urls, apiKey), searches)
}

func Start(r *gin.Engine) {
	// Start the ticker to run every 5 minutes
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	// Run the function immediately, then run by tick
	go func() {
		for {
			fmt.Println("Updating...")
			getScheduleJob()
			fmt.Println("Schedule updated")
			// Wait for the next tick
			<-ticker.C
		}
	}()

	AttachApi(r)
	r.Run()
}
