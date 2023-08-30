package main

import "fmt"

var p=fmt.Println

func check(n int){
	if n<1 || n>10 {
		panic("wrong input")
	} else {
		p("correct input")
	}
}

func main(){
	var n int
	p("Enter a number between 1 to 10")
	fmt.Scan(&n)
	check(n)
}