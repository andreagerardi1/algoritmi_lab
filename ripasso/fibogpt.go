package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Serie di Fibonacci:")
    	for i := 1; i <= 30; i++ {
        	fmt.Println(fibonacci(i))
    	}
}

