package calculations

import "math"

func Variance(values []int, average float64) float64 {
	var variance float64
	for _, integer := range values {
		variance += math.Pow((average - float64(integer)), 2)
	}
	variance /= float64(len(values))
	return variance
}
