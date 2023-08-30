package main

import "fmt"

func main(){
	var n int
	fmt.Println("Enter a number")
	fmt.Scan(&n)
	for i:=1; i<=10; i++{
		fmt.Println(n," x ", i," = ",n*i)
	}
}