package main

import (
	"fmt"
	"os"
)

// menu to select part 1 or part 2
func main() {
	for {
		fmt.Println("===== Advent of Code 2023 Day 02 =====")
		fmt.Println("1. Part 1")
		fmt.Println("2. Part 2")
		fmt.Println("3. Exit")
		fmt.Println("======================================")

		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		switch choice {
		case 1:
			fmt.Println("You chose Option 1.")
			part1()
		case 2:
			fmt.Println("You chose Option 2.")
			part2()
		case 3:
			fmt.Println("Exiting program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
