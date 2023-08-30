// GO Program to Calculate Area of Rectangle and Square

package main
import "fmt"

func main(){
	fmt.Printf( "Book 1 title : %s\n", Book1.title)
	var area float32
	var l float32
	var b float32
	fmt.Println("For rectangle enter length")
	fmt.Scan(&l)
	fmt.Println("enter breadth")
	fmt.Scan(&b)
	area=l*b
	fmt.Println("area=",area)
	fmt.Println("for square enter length")
	fmt.Scan(&l)
	area=l*l
	fmt.Println("area=",area)
}