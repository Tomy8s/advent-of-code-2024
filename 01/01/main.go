package main

import (
	"fmt"
	"os"
	"sort"
)

const (
	DEC_POS = 10000
	LIST_LEN = 1000
)

func main () {
	// Let's try to be "clever"
	in, _ := os.ReadFile("./input") // No error handling! Gulp!
	ints := [2][]int{make([]int, LIST_LEN), make([]int, LIST_LEN)}
	diff := 0
	listIdx := 0
	decPos := DEC_POS
	j := 0
	for i := 0; i < len(in); i++ {
		// Extract integer value from ASCII value (0 is ASCII value 48, 1 is 49...)
 		intVal := in[i] - 48

		// if byte represents an int char
		if intVal < 10 {
			ints[listIdx][j] += int(intVal) * decPos
			decPos /= 10
		} else {
			decPos = DEC_POS
			listIdx = listIdx ^ 1
			if in[i] == 10 {
				j++
			}
		}
	}
	// This is taking to long, just do it the easy way!!
	sort.Ints(ints[0])
	sort.Ints(ints[1])
	for k := 0; k < LIST_LEN; k++ {
		if ints[0][k] > ints[1][k] {
			diff += ints[0][k] - ints[1][k]
		} else {
			diff += ints[1][k] - ints[0][k]
		}
	}
	fmt.Println(diff)
}