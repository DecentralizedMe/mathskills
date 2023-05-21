package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read the file path from command-line argument
	filePath := os.Args[1]

	// Read data from the file
	data, err := readDataFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate average
	average := calculateAverage(data)

	// Sort the data
	sortedData := sortData(data)

	// Calculate median
	median := calculateMedian(sortedData)

	// Calculate variance
	variance := calculateVariance(data, average)

	// Calculate standard deviation
	standardDeviation := calculateStandardDeviation(variance)

	// Print the results rounded to the nearest number
	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", standardDeviation)
}

// readDataFromFile reads the data from the given file path and returns a slice of numbers.
func readDataFromFile(filePath string) ([]float64, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var data []float64
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}

	return data, nil
}

// calculateAverage calculates the average of the given data.
func calculateAverage(data []float64) float64 {
	sum := 0.0
	for _, num := range data {
		sum += num
	}
	return sum / float64(len(data))
}

// sortData sorts the given data in ascending order.
func sortData(data []float64) []float64 {
	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)
	return sortedData
}

// calculateMedian calculates the median of the given sorted data.
func calculateMedian(sortedData []float64) float64 {
	length := len(sortedData)
	middle := length / 2
	if length%2 == 0 {
		return (sortedData[middle-1] + sortedData[middle]) / 2.0
	}
	return sortedData[middle]
}

// calculateVariance calculates the variance of the given data, given the average.
func calculateVariance(data []float64, average float64) float64 {
	sumSquaredDiffs := 0.0
	for _, num := range data {
		diff := num - average
		sumSquaredDiffs += diff * diff
	}
	return sumSquaredDiffs / float64(len(data))
}

// calculateStandardDeviation calculates the standard deviation from the given variance.
func calculateStandardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}
