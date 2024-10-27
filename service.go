package main

import (
	"errors"
)

const (
	minPressure       = 0.05
	maxPressure       = 10.0
	liquidCoefficient = 2450
	liquidConstant    = 10325
	vaporCoefficient  = 29996500
	vaporConstant     = 299999825
	denominator       = 9950000
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
