package models

// PhaseChangeResponse representa la respuesta JSON esperada
type PhaseChangeResponse struct {
	SpecificVolumeLiquid float64 `json:"specific_volume_liquid"`
	SpecificVolumeVapor  float64 `json:"specific_volume_vapor"`
}
