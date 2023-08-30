package main

import (
"fmt"
)

func main() {
	var name map[int]string
	name = make(map[int] string)
	name[1]="Simple"
	name[2]="Supriya"
	name[3]="Aryan"
	name[4]="Aditya"
	name[5]="Manohar"
	name[6]="Vishwas"
	name[7]="Kola"
	for i, j := range name {
		fmt.Printf("Key: %d Value: %s\n", i, j)
	}
}
