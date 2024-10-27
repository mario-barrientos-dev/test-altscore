package main

import (
	_ "altscore/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Phase Change Diagram API
// @version 1.0
// @description API para obtener datos de volumen específico en función de la presión.
// @BasePath /

func main() {
	router := gin.Default()

	router.GET("/phase-change-diagram", getPhaseChangeDiagram)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("0.0.0.0:8000") // Ejecuta la API en el puerto 8080
}
