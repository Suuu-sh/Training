package main

import (
	"fmt"
	"os"

	"go-training/internal/greet"
)

// Go文法メモ:
// - 関数定義の形: func 関数名(引数名 型) 戻り値型 { 処理 }
// - 例: func main() { ... }（main は戻り値なし）
// - 条件分岐: if 条件 { ... }
// - 引数参照: os.Args[1]（コマンドライン引数）
// - 文字列出力: fmt.Println(...)
func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Println(greet.Message(name))
}
