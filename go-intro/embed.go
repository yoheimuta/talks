package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// User 構造体を埋め込んだ構造体をつくる
type Corp struct {
	CorpName string
	User
}

func NewCorp(corp, first, last string) Corp {
	return Corp{CorpName: corp, User: User{first, last}}
}

func main() {
	u := NewCorp("M.T.Burn", "yohei", "yoshimuta")
	fmt.Println(u.FirstName, u.LastName, u.CorpName)
	fmt.Println(u.FullName())
}
