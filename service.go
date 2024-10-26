package main

import (
	"errors"
	"math"
)

func calculateSpecificVolumes(pressure float64) (PhaseChangeResponse, error) {
	// Punto crítico
	if math.Abs(pressure-10.0) < 0.0001 {
		return PhaseChangeResponse{
			SpecificVolumeLiquid: 0.0035,
			SpecificVolumeVapor:  0.0035,
		}, nil
	}

	// Validar rango de presión para T > 30°C
	if pressure < 0.05 || pressure > 10.0 {
		return PhaseChangeResponse{}, errors.New("pressure out of range for T > 30°C")
	}

	// Puntos conocidos de la curva de saturación para T > 30°C
	pressureData := map[float64]PhaseChangeResponse{
		0.05: {SpecificVolumeLiquid: 0.0010, SpecificVolumeVapor: 30.0000},
		1.0:  {SpecificVolumeLiquid: 0.0012, SpecificVolumeVapor: 2.0000},
		5.0:  {SpecificVolumeLiquid: 0.0025, SpecificVolumeVapor: 0.5000},
		10.0: {SpecificVolumeLiquid: 0.0035, SpecificVolumeVapor: 0.0035},
	}

	// Encontrar puntos para interpolación
	var lowerPressure, upperPressure float64
	var lowerVolume, upperVolume PhaseChangeResponse
	found := false

	pressures := []float64{0.05, 1.0, 5.0, 10.0}
	for i, p := range pressures {
		if p == pressure {
			return pressureData[p], nil
		} else if p < pressure {
			lowerPressure = p
			lowerVolume = pressureData[p]
			if i+1 < len(pressures) {
				upperPressure = pressures[i+1]
				upperVolume = pressureData[pressures[i+1]]
				found = true
			}
		}
	}

	if !found {
		return PhaseChangeResponse{}, errors.New("could not interpolate values")
	}

	// Interpolación
	ratio := (pressure - lowerPressure) / (upperPressure - lowerPressure)

	// Interpolación lineal para líquido y vapor
	specificVolumeLiquid := lowerVolume.SpecificVolumeLiquid +
		ratio*(upperVolume.SpecificVolumeLiquid-lowerVolume.SpecificVolumeLiquid)
	specificVolumeVapor := lowerVolume.SpecificVolumeVapor +
		ratio*(upperVolume.SpecificVolumeVapor-lowerVolume.SpecificVolumeVapor)

	// Redondear a 4 decimales
	specificVolumeLiquid = math.Round(specificVolumeLiquid*10000) / 10000
	specificVolumeVapor = math.Round(specificVolumeVapor*10000) / 10000

	// Validar volumen mínimo
	if specificVolumeLiquid < 0.0010 {
		return PhaseChangeResponse{}, errors.New("volume below minimum threshold for T > 30°C")
	}

	response := PhaseChangeResponse{
		SpecificVolumeLiquid: specificVolumeLiquid,
		SpecificVolumeVapor:  specificVolumeVapor,
	}

	// Validar que el volumen de vapor sea mayor que el líquido excepto en el punto crítico
	if response.SpecificVolumeVapor <= response.SpecificVolumeLiquid && pressure < 10.0 {
		return PhaseChangeResponse{}, errors.New("invalid phase change state")
	}

	return response, nil
}
