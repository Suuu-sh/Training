# Go Syntax Summary

最終更新: 2026-02-11

## このファイルの目的
- ロードマップで完了した項目ごとに、実際に使ったGo文法を記録する。
- 次に進む前に「どの文法を使えたか」を短く確認する。

## 更新ルール
- `GO_TRAINING_ROADMAP.md` で項目を `[x]` にしたら、このファイルも同日に更新する。
- 追記単位は「完了した学習項目ごと」にする。

---

## 1) 開発環境セットアップ（完了）
対象: `go run`, `go test`, `go fmt`, `go vet`

- `go run ./cmd/hello Youta`
  - エントリポイント（`main`）を実行する。
- `go test ./...`
  - パッケージ単位でテストを実行する。
- `go fmt ./...`
  - Go標準フォーマッタで整形する。
- `go vet ./...`
  - よくあるミスを静的チェックする。

---

## 2) 変数・定数・基本型（完了）
主な実装: `cmd/basics/main.go`

- `const` 定数宣言
  - `const appName string = "go-training"`
- `var` による変数宣言
  - `var age int = 30`
  - `var height float64 = 170.5`
  - `var isInfraEngineer bool = true`
- `:=` による短縮宣言（関数内）
  - `name := "Youta"`
- 基本型
  - `string`, `int`, `float64`, `bool`
- 出力
  - `fmt.Println(...)`
  - `fmt.Printf("type(name)=%T\n", name)`（`%T` は型表示、`\n` は改行）

---

## 3) 制御構文（if/switch/for）（完了）
主な実装: `cmd/control/main.go`

- `if / else if / else`
  - 年齢で分岐（未成年/成人/シニア）
- `switch`（条件式なし）
  - `switch { case ... }` 形式で条件分岐
- `for`（3パターン）
  - カウントアップ: `for i := 1; i <= 5; i++`
  - while風: `for count < 3`
  - 早期終了: `return`（入力不正時）
- コマンドライン引数
  - `os.Args`
  - `len(os.Args)` で引数数を判定
  - `strconv.Atoi(...)` で文字列→整数変換

---

## 4) 関数（多値返却、defer）（完了）
主な実装: `cmd/functions/main.go`

- 関数定義
  - `func run() error`
  - `func parseTwoInts(args []string) (int, int, error)`
  - `func safeDivide(left int, right int) (int, error)`
- 多値返却
  - `left, right, err := parseTwoInts(os.Args)`
  - `result, err := safeDivide(left, right)`
- エラー処理の基本
  - `if err != nil { return ... }`
  - `errors.New(...)`
  - `fmt.Errorf("...: %w", err)`（wrap）
- `defer`
  - `defer fmt.Println("defer: run finished")`
  - 関数終了時に必ず実行されることを確認

---

## 5) ポインタの基本（完了）
主な実装: `cmd/pointer/main.go`

- ポインタ型
  - `*int`（int を指すポインタ型）
- アドレス取得
  - `ptr := &value`
- デリファレンス（参照先の値の読み書き）
  - 読み取り: `fmt.Println(*ptr)`
  - 書き込み: `*p = *p + 1`
- 値渡しと参照渡し風の違い
  - 値渡し: `func incrementByValue(v int) int`
  - 参照渡し風: `func incrementByPointer(p *int)`
- nil チェック
  - `if p == nil { return }`

---

## 6) 先取りで使った文法（補足）
主な実装: `cmd/hello/main.go`, `internal/greet/*`

- package/import
  - `package main`, `package greet`
  - 複数importと内部パッケージ呼び出し
- 別パッケージ関数の利用
  - `greet.Message(name)`
- テストの基本
  - `func TestXxx(t *testing.T)`
  - table-driven testの形（`[]struct{...}`）
  - サブテスト（`t.Run(...)`）

---

## 次回更新予定
- `構造体・メソッド` を完了したら、`struct` 定義とメソッド定義を追記する。
