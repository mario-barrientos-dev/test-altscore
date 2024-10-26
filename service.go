package main

import "errors"

// calculateSpecificVolumes calcula los volúmenes específicos en función de la presión.
func calculateSpecificVolumes(pressure float64) (PhaseChangeResponse, error) {
	// Definir los límites de presión
	const maxPressure = 10.0 // MPa
	const minPressure = 0.0  // MPa

	// Verificar si la presión está dentro del rango aceptable
	if pressure < minPressure || pressure > maxPressure {
		return PhaseChangeResponse{}, errors.New("pressure out of range")
	}

	// Asignación de volúmenes específicos en puntos de presión conocidos
	pressureData := map[float64]PhaseChangeResponse{
		0.0:  {SpecificVolumeLiquid: 0.0012, SpecificVolumeVapor: 12.0},
		1.0:  {SpecificVolumeLiquid: 0.001, SpecificVolumeVapor: 0.004},
		5.0:  {SpecificVolumeLiquid: 0.0025, SpecificVolumeVapor: 0.003},
		10.0: {SpecificVolumeLiquid: 0.0035, SpecificVolumeVapor: 0.0035},
	}

	var lowerPressure, upperPressure float64
	var lowerVolume, upperVolume PhaseChangeResponse

	// Encontrar los puntos de presión superior e inferior
	for p, volume := range pressureData {
		if p == pressure {
			return volume, nil
		} else if p < pressure {
			lowerPressure = p
			lowerVolume = volume
		} else if p > pressure {
			upperPressure = p
			upperVolume = volume
			break
		}
	}

	// Interpolación lineal
	ratio := (pressure - lowerPressure) / (upperPressure - lowerPressure)
	specificVolumeLiquid := lowerVolume.SpecificVolumeLiquid + ratio*(upperVolume.SpecificVolumeLiquid-lowerVolume.SpecificVolumeLiquid)
	specificVolumeVapor := lowerVolume.SpecificVolumeVapor + ratio*(upperVolume.SpecificVolumeVapor-lowerVolume.SpecificVolumeVapor)

	return PhaseChangeResponse{
		SpecificVolumeLiquid: specificVolumeLiquid,
		SpecificVolumeVapor:  specificVolumeVapor,
	}, nil
}
