package calculations

import (
	"sort"
)

func Median(values []int) float64 {
	var median float64
	sort.Ints(values)
	middle := (len(values) - 1) / 2
	if len(values)%2 == 0 {
		median = (float64(values[middle]) + float64(values[middle+1])) / 2
	} else if len(values)%2 != 0 {
		median = float64(values[middle])
	}
	return median
}
