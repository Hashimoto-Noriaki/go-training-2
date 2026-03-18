package main

import "fmt"

func main(){
	primes := [6]int{2,3,5,7,11,13}
	s := primes[1:4]
	s[0] = 100
	fmt.Println("スライス:",s)
	fmt.Println("配列:",primes)

	a := []int{1,2,3}
	double(a)
	fmt.Println("関数で変更後:",a)
}

func double(s []int){
	for i := range s {
		s[i] *= 2
	}
}
