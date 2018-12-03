package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	frequencies := make([]int, 0)
	seenFrequencies := make(map[int]bool)
	seenFrequencies[0] = true

	sum := 0
	scanner := bufio.NewScanner(file)

	foundDoubleFreq := false

	for scanner.Scan() {
		text := scanner.Text()
		value, err := strconv.Atoi(text)
		if err != nil {
			log.Fatal(err)
		}
		frequencies = append(frequencies, value)
		sum += value
		if seenFrequencies[sum] {
			fmt.Printf("First frequency appearing twice: %v\n", sum)
			foundDoubleFreq = true
			break
		} else {
			seenFrequencies[sum] = true
		}
	}

	if !foundDoubleFreq {
		frequenciesCount := len(frequencies)
		counter := 0
		for true {
			sum += frequencies[counter]
			if seenFrequencies[sum] {
				fmt.Printf("First frequency appearing twice: %v\n", sum)
				break
			} else {
				seenFrequencies[sum] = true
			}
			counter++
			counter %= frequenciesCount
		}
	}
}
