package main
import "fmt"


func avgOfArray(array[] int, n int) float64 {
   var sum int = 0
   var i int = 0
   for i = 0; i < n; i++ {
      sum += (array[i])
   }
   var avg float64 = float64(sum) / float64(n)
   return avg
}
func main() {
	var n int
	fmt.Println("enter number of elements")
	fmt.Scan(&n)
	var arr1=make([]int, n)
	fmt.Println("enter elements")
	for i := 0; i<n; i++{
		fmt.Scan(&arr1[i])
	}


   result := avgOfArray(arr1, n)
   fmt.Println("result:= ", result)
}