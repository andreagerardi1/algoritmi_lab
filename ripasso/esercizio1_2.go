package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
  "math"
)

func main() {
  str := os.Args[1]
  split1 := strings.Split(str, "x2")
  split2 := strings.Split(split1[1], "x")
  split3 := strings.Split(split2[1], "=0")

  a, _ := strconv.ParseFloat(split1[0], 64)
  b, _ := strconv.ParseFloat(split2[0], 64)
  c, _ := strconv.ParseFloat(split3[0], 64)

  delta := (b * b) - 4 * (a * c)
  if delta < 0 {
    fmt.Println("zero soluzioni reali")
  } else if delta == 0 {
    x1 := -b / (2 * a)
    fmt.Printf("una sola soluzione: %.3f", x1)
  } else {
    x1 := (-b + math.Sqrt(delta)) / (2 * a)
    x2 := (-b - math.Sqrt(delta)) / (2 * a)
    fmt.Printf("le due soluzioni reali sono: %.3f, %.3f", x1, x2)
  }
}
