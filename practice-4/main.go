package main

import "fmt"

func main(){
	numbers := []int{1, 2, 3, 4, 5}
	names := []string{"太郎","花子"}
	fmt.Println("数値:",numbers,"名前:",names)

	matrix := [][]int{{1, 2}, {3, 4}}
	fmt.Println("2次元:",matrix)

	s1 := []int{}
	var s2 []int
	fmt.Println("s1 == nil:", s1 == nil, "s2 == nil:", s2 == nil)
}
