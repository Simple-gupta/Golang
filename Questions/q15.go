
package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{1,2,3,5,4,7,6}
	fmt.Println(num)	
	if sort.IntsAreSorted(num)==false{
		sort.Ints(num)	
	}
	fmt.Println(num)

	text := []string{"Simple","Aryan","Somesh"}
	fmt.Println(text)	
	if sort.StringsAreSorted(text)==false{
		sort.Strings(text) 
	}
	fmt.Println(text)

	val := []float64{2.5,6.3,1.5,9.8,4.7,1.1,5.2}
	fmt.Println(val)	
	if sort.Float64sAreSorted(val)==false{
		sort.Float64s(val) 
	}
	fmt.Println(val)
}