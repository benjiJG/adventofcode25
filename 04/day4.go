package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 4 Advent of Code")
	input := GetInput()
	fmt.Println(input)
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