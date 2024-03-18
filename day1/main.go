package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	mostCalories := []int{0, 0, 0}

	scanner := bufio.NewScanner(f)

	var currentElfCalories = 0
	for scanner.Scan() {

		text := scanner.Text()
		// string to int

		if text == "" {

			if mostCalories[0] < currentElfCalories {
				mostCalories[0] = currentElfCalories
				sort.Ints(mostCalories)
				fmt.Printf("set most calories to: %d\n", mostCalories)
			}
			currentElfCalories = 0
		} else {
			fmt.Printf("current Food Calories are: %s\n", text)
			currentFoodCalories, err := strconv.Atoi(text)
			if err != nil {
				// ... handle error
				panic(err)
			}
			currentElfCalories += currentFoodCalories
			fmt.Printf("current Elf Calories are: %d\n", currentElfCalories)
		}

	}

	fmt.Printf("most Calories one Elf is carrying are: %d\n %d\n %d\n", mostCalories[0], mostCalories[1], mostCalories[2])
	print(mostCalories[0] + mostCalories[1] + mostCalories[2])

}
