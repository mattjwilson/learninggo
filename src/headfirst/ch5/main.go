// Package calculates the average of several numbers
package main

import (
	"bin/datafile"
	"fmt"
)

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	numbers, err := datafile.GetFloats("data.txt")
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Println("Average: %0.2f \n", sum/sampleCount)
}
