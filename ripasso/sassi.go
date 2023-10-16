package main

import("fmt")

func main() {
	fmt.Println(sassi(4))
}

func sassi(height int) int {
	if height == 1 {
		return 1
	}
	return height * height + sassi(height -1)
}
