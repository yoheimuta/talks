package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func NewUser(first, last string) User {
	return User{first, last}
}

// 構造体型（厳密には type で定義した型や後述するインターフェース型も）にメソッドを定義できる
// (u *User) のことをメソッドレシーバ( method receiver )と言う
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	u := NewUser("陽平", "吉牟田")
	fmt.Println(u.FirstName, u.LastName)
	fmt.Println(u.FullName())
}
