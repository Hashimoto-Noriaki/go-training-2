package main

import "fmt"

func main(){
	//関数とポインタ(値渡し、ポインタ渡し)
	x := 10
	doubleValue(x)
	fmt.Println("値渡し後: x =", x)

	doublePointer(&x)
	fmt.Println("ポインタ渡し後: x =", x)
}

func doubleValue(n int) {n = n * 2}
func doublePointer(n *int) { *n = *n *2 }
