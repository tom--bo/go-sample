package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func sleep(c *gin.Context) {
	sleepTime, err := strconv.Atoi(c.Param("stime"))
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(sleepTime) * time.Second)

	c.String(http.StatusOK, "OK, slept %s seconds.", c.Param("stime"))
}

func main() {
	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/sleep/:stime", sleep)

	r.Run(":8080")
}
