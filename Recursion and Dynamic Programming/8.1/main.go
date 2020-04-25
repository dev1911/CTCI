package main

import(
	"fmt"
)

//For memoization
var cache map[int]int

func noOfStepsCached(n int , cache map[int]int) int{
	if _ , prs := cache[n]; prs{
		return cache[n]
	}
	result := noOfStepsCached(n-1 , cache) + noOfStepsCached(n-2 , cache) + noOfStepsCached(n-3 , cache) 
	cache[n] = result
	return result
}

func main(){
	cache = make(map[int]int)
	cache[1] = 1
	cache[2] = 2
	cache[3] = 4
	n:=1000
	fmt.Printf("Number of ways to reach %d steps is %d\n" , n , noOfStepsCached(n , cache))
}