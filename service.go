package main

import (
	"errors"
)

func calculateSpecificVolumes(pressure float64) (PhaseChangeResponse, error) {
	// Validar rango de presión
	if pressure > 10 || pressure < 0.05 {
		return PhaseChangeResponse{}, errors.New("pressure out of range")
	}

	// Calcular volúmenes usando las fórmulas lineales
	specificVolumeLiquid := (2450*pressure + 10325) / 9950000
	specificVolumeVapor := (299999825 - 29996500*pressure) / 9950000

	return PhaseChangeResponse{
		SpecificVolumeLiquid: specificVolumeLiquid,
		SpecificVolumeVapor:  specificVolumeVapor,
	}, nil
}
