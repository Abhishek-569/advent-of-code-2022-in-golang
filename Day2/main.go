package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalScore := 0
	for scanner.Scan() {
		// put the value of the line into a variable
		line := scanner.Text()
		// Line is of form "opponent me" - split on the space
		words := strings.Split(line, " ")
		me := words[1]
		opponent := words[0]

		// A/X =rock
		// B/Y =paper
		// C/Z =scissors
		// rock beats scissors
		// paper beats rock
		// scissors beats paper
		// if I win add 6 to total score
		// if I lose add 0 to total score
		// if I tie add 3 to total score
		// Who won? Opponent A beats me Z, B beats me X, C beats me Y

		me = changeMe(opponent, me)
		totalScore += calculateScore(opponent, me)
		//add an extra score based on what i picked
		// X=1, Y=2, Z=3
		if me == "X" {
			totalScore += 1
		} else if me == "Y" {
			totalScore += 2
		} else if me == "Z" {
			totalScore += 3
		}

	}
	fmt.Println(totalScore)
}

// create a helper function which changes the value of me
// if my value is X change my value to lose opponent
// if my value is Y make me draw opponent
// if my value is Z make me win opponent
func changeMe(opponent, me string) string {
	if me == "X" && opponent == "A" {
		return "Z"
	} else if me == "Y" && opponent == "A" {
		return "X"
	} else if me == "Z" && opponent == "A" {
		return "Y"
	} else if me == "X" && opponent == "C" {
		return "Y"
	} else if me == "Y" && opponent == "C" {
		return "Z"
	} else if me == "Z" && opponent == "C" {
		return "X"
	}
	return me
}

func calculateScore(opponent string, me string) int {
	if opponent == "A" && me == "Y" {
		return 6
	} else if opponent == "B" && me == "Z" {
		return 6
	} else if opponent == "C" && me == "X" {
		return 6
	} else if opponent == "A" && me == "X" {
		return 3
	} else if opponent == "B" && me == "Y" {
		return 3
	} else if opponent == "C" && me == "Z" {
		return 3
	}
	return 0
}
