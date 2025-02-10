package main

import (
	"fmt"
	"math"
	"sort"
)

type Heuristic = func(*Puzzle) int

func LocateNumber(arr *[3][3]int, val int) (int, int) {
	for idx, i := range arr {
		for jdx, j := range i {
			if j == val {
				return idx, jdx
			}
		}
	}
	return -1, -1
}

func ManhattanDistance(puzzle *Puzzle) int {
	cost := 0

	for idx, i := range puzzle.current {
		for jdx, j := range i {
			if j != 0 {
				locI, locJ := LocateNumber(&puzzle.goal, j)

				cost += int(math.Abs(float64(locI-idx)) + math.Abs(float64(locJ-jdx)))
			}
		}
	}

	return cost
}

func GenerateNodes(puzzle Puzzle, h Heuristic) Puzzles {
	// Gets the indices for the valid swaps to generate children nodes
	GenValidSwaps := func(r, c int) [][2]int {
		var swaps [][2]int

		left := c - 1
		if left >= 0 {
			swaps = append(swaps, [2]int{r, left})
		}

		right := c + 1
		if right < len(puzzle.current) {
			swaps = append(swaps, [2]int{r, right})
		}

		top := r - 1
		if top >= 0 {
			swaps = append(swaps, [2]int{top, c})
		}

		bottom := r + 1
		if bottom < len(puzzle.current) {
			swaps = append(swaps, [2]int{bottom, c})
		}

		return swaps
	}

	var nodes Puzzles

	row, col := LocateNumber(&puzzle.current, 0)

	for _, points := range GenValidSwaps(row, col) {
		newNode := Puzzle{
			current: puzzle.current,
			goal:    puzzle.goal,
			parent:  &puzzle,
		}
		node := &newNode.current

		val := node[points[0]][points[1]]

		node[points[0]][points[1]] = 0
		node[row][col] = val

		newNode.h = h(&newNode)
		newNode.g = puzzle.g + 1
		//newNode.parent = puzzle

		nodes = append(nodes, newNode)
	}

	return nodes
}

func RemoveElement(slice Puzzles, idx int) (Puzzle, Puzzles) {
	popped := slice[idx]
	slice = append(slice[:idx], slice[idx+1:]...)

	return popped, slice
}

type NoSolution struct{}

func (ns NoSolution) Error() string {
	return "No Solution"
}

func AStar(puzzle Puzzle, h Heuristic) (SolvedPuzzle, error) {
	genCount := 0
	expCount := 0

	puzzle.h = h(&puzzle)

	openList := Puzzles{puzzle}
	closedList := Puzzles{}

	var currentPuzzle Puzzle

	for len(openList) > 0 {
		//fmt.Println(len(openList))
		sort.Sort(openList)

		currentPuzzle, openList = RemoveElement(openList, 0)
		//fmt.Println(currentPuzzle.f())
		if currentPuzzle.h == 0 { // Check if goal state has been reached
			return SolvedPuzzle{
				Puzzle:         currentPuzzle,
				expandedNodes:  expCount,
				generatedNodes: genCount,
			}, nil
		}

		// Expand the current best puzzle
		closedList = append(closedList, currentPuzzle)
		for _, neighbor := range GenerateNodes(currentPuzzle, h) {
			res := neighbor.In(&closedList)
			//fmt.Println(res)
			if !res {
				openList = append(openList, neighbor)
			}
		}
	}

	return SolvedPuzzle{
		Puzzle:         currentPuzzle,
		expandedNodes:  expCount,
		generatedNodes: genCount,
	}, NoSolution{}
}

func PrintResults(sp *SolvedPuzzle) {
	sp.TracePath()
	fmt.Printf("Nodes Generated: %v\n", sp.generatedNodes)
	fmt.Printf(" Nodes Expanded: %v\n", sp.expandedNodes)

}
