package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func getString() string {
	file, err := os.Open("../../puzzles/03/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var buffer bytes.Buffer

	for scanner.Scan() {
		line := scanner.Text()
		buffer.WriteString(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return buffer.String()
}

func runFunctions(text string) {
	var answer int

	r, _ := regexp.Compile("(mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\))")

	matches := r.FindAllStringSubmatch(text, -1)
	var do = true

	for _, match := range matches {
		if match[1] == "do()" {
			do = true
			continue
		} else if match[1] == "don't()" {
			do = false
			continue
		} else if do {
			left, _ := strconv.Atoi(match[2])
			right, _ := strconv.Atoi(match[3])

			answer += left * right
		}
	}

	fmt.Println(answer)
}

func main() {
	var text = getString()
	runFunctions(text)
}
