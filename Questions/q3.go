// GO Program to take user input and addition of two strings

package main
import "fmt"

func main(){
	var string1 string
	var string2 string
	fmt.Println("Enter first string")
	fmt.Scan(&string1)
	fmt.Println("enter second string")
	fmt.Scan(&string2)
	fmt.Println(string1+" " +string2)
}