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
	fmt.Println("Day 5 Advent of Code")
	input := GetInput()
	var ranges [][2]int
	freshIdsPt1 := 0
	var freshIdsPt2 []int

	for _, i := range input {
		if len(i) == 2 {
			i0, err := strconv.Atoi(i[0])
			i1, err := strconv.Atoi(i[1])
			if err != nil {
				log.Fatal(err)
			}
			intRange := [2]int{i0, i1}
			ranges = append(ranges, intRange)
			CheckFreshPt2(&freshIdsPt2, intRange)
		} else if i[0] != "" {
			freshIdsPt1 += CheckFresh(i, ranges)
		}
	}

	fmt.Println("Fresh pt1:", freshIdsPt1)
	fmt.Println("Fresh pt2:", len(freshIdsPt2))
}

func CheckFresh(idStr []string, idRanges [][2]int) int {
	id, err := strconv.Atoi(idStr[0])
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range idRanges {
		if id >= r[0] && id <= r[1] {
			return 1
		}
	}
	return 0
}

func CheckFreshPt2(ids *[]int, idRange [2]int) {
	for i := idRange[0]; i <= idRange[1]; i++ {
		if !slices.Contains(*ids, i) {
			*ids = append(*ids, i)
		} 
	}
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
				inputs = append(inputs, strings.Split(strings.TrimRight(line, "\n"), "-"))
				break
			} else {
				log.Fatal(err)
				break
			}
		}
		inputs = append(inputs, strings.Split(strings.TrimRight(line, "\n"), "-"))
	}
	return inputs
}