package main

import "fmt"


func main() {
    var n int
    fmt.Println("Ведите длину среза")
    fmt.Scanf("%d",&n)
    var x[]float64
    c:=cap(x)
    for len(x)< n {
      var i float64
      fmt.Println("Введите число")
      fmt.Scanf("%f",&i)
      x = append(x,i)
      if cap(x)!=c{
        c=cap(x)
        fmt.Println(x)
        }
      }
    for j=1,j<len(x),j++ {
      for k=0,k<(len(x)-1),k++ {
        if x[k]>x[k+1]{
           z = x[k]
           x[k] = x[k + 1]
           x[k + 1] = z
          }
        }
      fmt.Println(x)
      } 
    }
  

