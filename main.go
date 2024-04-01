package main

import (
	"github.com/Pourroy/provisionong-connector-go-api/scr/configuration/logger"
	"github.com/Pourroy/provisionong-connector-go-api/scr/controller/routes"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .enf file", err)
	}
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(requestid.New())
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		logger.Fatal("Error to initialize router", err)
	}
}
