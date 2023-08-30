package main

import "fmt"

func main(){
	var a,b int
	fmt.Println("Enter value for A")
	fmt.Scan(&a)
	fmt.Println("Enter value for B")
	fmt.Scan(&b)
	fmt.Println("value of a is ",a," value for b is ",b)
	b = a + b
   	a = b - a
   	b = b - a
	fmt.Println("new value of a is ",a," new value for b is ",b)
}