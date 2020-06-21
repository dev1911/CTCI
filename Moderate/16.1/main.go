package main

import(
	"fmt"
)

func swap(a int , b int)(int , int){
	a = a - b
	b = a + b
	a = b - a

	return a , b
}

func main(){
	a := 5
	b := 6

	fmt.Printf("Values before swapping : %d , %d\n" , a , b)
	a , b = swap(a , b)
	fmt.Printf("Values after swapping : %d , %d\n" , a , b)
}