package main

import (
	"fmt"
)

func main() {
	var a, b int
	var s string
	for i := 0; i < 10; i++{
		fmt.Scan(&s)
		if string(s[0]) == "a" {
			a++
		} else if string(s[0]) == "b" {
			b++
		}
	}
	fmt.Println("ci sono ", a, " parole che iniziano con la lettera 'a'")
	fmt.Println("ci sono ", b, " parole che iniziano con la lettera 'b'")
}
