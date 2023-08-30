package main
import (
   "fmt"
   "strings"
)
func main() {
   var a1 = "hi"
 var a2 = "hello"
   var a3 = "good morning"
   var a4 = "yes"
   
   fmt.Println(strings.Compare(a1,a2))
   
   fmt.Println(strings.Compare(a2, a1))
   
   fmt.Println(strings.Compare(a3, a4))
   
   fmt.Println(strings.Compare(a4, a3))
   
   fmt.Println(strings.Compare(a1, a1))
}