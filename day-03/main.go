package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Point struct {
	x     int
	y     int
	steps int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func trimFirst(s string) int {
	_, i := utf8.DecodeRuneInString(s)

	n, err := strconv.Atoi(s[i:])
	check(err)
	return n
}

func contains(s []Point, e Point) bool {
	for _, a := range s {
		if a.x == e.x && a.y == e.y {
			return true
		}
	}
	return false
}

// Reads input file at a given path, parses as string and splits into array based commas
func readInput(path string) ([]string, []string) {
	data, err := ioutil.ReadFile(path)
	check(err)

	stringSlice := strings.Split(string(data), "\n")

	if len(stringSlice) != 2 {
		panic("Invalid Input 👽")
	}

	lineOne := strings.Split(stringSlice[0], ",")
	lineTwo := strings.Split(stringSlice[1], ",")
	return lineOne, lineTwo
}

func plotPoints(line []string) []Point {
	x, y := 0, 0
	var points []Point
	totalSteps := 0

	for _, step := range line {
		direction := string(step[0])
		trimFirst(step)

		for i := 0; i < trimFirst(step); i++ {
			totalSteps++

			switch direction {
			case "U":
				y++
			case "D":
				y--
			case "L":
				x--
			case "R":
				x++
			}

			points = append(points, Point{x, y, totalSteps})
		}
	}

	return points
}

func correctCartesian(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func determineIntersections(lineOne []Point, lineTwo []Point) int {
	var lowestDistance int

	for _, v := range lineTwo {
		if contains(lineOne, v) {
			manhattanDistance := correctCartesian(v.x) + correctCartesian(v.y)

			if lowestDistance == 0 || manhattanDistance < lowestDistance {
				lowestDistance = manhattanDistance
			}
		}
	}

	return lowestDistance
}

func determineSteps(lineOne []Point, lineTwo []Point) int {
	var lowestDistance int

	for _, lineTwoPoint := range lineTwo {
		if contains(lineOne, lineTwoPoint) {
			lineOneSteps := 0
			for _, lineOnePoint := range lineOne {
				if lineOnePoint.x == lineTwoPoint.x && lineOnePoint.y == lineTwoPoint.y {
					lineOneSteps = lineOnePoint.steps
				}
			}

			intersectionSteps := lineTwoPoint.steps + lineOneSteps

			if lowestDistance == 0 || intersectionSteps < lowestDistance {
				lowestDistance = intersectionSteps
			}
		}
	}

	return lowestDistance
}

func main() {
	lineOne, lineTwo := readInput("./test.txt")
	first := plotPoints(lineOne)
	second := plotPoints(lineTwo)

	fmt.Println("Manhattan Distance", determineIntersections(first, second))
	fmt.Println("Lowest Step Number", determineSteps(first, second))
}
