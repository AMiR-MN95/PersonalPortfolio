package main

import (
	"PersonalPortfolio/handlers" // Import to run the init function in handler.go
)

func main() {
	// Run the server
	handlers.R.Run(":8080")
}
