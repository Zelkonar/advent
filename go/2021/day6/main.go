package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseInput() map[int]int {
	data, err := ioutil.ReadFile("../input/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.Split(strings.TrimSpace(string(data)), "\n")[0], ",")
	fishies := make(map[int]int)
	for _, s := range input {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		fishies[num]++
	}
	return fishies
}

func newDay(oldFishies map[int]int) map[int]int {
	newFishies := make(map[int]int)
	for i := 0; i < 9; i++ {
		if i == 0 {
			// special case -> new generation
			newFishies[6], newFishies[8] = oldFishies[0], oldFishies[0]
			continue
		}
		newFishies[i-1] += oldFishies[i]
	}
	return newFishies
}

func main() {
	// Part 1
	fishies := parseInput()
	for i := 0; i < 80; i++ {
		fishies = newDay(fishies)
	}
	totalFish := 0
	for _, count := range fishies {
		totalFish += count
	}
	log.Print(totalFish)

	// Part 2
	fishies = parseInput()
	for i := 0; i < 256; i++ {
		fishies = newDay(fishies)
	}
	totalFish = 0
	for _, count := range fishies {
		totalFish += count
	}
	log.Print(totalFish)
}
