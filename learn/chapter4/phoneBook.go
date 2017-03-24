package main

import "fmt"


func main() {
    var name string
    var number string
    entry := map[string]string{
        name: number,
    }
  Loop:
    for {
      var i int
      fmt.Println("1. Показать книгу")
      fmt.Println("2. Добавить контакт")
      fmt.Println("3. Удалить контакт")
      fmt.Println("4. Найти контакт")
      fmt.Println("5. Выход")
      fmt.Scanf("%d",&i)
      switch i{
      case 1: listEntry(&entry)
      case 2: newEntry(&entry)
      case 3: deleteEntry(&entry)
      case 4: searchEntry(&entry)
      case 5: break Loop
      default: fmt.Println("Данной команды не существует")
        }
    }
  }
func listEntry(entry *map[string]string) {
    for i, v:= range *entry{
        fmt.Println( i,v)
    }
}
func newEntry(entry *map[string]string){
  var name string
  fmt.Print("Введите имя: ")
  fmt.Scanf("%s", &name)
  var number string
  fmt.Print("Введите телефон: ")
  fmt.Scanf("%s", &number)
  (*entry)[name]= number
  return 
}
func deleteEntry(entry *map[string]string){
  var n string
  fmt.Print("Введите номер записи которую хотите удалить: ")
  fmt.Scanf("%s", &n)
  delete(*entry,n)
  return 
}
func searchEntry(entry *map[string]string){
 var search string
 fmt.Print("Введите имя которое нужно найти: ")
  fmt.Scanf("%s", &search)
 if name, ok := (*entry)[search]; ok {    
    fmt.Println(name)
} 
}

