package main

import (
	"fmt"
	"math"
)

var D float64

func roots(a, b, c float64) (float64, float64) {
	D = b*b - 4*a*c
	fmt.Println("Дискриминант равен ", D)
	if D > 0 {
		return -b + math.Sqrt(D), -b - math.Sqrt(D)
	} else {
		return 0, 0
	}
}
