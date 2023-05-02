package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ahmed-Mas/url_shorten_back/pkg/memstore"
	"github.com/Ahmed-Mas/url_shorten_back/pkg/router"
	"github.com/Ahmed-Mas/url_shorten_back/pkg/shorten"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	storage := memstore.NewMemStore()
	service := shorten.NewShorten(storage)
	handler := router.NewHandler(service)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	r.Use(cors.New(config))

	r.POST("/api/v1/short", handler.GenerateShort)
	r.GET("/:url", handler.RedirectToLong)

	log.Println("starting server at")
	r.Run(fmt.Sprintf(":%s", port))
}
