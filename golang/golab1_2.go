package main

import (
 "fmt"
 "math"
)

func main() {
 a := 0.05
 b := 0.06
 deltaX := 0.15
 points := []float64{0.15, 0.26, 0.37, 0.48, 0.56}

 for _, x := range points {
  xSquaredMinusBSquared := math.Pow(x, 2) - math.Pow(b, 2)
  xSquaredMinusASquared := math.Pow(x, 2) - math.Pow(a, 2)
  result := math.Acos(xSquaredMinusBSquared) / math.Asin(xSquaredMinusASquared)
  fmt.Printf("x=%.2f, y=%.4f\n", x, result)
 }
}