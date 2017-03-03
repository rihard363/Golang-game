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
		fmt.Println(searchMax(x...))
}
	func searchMax(x ...int) int{
		max:=x[0]
		for _,y:=range x{
			if y>max{
				max=y
			}
		}
		return max
}