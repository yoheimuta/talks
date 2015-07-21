// すべてのコード（≒ ファイル）は何かしらのパッケージに属する
// main パッケージから処理ははじまる
package main

// I/O を扱う標準パッケージ fmt を使えるようにする
import "fmt"

// main パッケージの main() が呼び出される
func main() {
	// 文字列の型であることを宣言する
	var hello string = "Hello"

	// 関数内なら省略できる（暗黙的な型宣言）
	world := "世界"

	// 最後は改行して標準出力する
	fmt.Println(hello + ", " + world)

	// 書式文字列を使って標準出力する
	fmt.Printf("%s, %s", hello, world)
}
