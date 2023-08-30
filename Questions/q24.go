package main

import "fmt"

func main(){
	var num,remainder int
	fmt.Println("Enter a number")
	fmt.Scan(&num)
	input_num := num
	res := 0
   for num>0 {
      remainder = num % 10
      res = (res * 10) + remainder
      num = num / 10
   }
   if input_num == res {

	  fmt.Println("Palindrome")
   } else {
	fmt.Println("Not Palindrome")
   }
}