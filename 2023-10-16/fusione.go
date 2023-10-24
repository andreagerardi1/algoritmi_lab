package main

import "fmt"

func main() {
	sl := []int{3, 5, 1, 8, 5, 7, 10}
	fmt.Println(mergeSort(sl))
}

func mergeSort(sl []int) []int {
	if len(sl) <= 1 {
		return sl
	}
	b := mergeSort(sl[:len(sl)/2])
	c := mergeSort(sl[len(sl)/2:])
	res := merge(b, c)
	return res
}

func merge(b, c []int) []int {
	lenghts := len(b) + len(c)
	a := make([]int, lenghts)
	var idxb, idxc int
	for i := 0; i < len(a); i++ {
		if idxb == len(b) {
			a[i] = c[idxc]
			idxc++
		} else if idxc == len(c) {
			a[i] = b[idxb]
			idxb++
		} else if b[idxb] > c[idxc] {
			a[i] = c[idxc]
			idxc++
		} else {
			a[i] = b[idxb]
			idxb++
		}
	}
	return a
}
