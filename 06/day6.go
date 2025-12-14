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
	input := GetInput()
	
	// PartOne(input)
	PartTwo(input)
}

func PartTwo(input [][]string) {
	for _, v := range input {
		fmt.Println(v)
	}
}

func PartOne(input [][]string) {
	totalSum := 0

	for i := 0; i < len(input[0]); i++ {
		curSum := 0
		curOperator := ""
		for j := len(input) - 1; j >= 0; j-- {
			if j == len(input) - 1 {
				curOperator = input[j][i]
				if curOperator == "*" {
					curSum = 1
				}
			} else {
				val, err := strconv.Atoi(input[j][i])
				if err != nil {
					log.Fatal(err)
				}
				switch curOperator {
				case "*":
					curSum *= val
				case "+":
					curSum += val
				}
			}
		}
		totalSum += curSum
	}
	fmt.Println("Total:", totalSum)
}

func GetInput() [][]string {
	file, err := os.Open("input_test.txt")
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
				inputs = append(inputs, strings.Fields(strings.TrimRight(line, "\n")))
				break
			} else {
				log.Fatal(err)
				break
			}
		}
		inputs = append(inputs, strings.Fields(strings.TrimRight(line, "\n")))
	}
	return inputs
}