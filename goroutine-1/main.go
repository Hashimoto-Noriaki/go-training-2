//ゴルーチン
package main

import(
	"fmt"
	"time"
) 

//お茶を沸かす(時間がかかる想定)
func makeTea(){
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("お茶を沸かしてます...")
	}
}

//カップラーメンにお湯を入れる
func makeNoodle(){
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("ラーメンのお湯を沸かして...")
	}
}

func main(){
	go makeTea()//お茶を沸かしておく
	makeNoodle()//その間にラーメンのお湯を沸かす
}
