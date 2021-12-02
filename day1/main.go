package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var prev [3]int
	increases := 0
	idx := 0
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		// load up first window
		if idx < 3 {
			prev[idx] = val
			idx++
			continue
		}

		if prev[0]+prev[1]+prev[2] < prev[1]+prev[2]+val {
			increases++
		}

		prev[0] = prev[1]
		prev[1] = prev[2]
		prev[2] = val
		idx++
	}

	log.Println(increases)
}
