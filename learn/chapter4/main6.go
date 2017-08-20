package main

import "fmt"
func main() {
	var x int = 0
	evenOdd := func() (int,bool){
		x++
		if x%2>0{
			return int(x)/2,false
		}else{
			return int(x)/2,true
		}
	}
	fmt.Println(evenOdd())
	fmt.Println(evenOdd())
	fmt.Println(evenOdd())
	fmt.Println(evenOdd())
	fmt.Println(evenOdd())
}