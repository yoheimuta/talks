package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://golang.org/",
		"http://fout.co.jp/",
		"http://mtburn.jp/",
	}
	for _, url := range urls {
		go func(url string) { // ゴルーチンを起動する
			res, err := http.Get(url) // URL 先のコンテンツを取得する
			if err != nil {
				return
			}
			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			fmt.Printf("------------\n%s\n------------\n", body)
		}(url)
	}

	time.Sleep(10 * 1000 * time.Millisecond) // ゴルーチンの完了を適当に待つ
	fmt.Println("DONE")
}
