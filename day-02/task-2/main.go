package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Rock     = "A"
	Paper    = "B"
	Scissors = "C"

	Loss = "X"
	Draw = "Y"
	Win  = "Z"

	RockScore     = 1
	PaperScore    = 2
	ScissorsScore = 3

	LossScore = 0
	DrawScore = 3
	WinScore  = 6
)

type Round struct {
	OpponentHand string
	Result       string
}

func main() {
	fmt.Println("advent of code 2022")
	fmt.Println("day 02")
	fmt.Println("task 2")

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	game := make([]Round, 0)

	linesCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		hands := strings.Split(line, " ")
		r := Round{OpponentHand: hands[0], Result: hands[1]}
		game = append(game, r)
		linesCount++
	}
	fmt.Println(linesCount, "lines read")

	// Test data.
	// game = []Round{
	// 	{"A", "Y"},
	// 	{"B", "X"},
	// 	{"C", "Z"},
	// }

	fmt.Println(game)
	fmt.Println("game size", len(game))

	score := playGame(game)

	fmt.Println("score", score)
}

func yourHandFromResult(r Round) string {
	if r.Result == Draw {
		return r.OpponentHand
	}
	if r.Result == Loss {
		switch r.OpponentHand {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
	}
	if r.Result == Win {
		switch r.OpponentHand {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	}
	panic(fmt.Sprintf("bad result: %s", r.Result))
}

func playGame(game []Round) (score int) {
	for _, r := range game {
		score += play(r)
	}
	return score
}

func play(r Round) (score int) {
	switch r.Result {
	case Loss:
		return LossScore + handScore(yourHandFromResult(r))
	case Draw:
		return DrawScore + handScore(yourHandFromResult(r))
	case Win:
		return WinScore + handScore(yourHandFromResult(r))
	}
	panic(fmt.Sprintf("bad result: %s", r.Result))
}

func handScore(hand string) int {
	switch hand {
	case Rock:
		return RockScore
	case Paper:
		return PaperScore
	case Scissors:
		return ScissorsScore
	}
	panic(fmt.Sprintf("bad hand: %s", hand))
}
