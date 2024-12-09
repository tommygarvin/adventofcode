package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	//log.Printf("Input: %s", input)
	part1(input)
	part2(input)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func getInput() string {
	file, err := os.Open("../input/03.txt")
	handleError(err)
	defer file.Close()

	var input string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		input += line
	}

	return input
}

func part1(input string) {
	log.Printf("Input: %s", input)
	// find all instances and store in slice
	filter := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := filter.FindAllString(input, -1)
	log.Printf("Matches: %s", matches)

	// a slice of the products
	var products []int
	for _, match := range matches {
		log.Printf("Processing: %s", match)
		// split the match into two numbers
		numbers := filter.FindStringSubmatch(match)
		//log.Printf("Numbers: %s", numbers)
		// convert the numbers to integers
		firstNumber, err := strconv.Atoi(numbers[1])
		handleError(err)
		log.Printf("First Number: %d", firstNumber)
		secondNumber, err := strconv.Atoi(numbers[2])
		handleError(err)
		log.Printf("Second Number: %d", secondNumber)
		// multiply the two numbers and store in products slice
		product := firstNumber * secondNumber
		products = append(products, product)
	}
	log.Printf("Products: %v", products)

	// sum the products
	sum := 0
	for _, product := range products {
		sum += product
	}
	log.Printf("Sum: %d", sum)
}

func part2(input string) {
	log.Printf("Input: %s", input)
	// find all instances and store in slice
	filterAll := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)
	filterMultiply := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	filterDigits := regexp.MustCompile(`\d+`)
	filterDo := regexp.MustCompile(`do\(\)`)
	filterDont := regexp.MustCompile(`don't\(\)`)
	matches := filterAll.FindAllString(input, -1)
	log.Printf("Matches: %s", matches)

	// a slice of the products
	var products []int

	isInstructionOn := true
	log.Printf("Instruction is on")
	for _, match := range matches {
		log.Printf("Processing: %s", match)

		// check if the match is do() or don't()
		if filterDo.MatchString(match) {
			isInstructionOn = true
			log.Printf("Instruction is on")
		} else if filterDont.MatchString(match) {
			isInstructionOn = false
			log.Printf("Instruction is off")
		}

		if isInstructionOn {
			if filterMultiply.MatchString(match) {
				// extract the two numbers
				numbers := filterDigits.FindAllString(match, -1)

				log.Printf("Numbers: %s", numbers)
				// convert the numbers to integers
				firstNumber, err := strconv.Atoi(numbers[0])
				handleError(err)
				log.Printf("First Number: %d", firstNumber)
				secondNumber, err := strconv.Atoi(numbers[1])
				handleError(err)
				log.Printf("Second Number: %d", secondNumber)
				// multiply the two numbers and store in products slice
				product := firstNumber * secondNumber
				products = append(products, product)
			}
		}
	}
	log.Printf("Products: %v", products)

	// sum the products
	sum := 0
	for _, product := range products {
		sum += product
	}
	log.Printf("Sum: %d", sum)
}
