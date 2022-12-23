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
	fmt.Println("task 1")

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	priorities := 0
	linesCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			priorities += rucksackPriority(line)
		}
		linesCount++
	}
	fmt.Println(linesCount, "lines read")
	fmt.Println(priorities)
}

func rucksackPriority(r string) int {
	rr := []rune(r)
	for _, r1 := range rr[0 : len(rr)/2] {
		for _, r2 := range rr[len(rr)/2 : len(rr)] {
			if r1 == r2 {
				return itemPriority(r1)
			}
		}
	}
	panic(fmt.Sprintf("no items shared in rucksack: %s", r))
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
