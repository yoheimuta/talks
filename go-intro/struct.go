package main

import "fmt"

// 構造体型は、フィールドの集まりで複数のデータを 1 つにまとめられる
// type 宣言は、既存の型を拡張して、独自の型を定義できる
type User struct {
	FirstName string
	LastName  string
}

// Go には構造体のコンストラクタにあたる構文はないので、New ではじまる関数を定義するのが慣習
func NewUser(first, last string) User {
	return User{first, last}
}

func main() {
	u := NewUser("yohei", "yoshimuta")
	fmt.Println(u.FirstName, u.LastName)
}
