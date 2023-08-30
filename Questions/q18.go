package main

import "fmt"

var factVal int64 = 1 
var n int

func factorial(n int) int64 {   
    if(n < 0){
        fmt.Print("Factorial of negative number doesn't exist.")    
    }else{        
        for i:=1; i<=n; i++ {
            factVal *= int64(i) 
        }
    }    
    return factVal 
}

func main(){    
    fmt.Print("Enter a number: ")
    fmt.Scan(&n)   
    fmt.Print("Factorial is: ",factorial(n))
}