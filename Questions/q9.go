package main
 
import "fmt"

 
func Fibonacci(n int) int {
 
  if(n<=1){
 
    return n
}
  return Fibonacci(n-1) + Fibonacci(n-2)
 
}
 
func main() {
 
	var num int
	fmt.Println("enter a number")
	fmt.Scan(&num)
  for i := 1; i <= num; i = i + 1{
 
    fmt.Println(Fibonacci(i))
 
  }
 
}