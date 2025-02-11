package main

import "fmt"

// Puzzle - Type that holds a given state, goal, and stats about an 8 problem board.
type Puzzle struct {
	current [3][3]int
	goal    [3][3]int
	h       int
	g       int
	parent  *Puzzle
}

// PrettyPrint - Prints the 8 Puzzle in an easier to read way
func (p *Puzzle) PrettyPrint() {
	for _, i := range p.current {
		fmt.Println(i)
	}
	fmt.Printf("\tf() = %v\n", p.h+p.g)
}

// TracePath - Runs up a Puzzle going up to the first ancestor, then printing down to the solution in order
func (p *Puzzle) TracePath() {
	if p.parent != nil {
		p.parent.TracePath()
	}
	p.PrettyPrint()
}

// F - Calculates the F value of a given state
func (p *Puzzle) f() int {
	return p.h + p.g
}

// Equals - Compares two puzzle states by their current state
func (p *Puzzle) Equals(other *Puzzle) bool {
	return p.current == other.current
}

// In - Checks if a given puzzle state is contained in a list of puzzle states
func (p *Puzzle) In(list *Puzzles) bool {
	for _, i := range *list {
		if i.Equals(p) {
			return true
		}
	}
	return false
}

// Puzzles - Type alias for a list of Puzzles
type Puzzles []Puzzle

// Len - needed for the sort.Sort function. Returns the len of Puzzles
func (p Puzzles) Len() int {
	return len(p)
}

// Less - needed for sort.Sort function. Compares two puzzles to see which one comes first based on F value
func (p Puzzles) Less(i, j int) bool {
	return p[i].f() < p[j].f()
}

// Swap - needed for sort.Sort function. Does as says
func (p Puzzles) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// SolvedPuzzle - Contains a solved puzzle which can be traversed using TracePath to print the puzzle steps.
// Also contains the number of nodes generated and expanded while solving the puzzle.
type SolvedPuzzle struct {
	Puzzle
	generatedNodes int
	expandedNodes  int
}

// PrintResults - Prints the results of a SolvedPuzzle. Prints in order as well as number of nodes generated and expanded.
func PrintResults(sp *SolvedPuzzle, full bool) {
	if full {
		sp.TracePath()
	}
	fmt.Printf("Nodes Generated: %v\n", sp.generatedNodes)
	fmt.Printf(" Nodes Expanded: %v\n", sp.expandedNodes)

}
