package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	win  = 6
	draw = 3
	loss = 0
)

func main() {
	lines := fileToLines("input.txt")
	// fmt.Printf("%v\n", lines)
	scores := linesToScoresPt2(lines)
	// fmt.Printf("%v\n", scores)
	total := getSumFromSliceofInt(scores)
	fmt.Print("Total: ", total, "\n")
}

func fileToLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return
}

func linesToScoresPt1(lines []string) (scores []int) {
	total := 0
	for _, line := range lines {
		match := strings.Split(line, " ")
		me := match[1]
		you := match[0]

		// Start with a loss
		score := 0
		// A = X = 1 rock
		// B = Y = 2 paper
		// C = Z = 3 scissors
		if (me == "X" && you == "C") || (me == "Y" && you == "A") || (me == "Z" && you == "B") { // win
			score = win
			fmt.Printf("with %s vs %s we got a win  and scored %d + %d", me, you, lettersToScore(me), win)
		} else if (me == "X" && you == "A") || (me == "Y" && you == "B") || (me == "Z" && you == "C") { //draw
			score = draw
			fmt.Printf("with %s vs %s we got a draw and scored %d + %d", me, you, lettersToScore(me), draw)
		} else {
			fmt.Printf("with %s vs %s we got a loss and scored %d + %d", me, you, lettersToScore(me), loss)
		}
		// add the score of the played hand
		score += lettersToScore(me)
		total += score
		fmt.Printf(" total : %d\n", total)
		// chuck in the slice
		scores = append(scores, score)

	}
	return
}

func lettersToScore(letter string) int {
	switch letter {
	case "A":
	case "X":
		return 1
	case "B":
	case "Y":
		return 2
	case "C":
	case "Z":
		return 3
	}
	return 0
}

func linesToScoresPt2(lines []string) (scores []int) {
	for _, line := range lines {
		match := strings.Split(line, " ")
		me := match[1]
		you := match[0]
		score := 0
		// A = 1 rock
		// B = 2 paper
		// C = 3 scissors
		switch me {
		case "X": // loss
			score += loss
			switch you {
			case "A": // they played rock
				score += 3 // we play scissors

			case "B": // they played paper
				score += 1 // we play rock

			case "C": // they played scissors
				score += 2 // we play paper
			}
		case "Y": // draw
			score += draw
			switch you {
			case "A": // they played rock
				score += 1 // we play rock

			case "B": // they played paper
				score += 2 // we play paper

			case "C": // they played scissors
				score += 3 // we play scissors
			}

		case "Z": // win
			score += win
			switch you {
			case "A": // they played rock
				score += 2 // we play paper

			case "B": // they played paper
				score += 3 // we play scissors

			case "C": // they played scissors
				score += 1 // we play rock
			}
		}

		// chuck in the slice
		scores = append(scores, score)

	}
	return
}

func getSumFromSliceofInt(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}
