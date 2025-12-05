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
	fmt.Println("Day 1 Advent of Code")
	dialPtr := 50
	password := 0
	instructions := GetInput()
	for _, v := range(instructions) {
		direction := string(v[0])
		distance, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatal(err)
		}

		simplifiedDistance := distance % 100
		// increment for each full rotation past 0 (i.e. any number over 100)
		password += distance / 100

		if direction == "L" {
			if dialPtr != 0 && dialPtr - simplifiedDistance < 0 {
				password++
			}
			simplifiedDistance = 100 - simplifiedDistance
		} else if dialPtr != 0 && dialPtr + simplifiedDistance > 100 {
			password++
		}

		RotateDial(&dialPtr, simplifiedDistance)
		if dialPtr == 0 {
			// increment password if dial ends on 0
			password++
		}
	}
	fmt.Println("Password:", password)
}

func RotateDial(ptr *int, dist int) {
	newPtr := *ptr + dist
	if newPtr >= 100 {
		*ptr = newPtr - 100
	} else {
		*ptr = newPtr
	}
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