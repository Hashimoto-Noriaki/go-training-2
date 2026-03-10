package main

import "fmt"

func main(){
	//ポインタの基本
	i := 56
	p := &i
	fmt.Println("i =", i, "*p =", *p)
	*p = 100
	fmt.Println("ポインタ渡し後: i =",i)
}
