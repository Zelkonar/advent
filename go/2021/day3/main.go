package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// this is absolutely horrendous and embarrassing, needs refactored but we all know I never will
func main() {
	data, err := ioutil.ReadFile("../../../input/day3.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	totalLines := len(lines)
	binarySize := len(lines[0])
	bitTotals := make([]int, binarySize)
	for _, line := range lines {
		for i, bit := range strings.Split(line, "") {
			bitValue, err := strconv.Atoi(bit)
			if err != nil {
				panic(err)
			}
			bitTotals[i] += bitValue
		}
	}

	gammaBinary := ""
	epsilonBinary := ""
	for _, gammaTotal := range bitTotals {
		if gammaTotal >= totalLines/2 {
			gammaBinary += "1"
			epsilonBinary += "0"
		} else {
			gammaBinary += "0"
			epsilonBinary += "1"
		}
	}
	gamma, err := strconv.ParseInt(gammaBinary, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(epsilonBinary, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Print("%d", gamma*epsilon)

	oxyBinary := ""
	oxyLines := lines
	co2Binary := ""
	co2Lines := lines

	fmt.Println("49 is \"1\", 48 is \"0\"")
	for i := 0; i < binarySize; i++ {
		oxyCommonBit := mostCommonBit(i, oxyLines)
		co2CommonBit := leastCommonBit(i, co2Lines)
		if len(oxyLines) != 1 {
			oxyLines = filterListViaBitCriteria(i, oxyCommonBit, oxyLines)
		}
		if len(co2Lines) != 1 {
			co2Lines = filterListViaBitCriteria(i, co2CommonBit, co2Lines)
		}
	}
	oxyBinary = oxyLines[0]
	co2Binary = co2Lines[0]
	oxyRating, err := strconv.ParseInt(oxyBinary, 2, 64)
	if err != nil {
		panic(err)
	}
	co2Rating, err := strconv.ParseInt(co2Binary, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d", oxyRating*co2Rating)
}

func mostCommonBit(idx int, lines []string) uint8 {
	total := 0
	for _, line := range lines {
		if line[idx] == "1"[0] { // lol
			total++
		}
	}
	if total*2 >= len(lines) {
		return "1"[0]
	}
	return "0"[0]
}

func leastCommonBit(idx int, lines []string) uint8 {
	total := 0
	for _, line := range lines {
		if line[idx] == "1"[0] { // lol
			total++
		}
	}
	if total*2 < len(lines) {
		return "1"[0]
	}
	return "0"[0]
}

func filterListViaBitCriteria(idx int, bit uint8, lines []string) []string {
	res := make([]string, 0)
	for _, line := range lines {
		if line[idx] == bit {
			res = append(res, line)
		}
	}
	return res
}
