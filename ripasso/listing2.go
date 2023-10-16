package main

import "fmt"

func main() {
	var arr = []int{1, 2, 5, 7, -2, 10, 9, 3, 8}
	fmt.Println(largest(arr))
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func largest ( numbers []int) int {
n := len( numbers )
if n == 1 {
return numbers[0]
}
return max ( numbers[n -1] , largest(numbers[:n-1]))
}

//la prima a terminare è (1,2), che sono gli argomenti, restituisce 2.
// la prima volta viene eseguita con 2,1
// l'ultima volta viene eseguita con 8, 21


