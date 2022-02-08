package main

import "fmt"

func main() {
	var a, b, c float64
	fmt.Println("Введите значения a,b,c")
	fmt.Scan(&a, &b, &c)
	fmt.Println(a, b, c)
	var x1, x2 float64
	x1, x2 = roots(a, b, c)
	if D > 0 {
		fmt.Println("Корни уравнения:", x1, " ", x2)
	} else {
		fmt.Println("У данного уравнения нет корней")
	}
}
