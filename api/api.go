package api

import (
	"fmt"
	"net/http"
	"time"
	"tvguide/getdata"
	"tvguide/model"
	"tvguide/textops"

	"github.com/gin-gonic/gin"
)

var urls = textops.OpenFile("url.txt")
var apiKey = textops.OpenFile("ak.txt")[0]
var searches = textops.OpenFile("channel.txt")
var channels []model.Channel

func AttachApi(r *gin.Engine) {
	// Routes
	r.GET("/schedule", getSchedule)
}

func getSchedule(c *gin.Context) {
	// This is for ?name= query
	name := c.Query("name")
	if len(name) == 0 {
		c.JSON(http.StatusOK, channels)
		return
	}
	result := getdata.SearchChannel(channels, name, false)
	c.JSON(http.StatusOK, result)
}

func getScheduleJob() {
	allChannelData := getdata.FetchScheduleData(urls, apiKey)
	channels = getdata.SearchThroughList(allChannelData, searches)
}

func RunGetScheduleJobOnce() []model.Channel {
	getScheduleJob()
	return channels
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
