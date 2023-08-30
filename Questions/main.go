package main

import (
	"fmt"
	parent "GO_PRACTICE/sub"
)

func main(){

	f:=new(parent.Father)
	fmt.Println(f.Data("Kaushik"))
}