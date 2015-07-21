package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := make(chan []byte) // 別ゴルーチンの URL 先取得結果を main 関数で受け取るためのチャネル

	urls := []string{
		"http://golang.org/",
		"http://fout.co.jp/",
		"http://mtburn.jp/",
	}
	for _, url := range urls {
		go func(url string, c chan []byte) {
			res, err := http.Get(url)
			if err != nil {
				return
			}
			body, err := ioutil.ReadAll(res.Body)
			res.Body.Close()

			c <- body // body := <-c に向けて URL 先取得結果を送信する
		}(url, c)
	}

	t := time.After(1 * time.Second) // 返り値は <-chan Time

	for {
		select {
		case body := <-c: // c <- body からの URL 先取得結果を受信する
			fmt.Printf("------------\n%s\n------------\n", body)
		case <-t: // 指定した秒数が経過すると現在時刻が送信されるので受信する
			fmt.Println("TIMEOUT")
			return // タイムアウトのため終了する
		}
	}
}
