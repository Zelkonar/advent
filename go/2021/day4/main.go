package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var bingos = [][]int{
	// horizontals
	{0, 1, 2, 3, 4},
	{5, 6, 7, 8, 9},
	{10, 11, 12, 13, 14},
	{15, 16, 17, 18, 19},
	{20, 21, 22, 23, 24},

	// verticals
	{0, 5, 10, 15, 20},
	{1, 6, 11, 16, 21},
	{2, 7, 12, 17, 22},
	{3, 8, 13, 18, 23},
	{4, 9, 14, 19, 24},
}

type bingoCell struct {
	val    int
	active bool
}

type bingoBoard struct {
	numbers []bingoCell
	hadBingo   bool
}

func (b *bingoBoard) activateCell(val int) {
	for i, cell := range b.numbers {
		if cell.val == val {
			b.numbers[i].active = true
			return
		}
	}
}

func (b *bingoBoard) hasBingo() bool {
	for _, bingo := range bingos {
		actives := 0
		for _, bingoIdx := range bingo {
			if b.numbers[bingoIdx].active {
				actives++
			}
		}
		if actives == 5 {
			return true // BINGO
		}
	}
	return false
}

func (b *bingoBoard) calculateScore(multiplier int) int {
	sum := 0
	for _, cell := range b.numbers {
		if !cell.active {
			sum += cell.val
		}
	}
	return sum * multiplier
}

func (b *bingoBoard) String() string {
	return fmt.Sprintf("%+v\n%+v\n%+v\n%+v\n%+v", b.numbers[0:4], b.numbers[4:9], b.numbers[9:14], b.numbers[14:19], b.numbers[19:24])
}

func main() {
	data, err := ioutil.ReadFile("../input/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var nums []int
	for _, s := range strings.Split(lines[0], ",") {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	var bingoBoards []bingoBoard
	nextBingoBoard := bingoBoard{}
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		for _, s := range strings.Split(line, " ") {
			if s == "" {
				continue
			}
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			nextBingoBoard.numbers = append(nextBingoBoard.numbers, bingoCell{val: num})
		}
		if len(nextBingoBoard.numbers) == 25 {
			bingoBoards = append(bingoBoards, nextBingoBoard)
			nextBingoBoard = bingoBoard{}
		}
	}

	if nums == nil {
		panic("nums is nil")
	}

	for i := 0; i < 4; i++ {
		for k := range bingoBoards {
			bingoBoards[k].activateCell(nums[i])
		}
	}

	foundBingos := 0
	totalBoards := len(bingoBoards)
	for i := 4; i < len(nums); i++ {
		for k, board := range bingoBoards {
			board.activateCell(nums[i])
			if board.hasBingo() && !board.hadBingo {
				bingoBoards[k].hadBingo = true
				foundBingos++
				if foundBingos == 1 {
					fmt.Printf("part1: %d\n", board.calculateScore(nums[i]))
				}
				if foundBingos == totalBoards {
					fmt.Printf("part2: %d\n", board.calculateScore(nums[i]))
					return
				}
			}
		}
	}
}
