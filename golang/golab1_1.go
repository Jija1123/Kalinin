package main

import (
 "fmt"
 "math"
)

func main() {
 a := 0.05
 b := 0.06
 deltaX := 0.15
 startX := 0.2
 endX := 0.95

 for x := startX; x <= endX; x += deltaX {
  xSquaredMinusBSquared := math.Pow(x, 2) - math.Pow(b, 2)
  xSquaredMinusASquared := math.Pow(x, 2) - math.Pow(a, 2)
  result := math.Acos(xSquaredMinusBSquared) / math.Asin(xSquaredMinusASquared)
  fmt.Printf("x=%.2f, y=%.4f\n", x, result)
 }
}