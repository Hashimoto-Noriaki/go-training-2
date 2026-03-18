package main

import "fmt"

func main(){
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("初期: len=%d cap=%d %v\n", len(s), cap(s),s)
	s = s[:3]
	fmt.Printf("s[:3]: len=%d cap=%d %v\n", len(s), cap(s),s)
	s = s[:5]
	fmt.Printf("s[:5]: len=%d cap=%d %v\n", len(s), cap(s),s)
}
