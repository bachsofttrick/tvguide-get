package main

import (
	"flag"
	"fmt"
	"tvguide/api"
	"tvguide/getdata"
	"tvguide/prettier"

	"github.com/gin-gonic/gin"
)

func main() {
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	single := flag.Bool("single", false, "Only run once to get schedule")
	name := flag.String("name", "", "Search for specific channels")
	flag.Parse()
	if !*single {
		apiRoute()
	}

	// Run get schedule once
	schedule := api.RunGetScheduleJobOnce()
	if len(*name) > 0 {
		schedule = getdata.SearchChannel(schedule, *name, false)
	}
	prettier.PrintFromObject(schedule, true)
}

func apiRoute() {
	r := gin.Default()
	api.Start(r)
}
