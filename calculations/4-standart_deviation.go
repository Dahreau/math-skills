package calculations

import "math"

func Stdv(variance float64) float64 {
	stdv := math.Sqrt(variance)
	return stdv
}
