package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	// Set up HTML template rendering
	r.LoadHTMLGlob("templates/*")

	// Define routes
	r.GET("/", showHomePage)

	// Run the server
	r.Run(":8080")
}

func showHomePage(c *gin.Context) {
	// Render the "index.html" template
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Your Portfolio",
	})
}
