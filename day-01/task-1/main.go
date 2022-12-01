package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("advent of code 2022")
	fmt.Println("day 01")
	fmt.Println("task 1")

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	linesCount := 0
	elvesCount := 0
	var m1, m2, m3 int
	currentElfCalories := 0

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		linesCount++
		if line != "" {
			currentElfCalories += stringToInt(line)
		} else {
			m1, m2, m3 = max3(m1, m2, m3, currentElfCalories)
			currentElfCalories = 0
			elvesCount++
		}
	}

	m1, m2, m3 = max3(m1, m2, m3, currentElfCalories)

	fmt.Println("number of lines:", linesCount)
	fmt.Println("number of elves:", elvesCount)
	fmt.Println("max calories:", m1, m2, m3)
	fmt.Println(m1 + m2 + m3)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func stringToInt(s string) int {
	if s == "" {
		return 0
	}
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func max3(i1, i2, i3, v int) (int, int, int) {
	if v > i1 {
		return v, i1, i2
	}
	if v > i2 {
		return i1, v, i2
	}
	if v > i3 {
		return i1, i2, v
	}
	return i1, i2, i3
}
