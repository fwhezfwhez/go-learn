package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/weather-report/", reportWeather)
	r.Run(":8080")
}

func reportWeather(c *gin.Context) {
	var weathers = []string{
		"晴朗 37°",
		"阴天 39°",
		"大雨 29°",
	}
	c.JSON(200, gin.H{"message": weathers[time.Now().Unix()%3]})
}
