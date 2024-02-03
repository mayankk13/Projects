package main

import "os"

// "github.com/Projects/ecommerce/controllers"
// "github.com/Projects/ecommerce/database"
// "github.com/Projects/ecommerce/routes"
// "github.com/Projects/ecommerce/middleware"
// "github.com/gin-gonic/gin"

// the main job this function is to start the surver i.e., which ports to run on
func main() {
	port := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}
}
