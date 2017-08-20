package main

import "fmt"

func main() {
	fmt.Print("Enter a value ")
	var F float64
	var M float64
	fmt.Scanf("%f", &F)
	M= F*0.3048
	fmt.Println(M)
}
