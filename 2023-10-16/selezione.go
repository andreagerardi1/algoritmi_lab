package main

import "fmt"

func main() {
	sl := []int{9, 2, 5, 7, 3, 2, 10}
	fmt.Println("iterativa", SelectionSortIter(sl))
	fmt.Println("ricorsiva", SelectionSortRic(sl))
}

func SelectionSortIter(sl []int) []int {
	var max, maxInd, swap int
	for i := len(sl) - 1; i > 0; i-- {
		max = sl[0]
		for j := 0; j <= i; j++ {
			if sl[j] >= max {
				max = sl[j]
				maxInd = j
			}
		}
		swap = sl[i]
		sl[i] = max
		sl[maxInd] = swap
	}
	return sl
}

func SelectionSortRic(sl []int) []int {
	if len(sl) <= 1 {
		return sl
	} else {
		max := sl[0]
		var maxInd, swap, appoggio int
		for i := 0; i < len(sl); i++ {
			if sl[i] >= max {
				max = sl[i]
				maxInd = i
			}
		}
		swap = sl[len(sl)-1]
		sl[len(sl)-1] = max
		sl[maxInd] = swap
		appoggio = sl[len(sl)-1]
		sl = SelectionSortRic(sl[:len(sl)-1])
		sl = append(sl, appoggio)
		return sl
	}
}
