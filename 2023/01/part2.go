package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2() {

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

	// conversion table, had to adjust for bad combos like "twone" being both two and one.
	replacements := map[string]string{
		"zero":  "z0o",
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	// loop through scanner results
	for scanner.Scan() {

		// cleanup line from the file
		line := strings.TrimSpace(scanner.Text())

		// call function to replace substrings
		//println("calling replacement function with", line)
		lineReplaced := replaceSubstrings(line, replacements)

		// declare empty array of integers to hold numbers from lines
		numbers := []int{}

		// for each character in the line
		for _, char := range lineReplaced {
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

// iterate over the input replacing keys with values from the conversion map
func replaceSubstrings(input string, replacementMap map[string]string) string {
	for old, new := range replacementMap {
		input = strings.Replace(input, old, new, -1)
		//println("replacing", old)
		//println("with", new)
		//println("replacement result", input)
	}
	return input
}
