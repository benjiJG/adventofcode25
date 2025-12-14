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
	input := GetInput()
	splits := 0
	for _, v := range input {
		for i := 0; i < len(v); i++ {
			if v[i] == '^' {
				if i - 1 > 0 && i - 1 >= 0 {
					splits++
				}
			}
		}
	}
	fmt.Println("Splits:", splits)
}

func GetInput() []string {
	file, err := os.Open("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputs []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				inputs = append(inputs, strings.TrimRight(line, "\n"))
				break
			} else {
				log.Fatal(err)
				break
			}
		}
		inputs = append(inputs, strings.TrimRight(line, "\n"))
	}
	return inputs
}