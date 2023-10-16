package main

import "fmt"

func main() {
	fmt.Println("Serie di Fibonacci:")
    	for i := 1; i <= 100; i++ {
        	fmt.Println(i, f_rec(i))
    	}
}

func f_rec(n int) uint64 {
	if n == 1 || n == 2 {
		return 1
	}
	return f_rec(n-1) + f_rec(n-2)
}
