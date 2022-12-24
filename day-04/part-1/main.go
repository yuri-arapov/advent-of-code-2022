package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("advent of code 2022")
	fmt.Println("day 04")
	fmt.Println("task 1")

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	linesCount := 0
	inclusions := 0
	overlaps := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			commaSplit := strings.Split(line, ",")
			r1 := parseRange(commaSplit[0])
			r2 := parseRange(commaSplit[1])
			fmt.Println(r1, r2)
			if r1.containsRange(r2) || r2.containsRange(r1) {
				inclusions++
			}
			if r1.overlaps(r2) || r2.overlaps(r1) {
				overlaps++
			}
		}
		linesCount++
	}
	fmt.Println(linesCount, "lines read")
	fmt.Println(inclusions, "inclusions")
	fmt.Println(overlaps, "overlaps")
}

type Range struct {
	from, to int
}

func parseRange(s string) Range {
	dashSplit := strings.Split(s, "-")
	return Range{
		from: toInt(dashSplit[0]),
		to:   toInt(dashSplit[1]),
	}
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

// contains returns true if range 'r' contains value 'i'.
func (r Range) contains(i int) bool {
	return r.from <= i && i <= r.to
}

// containsRange returns true if range 'r' contains range 'r2'.
func (r Range) containsRange(r2 Range) bool {
	return r.contains(r2.from) && r.contains(r2.to)
}

func (r Range) overlaps(r2 Range) bool {
	return r.containsRange(Range{r2.from, r2.from}) || r.containsRange(Range{r2.to, r2.to})
}

