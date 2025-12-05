package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Day 4 Advent of Code")
	input := GetInput()
	paddedInput := Pad2dArray(input)
	totalRolls := 0
	for {
		accessibleRolls := 0
		for _, line := range(paddedInput) {
			// display grid
			fmt.Println(line)
		}
		
		for i := 1; i < len(paddedInput) - 1; i++ {
			for j := 1; j < len(paddedInput[i]) - 1; j++ {
				accessibleRolls += CheckAdjacents(&paddedInput, i, j)
			}
		}

		totalRolls += accessibleRolls

		fmt.Println("Total rolls:", totalRolls)
		if accessibleRolls <= 0 {
			break;
		}
	}
}

func CheckAdjacents(gridPtr *[][]string, y int, x int) int {
	if (*gridPtr)[y][x] == "@" {
		adjacentRolls := 0
		for i := y - 1; i <= y + 1; i++ {
			for j := x - 1; j <= x + 1; j++ {
				if (*gridPtr)[i][j] == "@" {
					adjacentRolls++
					if adjacentRolls > 4 {
						return 0
					}
				}
			}
		}
		(*gridPtr)[y][x] = "."
		return 1
	}
	return 0
}

func Pad2dArray(toPad [][]string) [][]string {
	var padded [][]string
	topAndBottomRow := make([]string, len(toPad[0]) + 2)
	for i := range(topAndBottomRow) {
		topAndBottomRow[i] = "."
	}
	padded = append(padded, topAndBottomRow)
	for _, row := range(toPad) {
		paddedRow := slices.Insert(row, 0, ".")
		paddedRow = append(paddedRow, ".")
		padded = append(padded, paddedRow)
	}
	padded = append(padded, topAndBottomRow)
	return padded
}

func GetInput() [][]string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputs [][]string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				inputs = append(inputs, strings.Split(strings.TrimRight(line, "\n"), ""))
				break
			} else {
				log.Fatal(err)
				break
			}
		}
		inputs = append(inputs, strings.Split(strings.TrimRight(line, "\n"), ""))
	}
	return inputs
}