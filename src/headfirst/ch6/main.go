// Package calculates the average of several numbers
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var sum float64 = 0

	for _, argument := range arguments {
		// todo
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}

	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

	// numbers, err := datafile.GetFloats("data.txt")
	// var sum float64 = 0

	// for _, number := range numbers {
	// 	sum += number
	// }

	// sampleCount := float64(len(numbers))
	// fmt.Println("Average: %0.2f \n", sum/sampleCount)
}
