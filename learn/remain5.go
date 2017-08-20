package main

import "fmt"

func main() {
		 x := []int{
	 	33, 34, 56,
		78, 56, 45,
		24, 98, 11,
		45, 79, 99,
		999, 567, 567,
	}
		min:=x[0]
		for _,y:=range x{
			if y<min{
				min=y
			}
	}
		fmt.Println(min)
}
