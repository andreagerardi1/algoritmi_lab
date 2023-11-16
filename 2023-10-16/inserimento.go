package main

import (
	"fmt"
)

func main() {
	fmt.Println(Inserimento())
}

func Inserimento() []int {
	var sl []int
	var n int
	for {
		fmt.Scan(&n)
		if n != 0 {
			if len(sl) == 0 {
				sl = append(sl, n)
			} else {
				sl = append(sl, n)
				for i := len(sl) - 2; i >= 0; i-- {
					if sl[i] > n {
						sl[i+1] = sl[i]
						sl[i] = n
					}
				}
			}
		} else {
			return sl
		}
	}
}
