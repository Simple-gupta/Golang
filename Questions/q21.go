package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println

func main() {

    p("Contains:  ", s.Contains("hello", "he"))
	p("Containsany:  ", s.ContainsAny("hello", "h"))
	p("Count:  ", s.Count("hellohi", "h"))
	p("Equalfold:  ", s.EqualFold("hellohi", "hElloHi"))

}