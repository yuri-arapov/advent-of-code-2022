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
			r1, r2 := firstTwo(ymap(strings.Split(line, ","), parseRange))
			fmt.Println(r1, r2)
			if r1.includes(r2) || r2.includes(r1) {
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

func makeRange(from, to int) Range {
	return Range{
		from: from,
		to:   to,
	}
}

func parseRange(s string) Range {
	return makeRange(firstTwo(ymap(strings.Split(s, "-"), toInt)))
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

// includes returns true if range 'r' contains both ends of range 'r2'.
func (r Range) includes(r2 Range) bool {
	return r.contains(r2.from) && r.contains(r2.to)
}

// overlaps returns true if range 'r' overlaps with range 'r2'.
func (r Range) overlaps(r2 Range) bool {
	return r.contains(r2.from) || r.contains(r2.to)
}

// ymap applies function 'f' to every element of array 'args' of type 'T1' and
// returns array of type 'T2'.
func ymap[T1 any, T2 any](args []T1, f func(a T1) T2) []T2 {
	res := make([]T2, len(args))
	for i := range args {
		res[i] = f(args[i])
	}
	return res
}

// firstTwo returns first two elements of array 'args'.
func firstTwo[T any](args []T) (T, T) {
	if len(args) < 2 {
		panic("at least 2 elements in array expected")
	}
	return args[0], args[1]
}
