package main

import (
	"log"

	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/logger"
	"github.com/Pourroy/provisionong-connector-go-api/scr/controller/routes"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting Application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .enf file")
	}

	router := gin.Default()
	router.Use(requestid.New())
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Error to initialize router", err)
	}
}
