package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Go文法メモ:
// - 関数定義: func 関数名(引数名 型) 戻り値型 { ... }
// - 多値返却: func f() (int, int, error)
// - defer: 関数終了時に後処理を実行する。
// - エラー判定: if err != nil { ... }
// - エラーwrap: fmt.Errorf("文脈: %w", err)

// main は run の実行結果を受け取り、エラーがあれば表示する。
func main() {
	// run を実行して、処理全体のエラーを受け取る。
	err := run()
	// エラーがある場合だけメッセージを表示する。
	if err != nil {
		// エラー内容を標準出力に表示する。
		fmt.Println("error:", err)
	}
}

// run は引数の解析、割り算、結果表示を順番に実行する。
func run() error {
	// run が終了するときに必ず実行されるログを登録する。
	defer fmt.Println("defer: run finished")

	// コマンドライン引数から 2 つの整数を取得する。
	left, right, err := parseTwoInts(os.Args)
	// 取得に失敗したら、ここでエラーを返す。
	if err != nil {
		// 呼び出し元で原因が分かるように文脈付きで返す。
		return fmt.Errorf("parse arguments: %w", err)
	}

	// 2 つの整数で割り算を実行する。
	result, err := safeDivide(left, right)
	// 割り算に失敗したら、ここでエラーを返す。
	if err != nil {
		// 呼び出し元で原因が分かるように文脈付きで返す。
		return fmt.Errorf("divide values: %w", err)
	}

	// 入力値と計算結果を表示する。
	fmt.Printf("%d / %d = %d\n", left, right, result)
	// 正常終了なので nil を返す。
	return nil
}

// parseTwoInts は引数から 2 つの整数を取り出して返す。
func parseTwoInts(args []string) (int, int, error) {
	// 引数不足を判定して早期 return する。
	if len(args) < 3 {
		// 使い方が分かるエラーを返す。
		return 0, 0, errors.New("usage: go run ./cmd/functions <left> <right>")
	}

	// 1 つ目の値を文字列から整数へ変換する。
	left, err := strconv.Atoi(args[1])
	// 変換失敗を判定する。
	if err != nil {
		// 失敗理由を含めたエラーを返す。
		return 0, 0, fmt.Errorf("left is not int: %w", err)
	}

	// 2 つ目の値を文字列から整数へ変換する。
	right, err := strconv.Atoi(args[2])
	// 変換失敗を判定する。
	if err != nil {
		// 失敗理由を含めたエラーを返す。
		return 0, 0, fmt.Errorf("right is not int: %w", err)
	}

	// 正常時は 2 つの整数と nil エラーを返す。
	return left, right, nil
}

// safeDivide はゼロ除算を防ぎながら整数同士の割り算を行う。
func safeDivide(left int, right int) (int, error) {
	// ゼロ除算を先に判定する。
	if right == 0 {
		// 計算不可なのでエラーを返す。
		return 0, errors.New("division by zero")
	}

	// 整数除算を実行して返す。
	return left / right, nil
}
