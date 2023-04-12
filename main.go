package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ahmed-Mas/url_shorten_back/urlshorten/server"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	r.Use(cors.New(config))
	r.POST("/api/v1/short", server.GenerateShort)
	r.GET("/:url", server.RedirectToLong)

	log.Println("starting server at")
	r.Run(fmt.Sprintf(":%s", port))
}
