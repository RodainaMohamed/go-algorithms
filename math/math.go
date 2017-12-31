package math

// Fibonacci returns the fibonacci of the number passed in the parameter using a naive recursive approach.
func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func fibonacciMemoHelper(n int, memo []int) int {
	if n == 0 || n == 1 {
		return n
	}
	if memo[n-1] != -1 {
		return memo[n-1]
	}
	memo[n-1] = fibonacciMemoHelper(n-1, memo) + fibonacciMemoHelper(n-2, memo)
	return memo[n-1]
}

// FibonacciMemo returns the fibonacci of the number passed in the parameter using the dynamic programming memoization recursive approach.
func FibonacciMemo(n int) int {
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	return fibonacciMemoHelper(n, memo)
}

// FibonacciDP returns the fibonacci of the number passed in the parameter using the dynammic programming tabulation iterative approach.
func FibonacciDP(n int) int {
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}
