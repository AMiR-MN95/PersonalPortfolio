package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Start is the exported Gin router
var Start *gin.Engine

func init() {
	// Initialize the Gin router
	Start = gin.Default()

	// Serve static files from the "static" directory
	Start.Static("/static", "./static")

	// Set up HTML template rendering
	Start.LoadHTMLGlob("templates/*")

	// Initialize routes
	initRoutes()
}

// showHomePage handles the root route
func showHomePage(c *gin.Context) {
	// Render the "index.html" template
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "AMN-Portfolio",
	})
}
