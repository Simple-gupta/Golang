package main

import "fmt"

func main(){
	var n int
	var result int = 0
	fmt.Println("Enter the number")
	fmt.Scan(&n)
	result=((n*(n+1))/2)
	fmt.Println(result)
}