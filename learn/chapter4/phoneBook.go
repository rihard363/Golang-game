package main

import "fmt"

func main() {
    i:= 0
    fmt.Print(
      "1. Показать книгу"
      "2. Добавить контакт"
      "3.Удалить контакт"
      "4. Найти контакт"
      )
    fmt.Scanf("%i",&i)
    switch i{
      case 1: 
    }
    
    var name string
    fmt.Print("Введите имя: ")
    fmt.Scanf("%s", &name)
    var number string
    fmt.Print("Введите телефон: ")
    fmt.Scanf("%s", &number)
    entry := map[string]string{
        name: number,
    }
    
    for i, v:= range entry{
        fmt.Println( i,v)
    }
  
}