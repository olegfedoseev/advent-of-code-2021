package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	prev := 0
	increases := 0
	idx := 0
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		if idx > 0 && val > prev {
			increases++
		}
		prev = val
		idx++
	}

	log.Println(increases)
}
