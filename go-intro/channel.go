package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum // 合計値をチャネル c に送る
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // チャネルを初期化する

	go sum(a[:len(a)/2], c) // インデックス 0 1 2 の合計値の計算をする
	go sum(a[len(a)/2:], c) // インデックス 3 4 5 の合計値の計算をする

	x, y := <-c, <-c // 合計値の計算を待って受け取る（タイマーは使ってない）

	fmt.Println(x, y)
}
