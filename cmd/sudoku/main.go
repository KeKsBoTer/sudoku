package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KeKsBoTer/sudoku"
)

func main() {
	verbose := flag.Bool("verbose", false, "Prints single steps to console")
	debug := flag.Bool("debug", false, "Same as verbose, but stops after every step")
	delay := flag.Int64("delay", 100, "Delay in milliseconds between steps in verbose mode")
	file := flag.String("file", "sudoku.csv", "Path to the sudoku CSV file")
	flag.Parse()

	if *debug {
		*verbose = true
	}

	field, err := readCSV(*file)
	if err != nil {
		fmt.Printf("Error reading '%s': %s", *file, err)
		return
	}
	solved, err := sudoku.Solve(*field, func(updated sudoku.Field) {
		if *verbose {
			fmt.Println(updated)
			if *debug {
				fmt.Scanln()
			} else if *delay > 0 {
				time.Sleep(time.Duration(*delay) * time.Millisecond)
			}

		}
	})
	fmt.Println(solved)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Solved!")
	}
}

// read sodoku field from csv file
func readCSV(path string) (*sudoku.Field, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var field sudoku.Field
	for line := 0; ; line++ {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if len(row) != 9 {
			return nil, fmt.Errorf("Line %d has %d cells", line+1, len(row))
		}
		for i, v := range row {
			v = strings.Trim(v, " ")
			if len(v) == 0 {
				field[line][i] = 0
				continue
			}

			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("Column %d in line %d is not a number: '%s'", i, line+1, v)
			}
			field[line][i] = num
		}
	}
	return &field, nil
}
