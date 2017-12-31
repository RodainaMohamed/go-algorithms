# go-algorithms

### Array Utilities

* Searching
  * Linear Search
  * Binary Search

* Sorting
  * Bubble Sort
  * Selection Sort
  * Double Selection Sort
  * Insertion Sort

### String Manipulation

* String Reverse
* Palindrome Check

---

### Usage Example

~~~~
package main

import (
	"fmt"
	"strconv"

	"github.com/omardoma/go-algorithms/arrayutil/search"
	"github.com/omardoma/go-algorithms/arrayutil/sort"
	"github.com/omardoma/go-algorithms/benchmark"
	"github.com/omardoma/go-algorithms/math"
	"github.com/omardoma/go-algorithms/stringutil"
)

func main() {
	var arr []int
	var estimateExecTime func()

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("Sorted Array:", arr)
	fmt.Println("\nBinary Search")
	estimateExecTime = benchmark.Elapsed("Binary Search")
	fmt.Println("Key Index: ", search.BinarySearch(10, arr))
	estimateExecTime()

	arr = []int{3, 1, 4, 2, 6, 5, 8, 7, 10, 9}
	fmt.Println("\nDouble Selection Sort")
	estimateExecTime = benchmark.Elapsed("Double Selection Sort")
	sort.DoubleSelectionSort(arr)
	fmt.Println(arr)
	estimateExecTime()

	str := "Omar loves Tuffy so much"
	fmt.Println("\nReverse String")
	estimateExecTime = benchmark.Elapsed("Reverse String")
	fmt.Println(stringutil.Reverse(str))
	estimateExecTime()

	n := 10
	fmt.Println("\nMemoized Fibonacci")
	estimateExecTime = benchmark.Elapsed("Memoized Fibonacci")
	fmt.Println("Fibonacci of " + strconv.Itoa(n) + ": " + strconv.Itoa(math.FibonacciMemo(n)))
	estimateExecTime()
}

~~~~

© 2017 Omar Doma
