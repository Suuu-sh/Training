package main

import "fmt"

// Go文法メモ:
// - 定数宣言: const 名前 型 = 値
// - 変数宣言: var 名前 型 = 値
// - 短縮宣言: 名前 := 値（関数内のみ）
// - 出力: fmt.Println(...), fmt.Printf("...", 値)
// - 書式指定: %T（型表示）, \n（改行）
// main は変数・定数・基本型のサンプルを表示する。
func main() {
	// アプリ名を定数として定義する。
	const appName string = "go-training"

	// 年齢を int 型で定義する。
	var age int = 30

	// 身長を float64 型で定義する。
	var height float64 = 170.5

	// インフラエンジニアかどうかを bool 型で定義する。
	var isInfraEngineer bool = true

	// 名前を短縮宣言で定義する（string 型）。
	name := "Youta"

	// 定数 appName を表示する。
	fmt.Println("app:", appName)

	// 名前と年齢を表示する。
	fmt.Println("name:", name, "age:", age)

	// 身長と職種フラグを表示する。
	fmt.Println("height:", height, "isInfraEngineer:", isInfraEngineer)

	// name の型を表示する。
	fmt.Printf("type(name)=%T\n", name)

	// age の型を表示する。
	fmt.Printf("type(age)=%T\n", age)

	// height の型を表示する。
	fmt.Printf("type(height)=%T\n", height)

	// isInfraEngineer の型を表示する。
	fmt.Printf("type(isInfraEngineer)=%T\n", isInfraEngineer)
}
