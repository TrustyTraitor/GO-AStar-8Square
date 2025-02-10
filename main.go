package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ReadConfig(r io.Reader) (Puzzles, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var result Puzzles
	var puzzle Puzzle

	row := 0
	for scanner.Scan() {

		for idx, item := range strings.Split(scanner.Text(), " ") {
			x, err := strconv.Atoi(item)
			if err != nil {
				return result, err
			}

			if idx < 3 {
				puzzle.current[row%3][idx%3] = x
			} else {
				puzzle.goal[row%3][idx%3] = x
			}
		}

		if (row+1)%3 == 0 {
			result = append(result, puzzle)
		}
		row++
	}
	return result, scanner.Err()
}

func main() {
	file, _ := os.Open("puzzles.txt")
	puzzles, _ := ReadConfig(file)

	for _, puzzle := range puzzles {
		ans, err := AStar(puzzle, ManhattanDistance)
		if err == nil {
			PrintResults(&ans)
		} else {
			fmt.Println(err)
		}
	}
}
