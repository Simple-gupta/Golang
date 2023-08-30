package main

import "fmt"

func main(){
	var num int
	fmt.Println("Enter a number")
	fmt.Scan(&num)
	if(num==0){
		fmt.Println("zero")
	} else if(num%2==0){
		fmt.Println("even")
	} else{
		fmt.Println("odd")
	}
}