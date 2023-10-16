package main

import ( 
	"fmt"
)

func main() {
	fishes := []int{3,4,3,1,2}
	for i := 0; i < 80; i++ {
		fishes = NewDay(fishes)
		fmt.Println("after ", i+1, " days: ", fishes)
	}
	fmt.Println(len(fishes))
}

func AddFish(fishes *[]int) {
	*fishes = append(*fishes, 9)
}

func NewDay(fishes []int) []int {
	for i := 0; i < len(fishes); i++ {
		if fishes[i] > 0 && fishes[i] < 10 {
			fishes[i]--
		} else if fishes[i] == 0 {
			fishes[i] = 6
			AddFish(&fishes)
		}
	}
	return fishes
}
