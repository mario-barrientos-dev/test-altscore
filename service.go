package main

import (
	"errors"
)

func calculateSpecificVolumes(pressure float64) (PhaseChangeResponse, error) {
	if pressure > maxPressure || pressure < minPressure {
		return PhaseChangeResponse{}, errors.New("pressure out of range")
	}
	specificVolumeLiquid := (liquidCoefficient*pressure + liquidConstant) / denominator
	specificVolumeVapor := (vaporConstant - vaporCoefficient*pressure) / denominator

	return PhaseChangeResponse{
		SpecificVolumeLiquid: specificVolumeLiquid,
		SpecificVolumeVapor:  specificVolumeVapor,
	}, nil
}
