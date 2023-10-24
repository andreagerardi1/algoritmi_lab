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
	var flag bool
	for {
		fmt.Scan(&n)
		if n != 0 {
			if len(sl) == 0 {
				sl = append(sl, n)
			} else {
				flag = true
				sl = append(sl, 0)
				for i := len(sl) - 2; i >= 0; i-- {
					if sl[i] > n {
						sl[i+1] = sl[i]
						sl[i] = n
						flag = false
					}
					if flag {
						sl[len(sl)-1] = n
					}
				}
			}
		} else {
			return sl
		}
	}
}
