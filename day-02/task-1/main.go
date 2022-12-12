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

	RockScore     = 1
	PaperScore    = 2
	ScissorsScore = 3

	LossScore = 0
	DrawScore = 3
	WinScore  = 6
)

type Round struct {
	OpponentHand string
	YourHand     string
}

func main() {
	fmt.Println("advent of code 2022")
	fmt.Println("day 02")
	fmt.Println("task 1")

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
		r := Round{OpponentHand: hands[0], YourHand: hands[1]}
		game = append(game, r)
		linesCount++
	}
	fmt.Println(linesCount, "lines read")

	yourHandMaps := []map[string]string{
		{"X": Rock, "Y": Paper, "Z": Scissors},
		// The only first mapping needed for this task.
		// {"X": Rock, "Y": Scissors, "Z": Paper},
		// {"X": Paper, "Y": Rock, "Z": Scissors},
		// {"X": Paper, "Y": Scissors, "Z": Rock},
		// {"X": Scissors, "Y": Paper, "Z": Rock},
		// {"X": Scissors, "Y": Rock, "Z": Paper},
	}

	// Test data.
	// game = []Round{
	// 	{"A", "Y"},
	// 	{"B", "X"},
	// 	{"C", "Z"},
	// }

	fmt.Println(game)
	fmt.Println("game size", len(game))

	for _, m := range yourHandMaps {
		score := playGame(game, m)
		fmt.Println(score, m)
	}
}

func playGame(game []Round, yourHandMap map[string]string) (score int) {
	for _, r := range game {
		score += play(mapYourHand(r, yourHandMap))
	}
	return score
}

func play(r Round) (score int) {
	if r.OpponentHand == r.YourHand {
		return DrawScore + handScore(r.YourHand)
	}
	if isWin(r.OpponentHand, r.YourHand) {
		return WinScore + handScore(r.YourHand)
	}
	return LossScore + handScore(r.YourHand)
}

func isWin(opHand, yourHand string) bool {
	if opHand == Rock && yourHand == Paper {
		return true
	}
	if opHand == Paper && yourHand == Scissors {
		return true
	}
	if opHand == Scissors && yourHand == Rock {
		return true
	}
	return false
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

func mapYourHand(r Round, handMap map[string]string) Round {
	res := r
	res.YourHand = handMap[res.YourHand]
	return res
}
