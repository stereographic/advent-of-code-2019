package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Reads input file at a given path, parses as string and splits into array based on a new line
func readInput(path string) []string {
	data, err := ioutil.ReadFile(path)
	check(err)

	array := strings.Split(string(data), "\n")

	return array
}

// Iterates over array, parses each entry string as float64.
// Conversion from float64 to int removes floating points, essentially rounding "down"
// Gets total fuel usage
func getFuelUsage(input []string) {
	// Part 1 Result
	initialSum := 0
	// Part 2 Result
	totalSum := 0

	for i := 0; i <= (len(input) - 1); i++ {
		value, err := strconv.ParseFloat(input[i], 32)
		check(err)
		usage := calculateUsage(value)

		initialSum += usage
		totalSum += usage

		destinationUsage := usage
		for destinationUsage > 0 {
			destinationUsage = calculateUsage(float64(destinationUsage))
			if destinationUsage < 0 {
				destinationUsage = 0
			} else {
				totalSum += destinationUsage
			}
		}
	}

	fmt.Println("Part 1", initialSum)
	fmt.Println("Part 2", totalSum)
}

func calculateUsage(value float64) int {
	return int(value/3) - 2
}

func main() {
	input := readInput("./input.txt")

	getFuelUsage(input)
}
