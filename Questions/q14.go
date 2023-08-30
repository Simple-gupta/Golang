package main

import "fmt"

func main(){
	var rows int
	var temp int = 1
	fmt.Print("Enter rows")
	fmt.Scan(&rows)

	for i := 1; i <= rows; i++ {	
		
		for j := 1; j <= i; j++ {

			fmt.Printf(" %d",temp)				
			temp++
		}
		fmt.Println("")		
	}

}
