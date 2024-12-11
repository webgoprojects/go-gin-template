package main

import (
	"github.com/gin-gonic/gin"
	"github.com/webgoprojects/go-gin-template/controllers"
	"github.com/webgoprojects/go-gin-template/middleware"
)

func main() {
	r := gin.Default()

	// Set trusted proxies
	_ = r.SetTrustedProxies([]string{"192.168.0.1"})

	// Middleware
	r.Use(middleware.Logger())

	// Routes
	r.GET("/", controllers.Index)
	r.GET("/about", controllers.About)

	// Serve static files
	r.Static("/static", "./static")

	// Start the server
	r.Run(":8080")
}
