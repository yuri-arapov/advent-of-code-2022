package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("advent of code 2022")
	fmt.Println("day 03")
	fmt.Println("task 2")

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	var lines [3]string
	count := 0

	priorities := 0
	linesCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		linesCount++
		if line == "" {
			continue
		}
		lines[count] = line
		count++
		if count == 3 {
			priorities += itemPriority(commonItem(lines))
			count = 0
		}
	}
	fmt.Println(linesCount, "lines read")
	fmt.Println(priorities)
}

func commonItem(rucksacks [3]string) rune {
	for _, r1 := range []rune(rucksacks[0]) {
		for _, r2 := range []rune(rucksacks[1]) {
			if r1 == r2 {
				for _, r3 := range []rune(rucksacks[2]) {
					if r1 == r3 {
						return r1
					}
				}
			}
		}
	}
	panic("no common item!")
}

func itemPriority(r rune) int {
	if 'a' <= r && r <= 'z' {
		return int(r) - int('a') + 1
	}
	if 'A' <= r && r <= 'Z' {
		return int(r) - int('A') + 27
	}
	panic("bad item type")
}
