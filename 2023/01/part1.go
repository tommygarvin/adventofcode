package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {

	// open file, handle errors, close file after return
	file, err := os.Open("../input/01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create scanner
	scanner := bufio.NewScanner(file)

	// empty array of integers to hold calibration results
	calibrations := []int{}

	// loop through scanner results
	for scanner.Scan() {

		// cleanup line from the file
		line := strings.TrimSpace(scanner.Text())

		// declare empty array of integers to hold numbers from lines
		numbers := []int{}

		// for each character in the line
		for _, char := range line {
			// attempt to convert string to int
			val, err := strconv.Atoi(string(char))
			// if error then it cannot be converted to int, continue
			// if no error it can be converted to int, append to array
			if err != nil {
				continue
			} else {
				numbers = append(numbers, val)
			}
		}

		numberFirst := strconv.Itoa(numbers[0])                        // first number from array converted to string
		numberLast := strconv.Itoa(numbers[len(numbers)-1])            // lst number from array converted to string
		calibrationStr := fmt.Sprintf("%s%s", numberFirst, numberLast) // concat via string interpolation
		//println(calibrationStr)

		calibrationInt, _ := strconv.Atoi(calibrationStr)   // convert calibration result to int
		calibrations = append(calibrations, calibrationInt) // add calibration result to array of calibrations
		//println(calibrationInt)

		// total up the calibrations
		calibrationsSum := 0
		for _, value := range calibrations {
			calibrationsSum += value
		}
		fmt.Println("Running total:", calibrationsSum)
	}
}
