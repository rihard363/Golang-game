package main

import "fmt"
func main() {
	 x := 0
	evenOdd := func() (float32,bool){
		x++
		if x%2>0{
			return float32(x)/2,false
		}else{
			return float32(x)/2,true
		}
	}
	fmt.Println(evenOdd())
	fmt.Println(evenOdd())
	fmt.Println(evenOdd())
	fmt.Println(evenOdd())
	fmt.Println(evenOdd())
}