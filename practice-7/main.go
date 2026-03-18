package main

import "fmt"

func main(){
	var s []int
	fmt.Println("s =", s, "len =", len(s), "s == nil:", s == nil)
	s = append(s,1)
	fmt.Println("append後:", s, "s == nil:", s == nil)

	var nilSlice []int
	emptySlice := []int{}
	fmt.Println("nil:", nilSlice == nil, "空:", emptySlice == nil)
}
