package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	column0, column1 := getSlices()
	part1(column0, column1)
	part2(column0, column1)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func getSlices() ([]int, []int) {
	file, err := os.Open("../input/01-test.txt")
	handleError(err)
	input := bufio.NewScanner(file)
	var column0 []int
	var column1 []int

	for input.Scan() {
		line := input.Text()
		splitString := strings.Split(line, "   ")
		value0, err := strconv.Atoi(strings.Trim(splitString[0], " "))
		handleError(err)
		value1, err := strconv.Atoi(strings.Trim(splitString[1], " "))
		handleError(err)
		column0 = append(column0, value0)
		column1 = append(column1, value1)
	}

	return column0, column1
}

func part1(column0 []int, column1 []int) {
	sort.Ints(column0)
	sort.Ints(column1)

	log.Println(column0)
	log.Println(column1)

	var distanceApart []int
	for i := 0; i < len(column0); i++ {
		distance := column1[i] - column0[i]
		if distance < 0 {
			distance = -distance
		}
		distanceApart = append(distanceApart, distance)
	}

	log.Println(distanceApart)

	distanceTotal := 0
	for _, distance := range distanceApart {
		distanceTotal += distance
	}

	log.Println(distanceTotal)
	err := os.WriteFile("../output/01a.txt", []byte(strconv.Itoa(distanceTotal)), 0644)
	handleError(err)
}

func part2(column0 []int, column1 []int) {

	log.Println(column0)
	log.Println(column1)

	var similarityScore []int
	for i := 0; i < len(column0); i++ {
		score := 0
		for j := 0; j < len(column1); j++ {
			if column0[i] == column1[j] {
				score++
			}
		}
		similarityScore = append(similarityScore, column0[i]*score)
	}

	log.Println(similarityScore)

	scoreTotal := 0
	for _, score := range similarityScore {
		scoreTotal += score
	}

	log.Println(scoreTotal)
	err := os.WriteFile("../output/01b.txt", []byte(strconv.Itoa(scoreTotal)), 0644)
	handleError(err)
}
