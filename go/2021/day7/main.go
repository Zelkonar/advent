package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseInput() (positionToCount map[int]int, max int) {
	data, err := ioutil.ReadFile("../input/day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(strings.Split(strings.TrimSpace(string(data)), "\n")[0], ",")

	crabSubs := make(map[int]int)
	for _, s := range input {
		position, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		crabSubs[position]++
		if max < position {
			max = position
		}
	}
	return crabSubs, max
}

func findMinimumFuelPosition(crabSubs map[int]int, min, max int, fuelCost func(map[int]int, int) int) int {
	// Binary search approach, need to check fuel cost of min, compare against 1 lower or higher
	halfway := min + (max-min)/2
	minFuelCost := fuelCost(crabSubs, halfway)
	if minFuelCost > fuelCost(crabSubs, halfway-1) {
		return findMinimumFuelPosition(crabSubs, min, halfway, fuelCost)
	} else if minFuelCost > fuelCost(crabSubs, halfway+1) {
		return findMinimumFuelPosition(crabSubs, halfway, max, fuelCost)
	}
	return halfway
}

func calculateFuel(crabSubs map[int]int, position int) int {
	total := 0
	for crabPosition, count := range crabSubs {
		fuelCost := (crabPosition - position) * count
		if fuelCost < 0 {
			fuelCost *= -1
		}
		total += fuelCost
	}
	return total
}

func calculateFuelP2(crabSubs map[int]int, position int) int {
	total := 0
	for crabPosition, count := range crabSubs {
		distance := crabPosition - position
		if distance < 0 {
			distance *= -1
		}
		fuelCost := sumOfNaturalNumbers(distance) * count
		total += fuelCost
	}
	return total
}

func sumOfNaturalNumbers(n int) int { return (n * (n + 1)) / 2 }

func main() {
	crabSubs, max := parseInput()
	minPosition := findMinimumFuelPosition(crabSubs, 0, max, calculateFuel)
	fuelCost := calculateFuel(crabSubs, minPosition)
	log.Print(fuelCost)

	minPositionP2 := findMinimumFuelPosition(crabSubs, 0, max, calculateFuelP2)
	fuelCostP2 := calculateFuelP2(crabSubs, minPositionP2)
	log.Print(fuelCostP2)
}
