package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2 Advent of Code")
	codes := GetInput()
	pt1Sum := 0
	pt2Sum := 0
	for _, v := range(codes) {
		CheckRange(v, &pt1Sum, 1)
		CheckRange(v, &pt2Sum, 2)
	}
	fmt.Println("Part 1:", pt1Sum)
	fmt.Println("Part 2:", pt2Sum)
}

func CheckRange(codeRange string, sumPtr *int, taskPart uint8) {
	startEnd := strings.Split(codeRange, "-")
	start, err := strconv.Atoi(startEnd[0])
	if err != nil {
		log.Fatal(err)
	}
	end, err := strconv.Atoi(startEnd[1])
	if err != nil {
		log.Fatal(err)
	}
	switch taskPart {
		case 1:
			for i := start; i <= end; i++ {
				id := strconv.Itoa(i)
				mid := len(id) / 2
				firstSeq := id[:mid]
				secondSeq := id[mid:]
				if firstSeq == secondSeq {
					*sumPtr += i
				}
			}
		case 2:
			for i := start; i <= end; i++ {
				if IsInvalid(strconv.Itoa(i)) {
					*sumPtr += i
				}
			}
	}
}

func IsInvalid(id string) bool {
	idLength := len(id)
	for i := idLength / 2; i > 0; i-- {
		if idLength % i == 0 {
			substring := id[:i]
			repeatedSubstring := strings.Repeat(substring, idLength / i)
			if id == repeatedSubstring {
				return true
			}
		}
	}
	return false
}

func GetInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputs []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString(',')
		if err != nil {
			if err == io.EOF {
				inputs = append(inputs, strings.TrimRight(line, ","))
				break
			} else {
				log.Fatal(err)
				break
			}
		}
		inputs = append(inputs, strings.TrimRight(line, ","))
	}
	return inputs
}