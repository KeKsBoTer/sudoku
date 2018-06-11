package sudoku

import (
	"strconv"
)

// Possibilities holds all possible possible numbers in a cell
type Possibilities [9]bool

// NewPossibilities is a new possibility object with every number possible
func NewPossibilities() *Possibilities {
	var solver Possibilities
	for i := range solver {
		solver[i] = true
	}
	return &solver
}

// IsPossible checks if a certian number is possibile
func (p *Possibilities) IsPossible(n int) bool {
	return p[n-1]
}

// OnlyOne checks if only one number is possible and returns that number
func (p *Possibilities) OnlyOne() (bool, int) {
	found, count := 0, 0
	for k := 1; k <= 9; k++ {
		if p.IsPossible(k) {
			found = k
			count++
		}
	}
	return count == 1, found
}

// Remove removes number from possibilities
func (p *Possibilities) Remove(n int) {
	p[n-1] = false
}

// Add adds number to possibilities
func (p *Possibilities) Add(n int) {
	p[n-1] = true
}

// Empty checks if any number is possible
func (p *Possibilities) Empty() bool {
	for _, v := range p {
		if v {
			return false
		}
	}
	return true
}

// String turns Possibilities into human readable string
// It prints all possible numbers as list
// e.g. [4,5,8]
func (p Possibilities) String() string {
	out := "["
	for n, v := range p {
		if v {
			if out[len(out)-1] != '[' {
				out += ","
			}
			out += strconv.Itoa(n + 1)
		}
	}
	out += "]"
	return out
}

// merges multiple Possibilities into one,
// where a number is only possible if it is possible in all given Possibilities
func mergePossibilities(poses ...*Possibilities) *Possibilities {
	checker := NewPossibilities()
	for i := range poses[0] {
		for _, c := range poses {
			if !c.IsPossible(i+1) && checker.IsPossible(i+1) {
				checker.Remove(i + 1)
			}
		}
	}
	return checker
}
