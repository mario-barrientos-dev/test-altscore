package service

import (
	"altscore/constants"
	"altscore/models"
	"errors"
)

func CalculateSpecificVolumes(pressure float64) (models.PhaseChangeResponse, error) {
	if pressure > constants.MaxPressure || pressure < constants.MinPressure {
		return models.PhaseChangeResponse{}, errors.New("pressure out of range")
	}
	specificVolumeLiquid := (constants.LiquidCoefficient*pressure + constants.LiquidConstant) / constants.Denominator
	specificVolumeVapor := (constants.VaporConstant - constants.VaporCoefficient*pressure) / constants.Denominator

	return models.PhaseChangeResponse{
		SpecificVolumeLiquid: specificVolumeLiquid,
		SpecificVolumeVapor:  specificVolumeVapor,
	}, nil
}
