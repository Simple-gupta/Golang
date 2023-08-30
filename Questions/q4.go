//GO Program to Find LCM and GCD of given two numbers

package main

import "fmt"

func main() {

   var n1 int
   var n2 int
   var result int
   var result2 int
   fmt.Println("enter number 1")
	fmt.Scan(&n1)
	fmt.Println("enter number 2")
	fmt.Scan(&n2)	
   result=hcf(n1,n2)
   fmt.Println("GCD OF", n1, n2,"is",result)
   result2=lcm(n1,n2,result)
   fmt.Println("LCM OF", n1, n2,"is",result2)

}

func hcf(n1 int, n2 int) int {
   if (n2 != 0) {
      return hcf(n2, n1 % n2);
   } else {
      return n1;
   }
}

func lcm(n1 int, n2 int, result int) int {
	return (n1*n2)/result
}

