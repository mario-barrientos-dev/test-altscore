package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// getPhaseChangeDiagram devuelve datos específicos del volumen para la presión indicada
// @Summary Obtener diagrama de cambio de fase
// @Description Retorna volumen específico de líquido y vapor en función de la presión.
// @Param pressure query number true "Presión en MPa" format(float)
// @Success 200 {object} PhaseChangeResponse
// @Router /phase-change-diagram [get]
func getPhaseChangeDiagram(c *gin.Context) {
	pressureParam := c.Query("pressure")
	pressure, err := strconv.ParseFloat(pressureParam, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pressure parameter"})
		return
	}

	response, err := calculateSpecificVolumes(pressure)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
