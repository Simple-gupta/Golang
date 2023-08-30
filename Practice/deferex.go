package main

import "fmt"

var p=fmt.Println

func main(){
	//defer executes last
	//multiple defer LIFO
	p("first")
	defer p("second")
	defer p("third")
	defer p("fourth")
	p("fifth")

}

