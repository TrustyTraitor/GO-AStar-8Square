package main

import (
	"math"
	"sort"
)

// Heuristic - A type alias for easier reading and writing
type Heuristic = func(*Puzzle) int

// LocateNumber - Finds a number in a multidimensional array and returns it's indices
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

// ManhattanDistance - A heuristic function for the A* cost evaluation
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

// MisplacedTiles - A heuristic function for the A* cost evaluation
func MisplacedTiles(puzzle *Puzzle) int {
	cost := 0

	for idx, i := range puzzle.current {
		for jdx, j := range i {
			if j != puzzle.goal[idx][jdx] {
				cost++
			}
		}
	}

	return cost
}

// GenerateNodes - Creates and returns the nodes that can be generated from a given state.
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

// RemoveElement - Removes an element from a slice and returns it.
func RemoveElement(slice Puzzles, idx int) (Puzzle, Puzzles) {
	popped := slice[idx]
	slice = append(slice[:idx], slice[idx+1:]...)

	return popped, slice
}

// NoSolution - A custom error type returned when no solution is found.
type NoSolution struct{}

// Error - Implements the 'error' interface on the NoSolution type
func (ns *NoSolution) Error() string {
	return "No Solution"
}

// AStar - Implements A* algorithm. Takes in a heuristic function and a depth limit. Returns a SolvedPuzzle type and error.
// Error will be nil if no error occurs. Otherwise, returns a NoSolution
func AStar(puzzle Puzzle, h Heuristic, depthLimit int) (SolvedPuzzle, error) {
	puzzle.h = h(&puzzle)

	openList := Puzzles{puzzle}
	closedList := Puzzles{}

	var currentPuzzle Puzzle

	for len(openList) > 0 {
		if len(closedList)+len(openList) > depthLimit {
			break
		}
		// Sorts the openList by the f() value so that the item at the start has the lowest f()
		sort.Sort(openList)

		// Pops the first element off the openList
		currentPuzzle, openList = RemoveElement(openList, 0)
		if currentPuzzle.h == 0 { // Check if goal state has been reached
			return SolvedPuzzle{
				Puzzle:         currentPuzzle,
				expandedNodes:  len(closedList),
				generatedNodes: len(openList) + len(closedList),
			}, nil
		}

		// Adds the current node aka, the node currently being expanded, to the closed list
		closedList = append(closedList, currentPuzzle)
		// Expand the current best puzzle, and if they are not in the closedList, add them to the openList
		for _, neighbor := range GenerateNodes(currentPuzzle, h) {
			res := neighbor.In(&closedList)
			if !res {
				openList = append(openList, neighbor)
			}
		}
	}

	return SolvedPuzzle{
		Puzzle: Puzzle{},
	}, &NoSolution{}
}
