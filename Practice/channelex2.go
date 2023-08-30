package main
import "fmt"

func fun(c chan int){
	num := <- c
	fmt.Println(num)
}

func main(){
	c := make(chan int)
	go fun(c)
	c <- 1
	close(c)
}