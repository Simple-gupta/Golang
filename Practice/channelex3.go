package main
import "fmt"

func fun1(c chan int){
	num:= <-c
	fmt.Println(num)
}
func fun(){
	c := make(chan int)
	go fun1(c)
	c <- 10
	close(c)
}
func main(){
	fun()
}