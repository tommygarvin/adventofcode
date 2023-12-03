package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// part1
func part1() {
	// open file, handle errors, close file after return
	file, err := os.Open("../input/02.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// create scanner
	scanner := bufio.NewScanner(file)

	// total
	total := 0

	// loop through scanner results
	for scanner.Scan() {

		// cleanup line from the file
		line := strings.TrimSpace(scanner.Text())
		//fmt.Println(line)

		// get substring before ":", extract number
		game := strings.Split(line, ":")[0]
		gameNumber := extractNumber(game)
		fmt.Println("Game number is:", gameNumber)

		// get substring after ":", break down into rounds
		gameResults := strings.Split(line, ":")[1]
		//fmt.Println("Game results are:", gameResults)
		rounds := strings.Split(gameResults, ";")
		//fmt.Println("Game rounds are:", rounds)

		// empty variables to hold highest cube colors per game
		cubeHighestRed := 0
		cubeHighestGreen := 0
		cubeHighestBlue := 0

		// loop through rounds to find highest cube colors
		for _, round := range rounds {

			// split rounds into cubes
			cubes := strings.Split(round, ",")
			//fmt.Println("Cubes for this round are:", cubes)

			// process cubes by color
			for _, cube := range cubes {
				if strings.Contains(cube, "red") {
					//fmt.Println("Found red cube:", cube)
					cubeNumber := extractNumber(cube)
					//fmt.Println("Cube number is:", cubeNumber)
					//cubeColor := extractColor(cube)
					//fmt.Println("Cube number is:", cubeColor)
					if cubeNumber > cubeHighestRed {
						cubeHighestRed = cubeNumber
					}
				}
				if strings.Contains(cube, "green") {
					//fmt.Println("Found green cube:", cube)
					cubeNumber := extractNumber(cube)
					//fmt.Println("Cube number is:", cubeNumber)
					//cubeColor := extractColor(cube)
					//fmt.Println("Cube number is:", cubeColor)
					if cubeNumber > cubeHighestGreen {
						cubeHighestGreen = cubeNumber
					}
				}
				if strings.Contains(cube, "blue") {
					//fmt.Println("Found blue cube:", cube)
					cubeNumber := extractNumber(cube)
					//fmt.Println("Cube number is:", cubeNumber)
					//cubeColor := extractColor(cube)
					//fmt.Println("Cube number is:", cubeColor)
					if cubeNumber > cubeHighestBlue {
						cubeHighestBlue = cubeNumber
					}
				}
			}
		}
		//fmt.Println("Highest red:", cubeHighestRed)
		//fmt.Println("Highest green:", cubeHighestGreen)
		//fmt.Println("Highest blue:", cubeHighestBlue)

		possible := isGamePossible(cubeHighestRed, cubeHighestGreen, cubeHighestBlue)
		fmt.Println("Is game possible:", possible)

		if possible {
			fmt.Println("The game is possible, adding to total")
			total += gameNumber
		}
	}
	fmt.Println("The final total is:", total)
}

func extractNumber(s string) int {
	pattern := regexp.MustCompile(`\d+`)
	match := pattern.FindString(s)
	number, _ := strconv.Atoi(match)
	return number
}

func extractColor(s string) string {
	pattern := regexp.MustCompile("red|green|blue")
	color := pattern.FindString(s)
	return color
}

func isGamePossible(cubeHighestRed int, cubeHighestGreen int, cubeHighestBlue int) bool {
	// set max cubes
	cubeMaxRed := 12
	cubeMaxGreen := 13
	cubeMaxBlue := 14

	// check if maximum is exceeded
	if cubeMaxRed >= cubeHighestRed && cubeMaxGreen >= cubeHighestGreen && cubeMaxBlue >= cubeHighestBlue {
		return true
	}
	return false
}
