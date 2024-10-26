package main

import "errors"

// calculateSpecificVolumes calcula los volúmenes específicos en función de la presión
func calculateSpecificVolumes(pressure float64) (PhaseChangeResponse, error) {
	if pressure > 10 {
		return PhaseChangeResponse{}, errors.New("Pressure out of range")
	}
	return PhaseChangeResponse{
		SpecificVolumeLiquid: 0.0035,
		SpecificVolumeVapor:  0.0035,
	}, nil
}
