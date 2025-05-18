package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", showHome)

	r.Run(":8080")
}

func showHome(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", nil)
}
