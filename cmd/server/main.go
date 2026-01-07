package main

import (
	"log"

	"github.com/hans312456/go-gin-project/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.RegisterRoutes(r)

	log.Fatal(r.Run(":8080"))
}
