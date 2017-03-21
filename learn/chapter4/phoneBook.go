package main

import "fmt"


func main() {
    var name string
    var number string
    entry := map[string]string{
        name: number,
    }
    var i int
    fmt.Print("1. Показать книгу")
    fmt.Print("2. Добавить контакт")
    fmt.Print("3.Удалить контакт")
    fmt.Print("4. Найти контакт")
    fmt.Scanf("%i",&i)
    switch i{
      case 0: listEntry(&entry)
      case 1: newEntry(&entry)
      case 2: deleteEntry(&entry)
      case 3: searchEntry(&entry)
      default: fmt.Println("Данной команды не существует")
    }
}
func listEntry(entry *map[string]string) {
    for i, v:= range *entry{
        fmt.Println( i,v)
    }
}
func newEntry(entry *map[string]string) (name,number string){
  fmt.Print("Введите имя: ")
  fmt.Scanf("%s", &name)
  fmt.Print("Введите телефон: ")
  fmt.Scanf("%s", &number)
  *entry = map[string]string{
      name: number,
    }
}
func deleteEntry(entry *map[string]string){
  var n int
  fmt.Print("Введите номер записи которую хотите удалить: ")
  fmt.Scanf("%i", &n)
  delete(*entry,n)
}
func searchEntry(entry *map[string]string){
 var search string
 fmt.Print("Введите имя которое нужно найти: ")
  fmt.Scanf("%s", &search)
 if name, ok := *entry["search"]; ok {    
    fmt.Println(search, ok)
} 
}
