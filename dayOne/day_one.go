package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

const (
	MIN = 0
	MAX = 100
)

var partOneZeroCount = 0
var partTwoZeroCount = 0

func main() {
	partOne()
	partTwo()
}

func partOne() {
	current := 50
	for _, step := range puzzleInput {
		current = spinDialPartOne(current, step)
		if current == 0 {
			partOneZeroCount++
		}
	}
	fmt.Println("Part 1 zero count:", partOneZeroCount)
}

func partTwo() {
	current := 50
	for _, step := range puzzleInput {
		current = spinDialPartTwo(current, step)
	}
	fmt.Println("Part 2 zero count:", partTwoZeroCount)
}

func wrapValue(value int) int {
	return (value%MAX + MAX) % MAX
}

func spinDialPartOne(current int, value string) int {
	stepsInt, err := strconv.Atoi(value[1:])
	if err != nil {
		log.Fatal(err)
	}

	var newValue int
	switch value[0] {
	case 'L':
		newValue = current - stepsInt
	case 'R':
		newValue = current + stepsInt
	}

	return wrapValue(newValue)
}

func spinDialPartTwo(current int, value string) int {
	stepsInt, err := strconv.Atoi(value[1:])
	if err != nil {
		log.Fatal(err)
	}

	var newValue int
	switch value[0] {
	case 'L':
		newValue = current - stepsInt
		partTwoZeroCount += int(math.Floor(float64(current-1)/float64(MAX))) - int(math.Floor(float64(newValue-1)/float64(MAX)))
	case 'R':
		newValue = current + stepsInt
		partTwoZeroCount += int(math.Floor(float64(newValue)/float64(MAX))) - int(math.Floor(float64(current)/float64(MAX)))
	}

	return wrapValue(newValue)
}
