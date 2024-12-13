package main

import (
	"fmt"
	"os"
)


const (
	DEC_POS = 10000
	LIST_LEN = 1000
)


func main() {
	fmt.Println("Day 1, task 2")

	// let's try to be clever
	in, _ := os.ReadFile("./input") // No error handling! Gulp!
	var values [LIST_LEN]int
	countMap := make(map[int]bool, LIST_LEN)
	similarity := 0
	decPos := DEC_POS
	j := 0
	thisNum := 0
	for i := 0; i < len(in); i++ {
		// Extract integer value from ASCII value (0 is ASCII value 48, 1 is 49...)
 		intVal := in[i] - 48

		// if byte represents an int char
		if intVal < 10 {
			thisNum += int(intVal) * decPos
			decPos /= 10
		} else {
			decPos = DEC_POS
			if in[i] == 10 {
				// we're at the end of the line, thisNum is the value
				values[j] = thisNum
				j++
			} else {
				// we're halfway through the line, thisNum is key,
				// let's skip the next whitespaces
				i += 2
				countMap[thisNum] = true
			}
			thisNum = 0
		}
	}

	for l := 0; l < LIST_LEN; l++ {
		value := values[l]
		if _, exists := countMap[value]; exists {
			similarity += value
		}
	}

	fmt.Println("Similarity is", similarity)
}
