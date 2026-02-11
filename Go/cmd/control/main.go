package main

import (
	"fmt"
	"os"
	"strconv"
)

// Go文法メモ:
// - if / else if / else で条件分岐する。
// - switch { case 条件: ... default: ... } で条件分岐する。
// - for 初期化; 条件; 後処理 { ... } と for 条件 { ... } の2形がある。
// - 型変換: strconv.Atoi(文字列) -> (int, error)
// - 早期終了: return
// main は if/switch/for の基本動作をまとめて確認する。
func main() {
	// 年齢引数が1つ必須なので、引数不足を先に判定する。
	if len(os.Args) < 2 {
		// 引数が足りない場合の使い方を表示する。
		fmt.Println("使い方: go run ./cmd/control <年齢>")
		// 必須引数不足なので処理を終了する。
		return
	}

	// 先頭引数を整数に変換する。
	age, err := strconv.Atoi(os.Args[1])
	// 変換に失敗した場合の分岐を処理する。
	if err != nil {
		// 失敗時のメッセージを表示する。
		fmt.Println("年齢は整数で指定してください。例: go run ./cmd/control 25")
		// 不正入力時は処理を終了する。
		return
	}

	// age を if / else if / else で分類する。
	if age < 20 {
		// 20未満の分岐を表示する。
		fmt.Println("if: 未成年のカテゴリです。")
	} else if age < 65 {
		// 20以上65未満の分岐を表示する。
		fmt.Println("if: 成人のカテゴリです。")
	} else {
		// 65以上の分岐を表示する。
		fmt.Println("if: シニアのカテゴリです。")
	}

	// switch でも同じ分類を表現する。
	switch {
	// 20未満の条件を判定する。
	case age < 20:
		// switch の 20未満分岐を表示する。
		fmt.Println("switch: age < 20")
	// 65未満の条件を判定する。
	case age < 65:
		// switch の 20以上65未満分岐を表示する。
		fmt.Println("switch: 20 <= age < 65")
	// それ以外の条件を処理する。
	default:
		// switch の 65以上分岐を表示する。
		fmt.Println("switch: age >= 65")
	}

	// 1から5までの合計を保持する変数を初期化する。
	total := 0
	// for 文で 1 から 5 まで繰り返す。
	for i := 1; i <= 5; i++ {
		// 現在の i を total に加算する。
		total += i
		// 各ループの途中経過を表示する。
		fmt.Printf("for: i=%d total=%d\n", i, total)
	}

	// for を while 風に使うためのカウンタを初期化する。
	count := 0
	// count が 3 未満の間だけ繰り返す。
	for count < 3 {
		// 現在の count を表示する。
		fmt.Printf("for-while: count=%d\n", count)
		// ループ終了に向けて count を増やす。
		count++
	}
}
