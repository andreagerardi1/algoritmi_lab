package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var s []string
	var flag bool
	var vocale string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	for _, l := range s {
		flag = false
		for _, k := range l {
			if k == 'a' || k == 'e' || k == 'i' || k == 'o' || k == 'u' {
				vocale = string(k) 
				flag = true
				break
			}
		}
		if flag == true {
			fmt.Println("la prima vocale di ", l, " è: ", vocale)
		} else {
			fmt.Println(l, " non contiene nessuna vocale")
		}
	}
}
