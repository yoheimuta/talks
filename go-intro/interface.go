package main

import "fmt"

type Duck interface { // インターフェース型( interface type )は、メソッド群で定義される
	Quack() // その型がどのようなメソッドを実装すべきかを規定する役割を持つ（このケースでは Quack()）
}

type Donald struct {
}

func (d Donald) Quack() { // ドナルド・ダックさんはガーガー鳴く
	fmt.Println("quack quack!")
}

type Daisy struct {
}

func (d Daisy) Quack() { // デイジーダックさんもガーガー鳴く
	fmt.Println("~quack ~quack")
}

type Dog struct {
}

func (d Dog) Bark() { // 犬はガーガー鳴いたりしない
	fmt.Print("woof woof")
}

func sayQuack(duck Duck) { // Quack メソッドを実装している型を引数にとる
	duck.Quack()
}

func main() {
	donald := Donald{}
	sayQuack(donald) // quack quack!（implements などの明示的な宣言がないところがポイント）

	daisy := Daisy{}
	sayQuack(daisy) // ~quack ~quack（implements などの明示的な宣言がないところがポイント）

	dog := Dog{}
	sayQuack(dog) // コンパイルエラーになる（ランタイムエラーではないところがポイント）
}
