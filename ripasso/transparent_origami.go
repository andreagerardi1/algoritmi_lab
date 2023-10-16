package main


import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Dot struct {
	x,y int
}

func main() {
	flag := true
	var dots []Dot
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			flag = false
		}
		if flag == true { //sto leggendo punti, li salvo in dots
			sl := strings.Split(scanner.Text(), ",")
			a, _ := strconv.Atoi(sl[0])
			b, _ := strconv.Atoi(sl[1])
			dots = append(dots, Dot{x: a, y: b})
		} else { //sto leggendo dove foldare 
			sl := strings.Fields(scanner.Text())
			aux := strings.Split(sl[2], "=")
			if aux[0] == "y" {
				FoldY(aux[1], dots)
			} else if aux[0] == "x" {
				FoldX(aux[1], dots)
			}	
		}
	}
}

func FoldY(ord int, dots []Dot) {
	
}

func FoldX(asc int, dots []Dot) {

}
