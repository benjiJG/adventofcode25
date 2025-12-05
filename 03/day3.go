package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 3 Advent of Code")
	banks := GetInput()
	PartOne(&banks)
	PartTwo(&banks)
}

func PartTwo(banksPtr *[][]string) {
	maxJoltage := 0
	numBatteries := 12

	for _, bank := range(*banksPtr) {
		startCursor := 0
		endCursor := len(bank) - (numBatteries - 1)
		var joltageString strings.Builder
		for i := numBatteries - 1; i >= 0; i-- {
			selectionGroup := bank[startCursor:endCursor]
			max := slices.Max(selectionGroup)
			maxIdx := slices.Index(selectionGroup, max)
			startCursor += maxIdx + 1
			endCursor = len(bank) - (i - 1)
			joltageString.WriteString(max)
		}
		joltage, err := strconv.Atoi(joltageString.String())
		if err != nil {
			log.Fatal(err)
		}
		maxJoltage += joltage
	}

	fmt.Println("Part two joltage:", maxJoltage)
}

func PartOne(banksPtr *[][]string) {
	maxJoltage := 0

	for _, i := range(*banksPtr) {
		bankMinusLast := i[:len(i)-1]
		firstMax := slices.Max(bankMinusLast)
		firstIdx := slices.Index(bankMinusLast, firstMax)

		bankMinusInvalid := i[firstIdx + 1:]
		lastMax := slices.Max(bankMinusInvalid)

		joltage, err := strconv.Atoi(firstMax + lastMax)
		if err != nil {
			log.Fatal(err)
		}
		maxJoltage += joltage
	}

	fmt.Println("Part one joltage:", maxJoltage)
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