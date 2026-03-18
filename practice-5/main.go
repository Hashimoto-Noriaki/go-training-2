package main

import "fmt"

func main(){
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("s[:3] =",s[:3])
	fmt.Println("s[2:] =",s[2:])
	fmt.Println("s[:] =",s[:])

	data := []string{"A", "B", "C"}
	fmt.Println("先頭以外:",data[1:],"末尾以外:",data[:2])
}
