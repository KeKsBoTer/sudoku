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
	verbose := flag.Bool("v", false, "Verbose: prints single steps to console")
	debug := flag.Bool("d", false, "Debug: Same as verbose, but stops after every step")
	delay := flag.Int64("delay", 100, "Delay in milliseconds between steps in verbose mode")
	file := flag.String("file", "sudoku.csv", "Path to the sudoku CSV file")
	flag.Parse()

	if *debug {
		*verbose = true
	}

	// read csv file
	csvFile, err := os.Open(*file)
	if err != nil {
		fmt.Printf("Error reading '%s': %s", *file, err)
		return
	}
	field, err := readCSV(bufio.NewReader(csvFile))
	if err != nil {
		fmt.Printf("Error parsing file '%s': %s", *file, err)
		return
	}

	// remove all new entered lines
	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			reader.ReadLine()
			Clear(1)
		}
	}()

	printed := false
	printField := func(new sudoku.Field) {
		if printed {
			Clear(13)
		}
		fmt.Println(new.PrettyPrint(field))
	}
	solved, err := sudoku.Solve(*field, func(updated sudoku.Field) {
		if *verbose {
			printField(updated)
			printed = true
			if *debug {
				fmt.Scanln()
			} else if *delay > 0 {
				time.Sleep(time.Duration(*delay) * time.Millisecond)
			}
		}
	})
	printField(*solved)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Solved!")
	}
}

// read sodoku field from csv file
func readCSV(r io.Reader) (*sudoku.Field, error) {
	reader := csv.NewReader(r)
	var field sudoku.Field
loop:
	for line := 0; ; line++ {
		row, err := reader.Read()
		// error handling
		switch {
		case err == io.EOF:
			if line != 9 {
				return nil, fmt.Errorf("Too little lines (%d)", line+1)
			}
			break loop
		case err != nil:
			return nil, err
		case line > 8:
			return nil, fmt.Errorf("Too many lines (%d)", line+1)
		case len(row) != 9:
			return nil, fmt.Errorf("Line %d has %d cells", line+1, len(row))
		}
		for i, v := range row {
			v = strings.Trim(v, " \r\n\t")
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
