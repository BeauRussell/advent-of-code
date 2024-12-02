package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getArrays() [][]string {
	file, err := os.Open("../../puzzles/02/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")
		result = append(result, lineSplit)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func findSafeReports(puzzle [][]string) {
	var answer int
	for _, row := range puzzle {
		var inc = true
		var safe = true
		for i, _ := range row {
			if i == 0 {
				continue
			}
			current, _ := strconv.Atoi(row[i])
			last, _ := strconv.Atoi(row[i-1])
			if i == 1 && current < last {
				inc = false
			}
			if (inc && current <= last || current-last > 3) || (!inc && current >= last || last-current > 3) {
				safe = false
				break
			}
		}
		if safe {
			answer++
		}
	}

	fmt.Println("Part 1:", answer)
}

func part2(input [][]string) {
	output := 0

	for _, line := range input {
		levels := make([]int, len(line))

		for i, part := range line {
			num, _ := strconv.Atoi(part)
			levels[i] = num
		}

		isSafe := false
		isSafe = levelIsSafe(levels)

		if !isSafe {
			for i := 0; i < len(levels); i++ {
				modifiedLevels := []int{}
				modifiedLevels = append(modifiedLevels, levels[:i]...)
				modifiedLevels = append(modifiedLevels, levels[i+1:]...)
				isSafe = levelIsSafe(modifiedLevels)

				if isSafe {
					break
				}
			}
		}

		if isSafe {
			output += 1
		}
	}

	fmt.Println("Part 2", output)
}

func levelIsSafe(levels []int) bool {
	isIncreasing := true
	isSafe := true
	badLevel := 0

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff == 0 {
			badLevel++
			isSafe = false
			break
		}

		if diff < -3 || diff > 3 {
			isSafe = false
			break
		}

		if i == 1 && diff < 0 {
			isIncreasing = false
		} else if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
			isSafe = false
			break
		}
	}

	return isSafe
}

func removeIndex(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	var puzzle = getArrays()
	findSafeReports(puzzle)
	part2(puzzle)
}
