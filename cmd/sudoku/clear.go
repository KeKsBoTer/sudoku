package main

import (
	"runtime"
)

// Clear removes the last n lines from the console output
func Clear(lines int) {
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		for i := 0; i < lines; i++ {
			print("\033[A\033[2K")
		}
	}
}
