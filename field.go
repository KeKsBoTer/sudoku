package sudoku

import (
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// EmptyCell means a cell has no value
const EmptyCell = 0

// Field represents a sudoku field and holds the number for every cell
// 0 means a cell is empty
type Field [9][9]int

// EmptyCells returns the count of empty cells
func (f *Field) EmptyCells() int {
	var empty int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if f[i][j] == EmptyCell {
				empty++
			}
		}
	}
	return empty
}

// Check validates if all cells are filled according to the sudoku rules
func (f *Field) Check() error {
	var err error
	for i := 0; i < 9; i++ {
		if err = f.checkColumn(i); err != nil {
			return err
		}
		if err = f.checkRow(i); err != nil {
			return err
		}
		if err = f.checkSquare((i-(i/3)*3)*3, (i/3)*3); err != nil {
			return err
		}
	}
	return nil
}

// checks if row in field is valid
func (f *Field) checkRow(row int) error {
	var pos Possibilities
	var num int
	for i := 0; i < 9; i++ {
		num = f[row][i]
		if num > 0 {
			if pos.IsPossible(num) {
				return ErrorSudoku{
					num:   num,
					x:     i,
					y:     row,
					eType: Row,
				}
			}
			pos.Add(num)
		}
	}
	return nil
}

func (f *Field) checkColumn(col int) error {
	var pos Possibilities
	var num int
	for i := 0; i < 9; i++ {
		num = f[i][col]
		if num > 0 {
			if pos.IsPossible(num) {
				return ErrorSudoku{
					num:   num,
					x:     col,
					y:     i,
					eType: Column,
				}
			}
			pos.Add(num)
		}
	}
	return nil
}

func (f *Field) checkSquare(x, y int) error {
	var pos Possibilities
	var num int
	for i := y; i < y+3; i++ {
		for j := x; j < x+3; j++ {
			num = f[i][j]
			if num > 0 {
				if pos.IsPossible(num) {
					return ErrorSudoku{
						num:   num,
						x:     j,
						y:     i,
						eType: Square,
					}
				}
				pos.Add(num)
			}
		}
	}
	return nil
}

// String turns object into human readable string
// See Field.PrettyPrint
func (f Field) String() string {
	return f.PrettyPrint(nil)
}

// PrettyPrint turns object into human readable string
// changes to the initial field are printed in green
// e.g.:
// ╔═════╦═════╦═════╗
// ║7│ │ ║ │ │4║8│1│ ║
// ║1│8│ ║ │7│5║4│ │ ║
// ║ │4│9║ │ │1║7│ │ ║
// ╠═════╬═════╬═════╣
// ║4│ │1║5│3│ ║ │9│7║
// ║8│ │7║4│ │ ║ │ │ ║
// ║ │ │ ║6│1│7║ │8│4║
// ╠═════╬═════╬═════╣
// ║ │3│8║7│5│ ║ │4│1║
// ║ │1│4║2│ │ ║3│7│5║
// ║ │7│ ║1│4│3║ │ │8║
// ╚═════╩═════╩═════╝
func (f Field) PrettyPrint(initial *Field) string {
	b := strings.Builder{}
	var num string
	b.WriteString("╔═════╦═════╦═════╗\n")
	for i := 0; i < 9; i++ {
		b.WriteString("║")
		for j := 0; j < 9; j++ {
			if f[i][j] == 0 {
				b.WriteString(" ")
			} else {
				num = strconv.Itoa(f[i][j])
				if initial != nil && f[i][j] != initial[i][j] {
					num = color.HiGreenString(num)
				}
				b.WriteString(num)
			}
			if j < 8 {
				if (j+1)%3 == 0 {
					b.WriteString("║")
				} else {
					b.WriteString("│")
				}
			}
		}
		b.WriteString("║")
		b.WriteString("\n")
		if (i+1)%3 == 0 && i != 8 {
			b.WriteString("╠═════╬═════╬═════╣\n")
		}
	}
	b.WriteString("╚═════╩═════╩═════╝")
	return b.String()
}
