package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func getArrays() [][]string {
	file, err := os.Open("../../puzzles/04/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "")
		result = append(result, lineSplit)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func findXmas(puzzle [][]string) {
	var occurrences int
	for i, row := range puzzle {
		for j, cell := range row {
			if cell == "X" {
				occurrences += checkRow(i, j, puzzle)
				occurrences += checkCol(i, j, puzzle)
				occurrences += checkDiags(i, j, puzzle)
			}
		}
	}

	fmt.Println(occurrences)
}

func checkRow(i int, j int, puzzle [][]string) int {
	var occurrences int
	if len(puzzle[i])-j >= 4 && reflect.DeepEqual(puzzle[i][j:j+4], []string{"X", "M", "A", "S"}) {
		occurrences++
	}
	if j >= 3 && reflect.DeepEqual(puzzle[i][j-3:j+1], []string{"S", "A", "M", "X"}) {
		occurrences++
	}

	return occurrences
}

func checkCol(row int, col int, puzzle [][]string) int {
	var occurrences int
	columnSlice := make([]string, len(puzzle))

	for i := 0; i < len(puzzle); i++ {
		columnSlice[i] = puzzle[i][col]
	}

	if len(columnSlice)-row >= 4 && reflect.DeepEqual(columnSlice[row:row+4], []string{"X", "M", "A", "S"}) {
		occurrences++
	}
	if row >= 3 && reflect.DeepEqual(columnSlice[row-3:row+1], []string{"S", "A", "M", "X"}) {
		occurrences++
	}

	return occurrences
}

func checkDiags(row int, col int, puzzle [][]string) int {
	var occurrences int
	var downRight = make([]string, 4)
	var downLeft = make([]string, 4)
	var upRight = make([]string, 4)
	var upLeft = make([]string, 4)
	if len(puzzle)-row >= 4 && len(puzzle[row])-col >= 4 {
		for j := 0; j < 4; j++ {
			downRight[j] = puzzle[row+j][col+j]
		}
		if reflect.DeepEqual(downRight, []string{"X", "M", "A", "S"}) {
			occurrences++
		}
	}

	if len(puzzle)-row >= 4 && col >= 3 {
		for j := 0; j < 4; j++ {
			downLeft[j] = puzzle[row+j][col-j]
		}
		if reflect.DeepEqual(downLeft, []string{"X", "M", "A", "S"}) {
			occurrences++
		}
	}

	if row >= 3 && len(puzzle[row])-col >= 4 {
		for j := 0; j < 4; j++ {
			upRight[j] = puzzle[row-j][col+j]
		}
		if reflect.DeepEqual(upRight, []string{"X", "M", "A", "S"}) {
			occurrences++
		}
	}

	if row >= 3 && col >= 3 {
		for j := 0; j < 4; j++ {
			upLeft[j] = puzzle[row-j][col-j]
		}
		if reflect.DeepEqual(upLeft, []string{"X", "M", "A", "S"}) {
			occurrences++
		}
	}

	return occurrences
}

func findMasX(puzzle [][]string) {
	var occurrences int

	for i, row := range puzzle {
		for j, cell := range row {
			if cell == "A" {
				if i >= 1 && j >= 1 && len(puzzle)-i-1 >= 1 && len(puzzle[i])-j-1 >= 1 {
					if checkDiag([]string{puzzle[i-1][j-1], cell, puzzle[i+1][j+1]}) && checkDiag([]string{puzzle[i-1][j+1], cell, puzzle[i+1][j-1]}) {
						occurrences++
					}
				}
			}
		}
	}

	fmt.Println(occurrences)
}

func checkDiag(diag []string) bool {
	var safe bool
	if reflect.DeepEqual(diag, []string{"M", "A", "S"}) || reflect.DeepEqual(diag, []string{"S", "A", "M"}) {
		safe = true
	} else {
		safe = false
	}
	return safe
}

func main() {
	puzzle := getArrays()
	findXmas(puzzle)
	findMasX(puzzle)
}
