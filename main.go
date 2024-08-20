package main

import (
	"fmt"
	"tvguide/api"

	"github.com/gin-gonic/gin"
)

func main() {
	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()

	r := gin.Default()
	api.Start(r)
}
