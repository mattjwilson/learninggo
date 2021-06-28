// Package datafile allows reading data samples from files.
package datafile

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// GetFloats reas a float64 from each line of a file.
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, err
	}

	return numbers, nil
}

// Maximum takes a slice of float64s and returns the maximum in the list.
func Maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}
