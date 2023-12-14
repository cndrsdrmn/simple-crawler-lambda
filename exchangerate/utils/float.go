package utils

import "math"

func Round(val float64, precision uint) float64 {
	ratio := math.Pow10(int(precision))
	return math.Round(val*ratio) / ratio
}
