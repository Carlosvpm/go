package main

import "fmt"


func main(){
	a:=10
	
	
	fmt.Println(&a)

	var ponteiro *int = &a

	*ponteiro = 50

	fmt.Println(a)

}
