// Package benchmark contains utility functions to calculate the execution time of functions.
package benchmark

import (
	"fmt"
	"time"
)

// Elapsed returns the time that a function took to execute in microseconds.
func Elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
