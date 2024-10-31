package main

import (
	_ "altscore/docs"

	"altscore/handlers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Phase Change Diagram API
// @version 1.0
// @description API para obtener datos de volumen específico en función de la presión.
// @BasePath /

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/phase-change-diagram", handlers.GetPhaseChangeDiagram)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
