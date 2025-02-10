package main

import "fmt"

type Puzzle struct {
	current [3][3]int
	goal    [3][3]int
	h       int
	g       int
	parent  *Puzzle
}

func (p *Puzzle) PrettyPrint() {
	for _, i := range p.current {
		fmt.Println(i)
	}
	fmt.Printf("\tf() = %v\n", p.h+p.g)
}

func (p *Puzzle) TracePath() {
	if p.parent != nil {
		p.parent.TracePath()
	}
	p.PrettyPrint()
}

func (p *Puzzle) f() int {
	return p.h + p.g
}

func (p *Puzzle) Equals(other *Puzzle) bool {
	return p.current == other.current
}

func (p *Puzzle) In(list *Puzzles) bool {
	for _, i := range *list {
		if i.Equals(p) {
			return true
		}
	}
	return false
}

type Puzzles []Puzzle

func (p Puzzles) Len() int {
	return len(p)
}

func (p Puzzles) Less(i, j int) bool {
	return p[i].f() < p[j].f()
}

func (p Puzzles) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type SolvedPuzzle struct {
	Puzzle
	generatedNodes int
	expandedNodes  int
}
