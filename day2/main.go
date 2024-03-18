package main

import (
	"advent_of_code_2022/utils"
	"strings"
)

type ENEMY string

const (
	ENEMY_ROCK     = "A"
	ENEMY_PAPER    = "B"
	ENEMY_SCISSORS = "C"
)

type ME string

const (
	MY_ROCK     = "X"
	MY_PAPER    = "Y"
	MY_SCISSORS = "Z"
)

func main() {

	// How we score

	// ROCK: 1
	// PAPER: 2
	// SCISSORS: 3

	// LOSE: 0
	// DRAW: 3
	// WIN: 6

	var result = 0

	scanner := utils.GetFileScanner("day2/input.txt")
	for scanner.Scan() {
		//one line
		var input string = scanner.Text()

		res := strings.Split(input, " ")
		if len(res) == 2 {

			enemy_val := getVal(res[0])
			my_val := getVal(res[1])

			diff := my_val - enemy_val

			if diff == 1 {
				//we win
				my_val += 6
			} else if diff == 2 {
				//we loose
			} else if diff == -1 {
				//we loose
			} else if diff == -2 {
				//we win
				my_val += 6
			} else if diff == 0 {
				//draw
				my_val += 3
			}
			result += my_val
		}

	}

	print(result)

}

func getVal(s string) int {
	switch s {
	case MY_ROCK:
		return 1
	case MY_PAPER:
		return 2
	case MY_SCISSORS:
		return 3
	case ENEMY_ROCK:
		return 1
	case ENEMY_PAPER:
		return 2
	case ENEMY_SCISSORS:
		return 3

	}
	return 0
}
