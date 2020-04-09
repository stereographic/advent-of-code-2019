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
func readInput(path string) []int {
	data, err := ioutil.ReadFile(path)
	check(err)

	stringSlice := strings.Split(string(data), ",")
	numSlice := make([]int, 0, len(stringSlice))

	for _, v := range stringSlice {
		number, err := strconv.Atoi(v)
		check(err)

		numSlice = append(numSlice, number)
	}

	return numSlice
}

func manageIntcodes(intcodes []int, positionOne int, positionTwo int) []int {
	intcodes[1] = positionOne
	intcodes[2] = positionTwo

	return intcodes
}

func processIntcode(intcodes []int) []int {
	statusCode := 0
	index := 0

	for statusCode != 99 {
		opperator := intcodes[index]
		positionOne := index + 1
		positionTwo := index + 2
		productPosition := index + 3

		if opperator == 1 {
			intcodes[intcodes[productPosition]] = intcodes[intcodes[positionOne]] + intcodes[intcodes[positionTwo]]
			index = index + 4
		} else if opperator == 2 {
			intcodes[intcodes[productPosition]] = intcodes[intcodes[positionOne]] * intcodes[intcodes[positionTwo]]
			index = index + 4
		} else if opperator == 99 {
			statusCode = 99
			break
		} else {
			panic("Unhandled Intcode")
		}
	}

	return intcodes
}

func partOne(input []int) {
	partOneInput := make([]int, len(input))
	copy(partOneInput, input)
	result := processIntcode(manageIntcodes(partOneInput, 12, 2))[0]
	fmt.Println("Part One:", result)
}

func partTwo(input []int) {
	for i := 0; i <= 99; i++ {
		for n := 0; n <= 99; n++ {

			currentInput := make([]int, len(input))
			copy(currentInput, input)

			result := processIntcode(manageIntcodes(currentInput, i, n))[0]

			if result == 19690720 {
				fmt.Println("Part Two:", 100*i+n)
			}
		}
	}
}

func main() {
	inputSlice := readInput("./input.txt")

	partOne(inputSlice)
	partTwo(inputSlice)
}
