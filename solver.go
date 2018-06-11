package sudoku

import (
	"fmt"
)

// SolverField is a array to manage the possible numbers for all fields
type SolverField [9][9]*Possibilities

// PossibilitiesInRow calculates all posibile numbers in a row
func (f *Field) PossibilitiesInRow(row int) *Possibilities {
	checker := NewPossibilities()
	for i := 0; i < 9; i++ {
		value := f[row][i]
		if value > 0 {
			checker.Remove(value)
		}
	}
	return checker
}

// PossibilitiesInColumn calculates all posibile numbers in a column
func (f *Field) PossibilitiesInColumn(col int) *Possibilities {
	checker := NewPossibilities()
	for i := 0; i < 9; i++ {
		value := f[i][col]
		if value > 0 {
			checker.Remove(value)
		}
	}
	return checker
}

// PossibilitiesInSquare calculates all posibile numbers in a 3x3 square
// x and y are the offset for the squares left upper corner and must be a multiple of three
func (f *Field) PossibilitiesInSquare(x, y int) *Possibilities {
	if x%3 != 0 || y%3 != 0 {
		panic(fmt.Sprintf("%d and %d are not both a multiple of 3 or zero", x, y))
	}
	checker := NewPossibilities()
	for i := y; i < y+3; i++ {
		for j := x; j < x+3; j++ {
			value := f[i][j]
			if value > 0 {
				checker.Remove(value)
			}
		}
	}
	return checker
}

// UpdateFunc is called when field is updated
type UpdateFunc func(f Field)

// Solve solves sudoku field
func Solve(f Field, onUpdate UpdateFunc) (*Field, error) {

	var check SolverField
	var updated bool

	// function to update cell
	setCell := func(x, y, num int) {
		f[y][x] = num
		f.updatePossibilities(&check)
		onUpdate(f)
		updated = true
	}

	empty := f.EmptyCells()
	for empty > 0 {
		updated = false
		f.updatePossibilities(&check)

		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if ok, num := check[i][j].OnlyOne(); ok {
					setCell(j, i, num)
				}
			}
		}
		// check if number can only places at one call in row
		for i := 0; i < 9; i++ {
		num:
			for n := 1; n <= 9; n++ {
				pos := -1
				for j := 0; j < 9; j++ {
					if f[i][j] == 0 && check[i][j].IsPossible(n) {
						if pos != -1 {
							continue num
						}
						pos = j
					}
				}
				if pos != -1 {
					setCell(pos, i, n)
				}
			}
		}

		// check if number can only places at one call in column
		for i := 0; i < 9; i++ {
		num2:
			for n := 1; n <= 9; n++ {
				pos := -1
				for j := 0; j < 9; j++ {
					if f[j][i] == 0 && check[j][i].IsPossible(n) {
						if pos != -1 {
							continue num2
						}
						pos = j
					}
				}
				if pos != -1 {
					setCell(i, pos, n)
				}
			}
		}

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
			num3:
				for n := 1; n <= 9; n++ {
					posX, posY := -1, -1

					for y := 0; y < 3; y++ {
						for x := 0; x < 3; x++ {
							if f[i*3+y][j*3+x] == 0 && check[i*3+y][j*3+x].IsPossible(n) {
								if posX != -1 {
									continue num3
								}
								posX, posY = j*3+x, i*3+y
							}
						}
					}
					if posX != -1 {
						setCell(posX, posY, n)
					}
				}
			}
		}

		// solver is stuck if it cannot set new cells
		if !updated {
			return &f, fmt.Errorf("Stuck :(")
		}

		// check if new entered numbers are correct
		if err := f.Check(); err != nil {
			return &f, fmt.Errorf("Solution is false: %v", err)
		}
		empty = f.EmptyCells()
	}
	return &f, nil
}

// updates possible numbers for all cells
func (f *Field) updatePossibilities(c *SolverField) {
	var rows [9]*Possibilities
	var cols [9]*Possibilities
	var squares [3][3]*Possibilities
	for i := 0; i < 9; i++ {
		rows[i] = f.PossibilitiesInRow(i)
		cols[i] = f.PossibilitiesInColumn(i)
		squares[i/3][i-(i/3)*3] = f.PossibilitiesInSquare((i-(i/3)*3)*3, (i/3)*3)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if f[i][j] == 0 {
				c[i][j] = mergePossibilities(
					rows[i],
					cols[j],
					squares[i/3][j/3],
				)
			} else {
				// Field is allready set, no numbers are possible
				c[i][j] = new(Possibilities)
			}
		}
	}
}

// String turns the object into a human readable string for debugging
// It prints all possible numbers for every cell as list
// e.g.:
// (0,0): [1,2]
// (1,0): [3]
// ...
func (c SolverField) String() string {
	var out string
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			out += fmt.Sprintf("(%d,%d): %v\n", j, i, *(c[i][j]))
		}
	}
	return out
}
