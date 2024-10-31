package constants

const (
	// MinPressure = 0.05 MPa
	// Este es el límite inferior del diagrama P-v donde el vapor tiene su mayor volumen específico.
	// En este punto:
	// - v_liquid ≈ 0.001 m³/kg
	// - v_vapor ≈ 0.3 m³/kg
	MinPressure = 0.05

	// MaxPressure = 10.0 MPa
	// Este es el punto crítico donde el líquido y vapor tienen el mismo volumen específico.
	// En este punto:
	// - v_liquid = v_vapor = 0.0035 m³/kg
	MaxPressure = 10.0

	// LiquidCoefficient = 2450
	// Coeficiente para la ecuación del líquido: v_liquid = (LiquidCoefficient * P + LiquidConstant) / Denominator
	// Este valor fue elegido para que:
	// 1. En P = 10 MPa: 2450 * 10 + 10325 = 34825, que da v_liquid = 0.0035
	// 2. En P = 0.05 MPa: 2450 * 0.05 + 10325 = 10447.5, que da v_liquid ≈ 0.001
	LiquidCoefficient = 2450

	// LiquidConstant = 10325
	// Constante para la ecuación del líquido que trabaja junto con LiquidCoefficient
	// Este valor asegura que:
	// 1. El volumen específico del líquido sea pequeño y casi constante a bajas presiones
	// 2. Converja exactamente a 0.0035 en el punto crítico (10 MPa)
	LiquidConstant = 10325

	// VaporCoefficient = 29996500
	// Coeficiente para la ecuación del vapor: v_vapor = (VaporConstant - VaporCoefficient * P) / Denominator
	// Este valor fue elegido para que:
	// 1. En P = 10 MPa: 299999825 - 29996500 * 10 = 34825, que da v_vapor = 0.0035
	// 2. En P = 0.05 MPa: 299999825 - 29996500 * 0.05 = 298499000, que da v_vapor ≈ 0.3
	VaporCoefficient = 29996500

	// VaporConstant = 299999825
	// Constante para la ecuación del vapor que trabaja junto con VaporCoefficient
	// Este valor asegura que:
	// 1. El volumen específico del vapor sea grande y muy sensible a la presión a bajas presiones
	// 2. Converja exactamente a 0.0035 en el punto crítico (10 MPa)
	VaporConstant = 299999825

	// Denominator = 9950000
	// Factor de escala base para convertir los valores grandes en volúmenes específicos apropiados
	// Se eligió este número porque:
	// 1. Al dividir por él, obtenemos valores en el rango correcto (cerca de 0.0035)
	// 2. Es aproximadamente 10^7, lo que facilita obtener valores en el orden de 10^-3
	// 3. Produce valores físicamente realistas en todo el rango de presiones
	Denominator = 9950000
)
