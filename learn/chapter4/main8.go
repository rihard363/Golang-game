package main

import "fmt"
func makeOddGenerator() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = 2*i+1
        i += 1
        return
    }
}
func main() {
    nextOdd := makeOddGenerator()
    fmt.Println(nextOdd()) // 1
    fmt.Println(nextOdd()) // 3
    fmt.Println(nextOdd()) // 5
}