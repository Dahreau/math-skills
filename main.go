package main

import (
	"calculations"
	"tools"
)

func main() {
	values := tools.Initialization()

	average := calculations.Average(values)
	median := calculations.Median(values)
	variance := calculations.Variance(values, average)
	stdv := calculations.Stdv(variance)

	tools.Display(average, median, variance, stdv)

}
