package calculations

func Average(values []int) float64 {
	var average float64
	for _, integer := range values {
		average += float64(integer)
	}
	average /= float64(len(values))
	return average
}
