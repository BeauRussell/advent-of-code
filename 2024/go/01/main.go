package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getArrays() ([]string, []string) {
	file, err := os.Open("../../puzzles/01/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left []string
	var right []string
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "   ")
		left = append(left, lineSplit[0])
		right = append(right, lineSplit[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return left, right
}

func sortSlice(slice []string) []string {
	sort.Strings(slice)
	return slice
}

func findDistance(right []string, left []string) int {
	var distance int

	for i := 0; i < len(right); i++ {
		leftVal, _ := strconv.Atoi(left[i])
		rightVal, _ := strconv.Atoi(right[i])
		distance += int(math.Abs(float64(leftVal) - float64(rightVal)))
	}

	return distance
}

func findSimilarity(left []string, right []string) int {
	var similarity int

	for i := 0; i < len(left); i++ {
		leftVal, _ := strconv.Atoi(left[i])
		var appearances int
		for j := 0; j <= len(right); j++ {
			rightVal, _ := strconv.Atoi(right[j])
			if leftVal < rightVal {
				break
			} else if leftVal == rightVal {
				appearances++
			}
		}

		similarity += appearances * leftVal
	}

	return similarity
}

func main() {
	left, right := getArrays()
	left = sortSlice(left)
	right = sortSlice(right)

	fmt.Println(findDistance(left, right))
	fmt.Println(findSimilarity(left, right))
}
