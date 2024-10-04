package tools

import (
	"fmt"
	"math"
)

func Display(average float64, median float64, variance float64, stdv float64) {
	fmt.Println("Average:", int(math.Round(average)), "\nMedian:", int(math.Round(median)), "\nVariance:", int(math.Round(variance)), "\nStandard Deviation:", int(math.Round(stdv)))
}
