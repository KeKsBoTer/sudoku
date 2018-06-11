package sudoku

import (
	"fmt"
)

// ErrorType is the rule which is violated
type ErrorType string

const (
	// Column means a number is represented in a column more than once
	Column = ErrorType("column")

	// Row means a number is represented in a row more than once
	Row = ErrorType("row")

	// Square means a number is represented in a 3x3 square more than once
	Square = ErrorType("square")
)

// ErrorSudoku is a violation of the sukodu rules
type ErrorSudoku struct {
	x, y  int
	num   int
	eType ErrorType
}

func (err ErrorSudoku) Error() string {
	return fmt.Sprintf("Number %d is present more than once in %s at (%d,%d)", err.num, err.eType, err.x, err.y)
}
