# Sudoku Solver
A single CLI to solve sudokus.

# Installation

```
go get github.com/KeKsBoTer/sudoku
```

# Usage
Run the program in the console and add optional flags:
```
-debug
    Same as verbose, but stops after every step
-delay int
    Delay in milliseconds between steps in verbose mode (default 100)
-file string
    Path to the sudoku CSV file (default "sudoku.csv")
-verbose
    Prints single steps to console
```


# Roadmap

- [ ] Add unit tests
- [ ] Add colored output