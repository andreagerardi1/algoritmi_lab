package main

import ("fmt")

func main() {
	fmt.Println(multiply(0,0))
}

func multiply (x , y int) int {
	if y == 0 {
		return 0
	} else {
		return x + multiply (x , y -1)
	}
}
