package main

import "fmt"

func main() {
	var vklad, perc, sum float64
	fmt.Println("Введите значения сумму вклада и процент")
	fmt.Scan(&vklad, &perc)
	years := 5
	sum = vklad
	for i := 0; i < years; i++ {
		sum = sum * ((100 + perc) / 100)
	}
	fmt.Println("Сумма вклада через ", years, "лет составляет:", sum)
}
