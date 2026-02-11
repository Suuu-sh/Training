package main

import "fmt"

// Go文法メモ:
// - 関数定義の形: func 関数名(引数名 型) 戻り値型 { 処理 }
// - ポインタ型: *int（int を指すポインタ）
// - アドレス取得: &変数
// - デリファレンス: *ポインタ（参照先の値を読む/書く）
// - 値渡し: func f(v int) int
// - 参照渡し風: func f(p *int)

// main はポインタの基本を確認する。
func main() {
	// 初期値として整数を用意する。
	value := 10

	// value のメモリアドレスを取得して ptr に入れる。
	ptr := &value

	// value の値を表示する。
	fmt.Println("value:", value)
	// value のアドレスを表示する。
	fmt.Println("&value:", &value)
	// ptr が持つアドレスを表示する。
	fmt.Println("ptr:", ptr)
	// ptr の参照先の値を表示する。
	fmt.Println("*ptr:", *ptr)

	// 値渡し関数を呼んで、戻り値を received に受け取る。
	received := incrementByValue(value)
	// 値渡し後の元の value を表示する（変化しない）。
	fmt.Println("after incrementByValue, value:", value)
	// 値渡し関数の戻り値を表示する。
	fmt.Println("received from incrementByValue:", received)

	// ポインタ渡し関数を呼んで value を直接更新する。
	incrementByPointer(ptr)
	// ポインタ渡し後の value を表示する（変化する）。
	fmt.Println("after incrementByPointer, value:", value)
	// ptr の参照先の値も同じであることを表示する。
	fmt.Println("after incrementByPointer, *ptr:", *ptr)
}

// incrementByValue は値渡しで受け取った整数を 1 増やして返す。
func incrementByValue(v int) int {
	// 受け取った v を 1 増やす。
	v = v + 1
	// 更新した v を返す。
	return v
}

// incrementByPointer は参照先の整数を 1 増やす。
func incrementByPointer(p *int) {
	// nil ポインタを受けた場合に備えて先に判定する。
	if p == nil {
		// nil の場合は何もしないで終了する。
		return
	}
	// 参照先の値を 1 増やす。
	*p = *p + 1
}
