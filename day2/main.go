package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	horizontalPosition := 0
	depth := 0
	aim := 0
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		direction := tokens[0]
		amount, _ := strconv.Atoi(tokens[1])

		switch direction {
		case "forward":
			horizontalPosition += amount
			depth += aim * amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}

	log.Println(horizontalPosition * depth)
}
