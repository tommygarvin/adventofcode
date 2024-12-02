package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	//log.Println(input)
	part1(input)
	//part2(input)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func getInput() [][]int {
	file, err := os.Open("../input/02.txt")
	handleError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var slice2D [][]int
	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Fields(report)
		row := make([]int, len(levels))
		for i, field := range levels {
			num, err := strconv.Atoi(field)
			handleError(err)
			row[i] = num
		}
		slice2D = append(slice2D, row)
	}

	return slice2D
}

func isSortedDescending(slice []int) bool {
	reversed := make([]int, len(slice))
	copy(reversed, slice)
	slices.Reverse(reversed)
	return slices.IsSorted(reversed)
}

func isLevelDifferenceWithinTolerance(slice []int) bool {
	//log.Println(slice)
	//var levelDifferences []int
	for i := 0; i < len(slice)-1; i++ {
		difference := slice[i] - slice[i+1]
		if difference < 0 {
			difference = -difference
		}
		//levelDifferences = append(levelDifferences, difference)
		if difference < 1 || difference > 3 {
			return false
		}
	}
	//log.Println(levelDifferences)
	return true
}

func part1(input [][]int) {
	var isSafe []bool
	for _, row := range input {
		//log.Println(slices.IsSorted(row))
		//log.Println(isSortedDescending(row))
		//log.Println(isLevelDifferenceWithinTolerance(row))
		if (slices.IsSorted(row) || isSortedDescending(row)) && isLevelDifferenceWithinTolerance(row) {
			isSafe = append(isSafe, true)
		} else {
			isSafe = append(isSafe, false)
		}
	}
	//log.Println(isSafe)

	countSafe := 0
	for _, value := range isSafe {
		if value {
			countSafe++
		}
	}
	log.Println(countSafe)
	err := os.WriteFile("../output/02a.txt", []byte(strconv.Itoa(countSafe)), 0644)
	handleError(err)
}
