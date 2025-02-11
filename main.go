package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// ReadConfig - Reads the file containing the problems and returns them in a slice of Puzzle(s)
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
	printPaths := true

	fmt.Println("Print the Paths? (0 - no ; 1 - yes)")
	fmt.Println("\t Due to the data structure used to store the path, it is difficult to print it horizontally.")
	fmt.Println("\t Printing paths takes a lot of room: For a cleaner output, type 0")
	_, err := fmt.Scanf("%v", &printPaths)

	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	file, _ := os.Open("puzzles.txt")
	puzzles, _ := ReadConfig(file)

	depthLimit := 20000

	hs := []Heuristic{ManhattanDistance, MisplacedTiles}
	hsName := []string{"Manhattan Distance", "Misplaced Tiles"}

	for idx, puzzle := range puzzles {
		fmt.Println("<[==========#***#==========]>")
		for jdx, h := range hs {
			ans, err := AStar(puzzle, h, depthLimit)
			fmt.Printf("Puzzle #%v\n", idx+1)
			fmt.Printf("\tHeuristic %v\n", hsName[jdx])
			if err == nil {
				PrintResults(&ans, printPaths)
				fmt.Println()
			} else {
				fmt.Printf("No Solution Found\n")
				fmt.Printf("\tDepth Limit: %v\n\n", depthLimit)
			}

		}
	}
}
