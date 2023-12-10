package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// R is the exported Gin router
var R *gin.Engine

func init() {
	// Initialize the Gin router
	R = gin.Default()

	// Serve static files from the "static" directory
	R.Static("/static", "./static")

	// Set up HTML template rendering
	R.LoadHTMLGlob("templates/*")

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
