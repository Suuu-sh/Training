package main

import "fmt"

// Go文法メモ:
// - 構造体定義: type 型名 struct { フィールド名 型 }
// - インスタンス生成: 変数 := 型名{フィールド: 値}
// - メソッド定義（値レシーバ）: func (r 型名) メソッド名(...) 戻り値 { ... }
// - メソッド定義（ポインタレシーバ）: func (r *型名) メソッド名(...) 戻り値 { ... }
// - 値レシーバはコピーに対して動作し、ポインタレシーバは元データを更新できる。

// Profile は学習者プロフィールを表す構造体。
type Profile struct {
	// Name は名前を保持する。
	Name string
	// Years は学習年数ではなく、ここでは年齢を保持する。
	Years int
}

// Intro はプロフィール情報を文字列で返す（値レシーバ）。
func (p Profile) Intro() string {
	// 紹介文をフォーマットして返す。
	return fmt.Sprintf("name=%s, years=%d", p.Name, p.Years)
}

// RenameByValue は値レシーバで名前を変更する（元データは変わらない）。
func (p Profile) RenameByValue(newName string) {
	// レシーバのコピー側の Name を更新する。
	p.Name = newName
	// コピー側の結果を表示する。
	fmt.Println("inside RenameByValue:", p.Intro())
}

// RenameByPointer はポインタレシーバで名前を変更する（元データが変わる）。
func (p *Profile) RenameByPointer(newName string) {
	// 受け取ったポインタが nil の場合は何もしない。
	if p == nil {
		// nil のときは早期 return する。
		return
	}
	// 元データの Name を更新する。
	p.Name = newName
}

// AddYear はポインタレシーバで年齢を 1 増やす。
func (p *Profile) AddYear() {
	// 受け取ったポインタが nil の場合は何もしない。
	if p == nil {
		// nil のときは早期 return する。
		return
	}
	// 元データの Years を 1 増やす。
	p.Years++
}

// main は構造体とメソッドの基本を確認する。
func main() {
	// Profile のインスタンスを生成する。
	profile := Profile{
		// 初期の名前を設定する。
		Name: "Youta",
		// 初期の年齢を設定する。
		Years: 30,
	}

	// 初期状態を表示する。
	fmt.Println("initial:", profile.Intro())

	// 値レシーバのメソッドを呼び出す（元データは変わらない）。
	profile.RenameByValue("Taro")
	// 値レシーバ呼び出し後の状態を表示する。
	fmt.Println("after RenameByValue:", profile.Intro())

	// ポインタレシーバのメソッドを呼び出す（元データが変わる）。
	profile.RenameByPointer("Hanako")
	// 名前変更後の状態を表示する。
	fmt.Println("after RenameByPointer:", profile.Intro())

	// ポインタレシーバで年齢を 1 増やす。
	profile.AddYear()
	// 年齢更新後の状態を表示する。
	fmt.Println("after AddYear:", profile.Intro())
}
