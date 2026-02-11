# Go Training Roadmap & Progress

最終更新: 2026-02-11

## 学習者プロフィール
- バックグラウンド: `Infra Engineer`
- Go経験: `ほぼ未経験`
- 学習方針: `基礎を固めつつ、運用を意識した実装力を獲得`

## 使い方
- 完了した項目は `[x]` に変更
- 日付と成果を `進捗ログ` に追記
- 毎週末に `次週の目標` を更新
- 項目を完了したら `GO_SYNTAX_SUMMARY.md` に使用した文法を追記

## 進捗サマリー
- 全体進捗: `5 / 32` 完了
- 現在のフェーズ: `Phase 1: 基礎`
- 今週の目標: `Goの土台を作る（文法・実行・テストの入口）`

---

## Phase 1: Go基礎
- [x] 開発環境セットアップ（Go, fmt, lint, test実行）
- [x] 変数・定数・基本型
- [x] 制御構文（if/switch/for）
- [x] 関数（多値返却、defer）
- [x] ポインタの基本
- [ ] 構造体・メソッド
- [ ] インターフェースの基本
- [ ] package分割の基本

## Phase 2: 標準ライブラリ実践
- [ ] `errors` と wrapping（`fmt.Errorf`, `errors.Is/As`）
- [ ] `context` の使い方
- [ ] `encoding/json` のmarshal/unmarshal
- [ ] `time` とタイムゾーン
- [ ] `net/http` でHTTPクライアント/サーバ基礎
- [ ] `io` / `os` でファイル操作

## Phase 3: 並行処理
- [ ] goroutineの基礎
- [ ] channelの基礎
- [ ] selectとタイムアウト制御
- [ ] worker pool実装
- [ ] race conditionの検知（`-race`）
- [ ] 同期プリミティブ（mutex, waitgroup）

## Phase 4: テストと品質
- [ ] table-driven test
- [ ] サブテスト・ヘルパー関数
- [ ] カバレッジ計測
- [ ] benchmark基礎
- [ ] テストしやすい設計（依存分離）

## Phase 5: 実務設計
- [ ] レイヤ設計（handler/usecase/repository）
- [ ] バリデーション方針
- [ ] ログ設計（構造化ログ）
- [ ] 設定管理（環境変数・設定構造体）
- [ ] エラーハンドリング方針統一

## Phase 6: API実装
- [ ] CRUD APIの実装
- [ ] middleware（認証/ロギング/リカバリ）
- [ ] SQL接続・トランザクション
- [ ] migration運用
- [ ] APIテスト（正常系/異常系）

## Phase 7: 仕上げ
- [ ] Docker化
- [ ] ヘルスチェック/Graceful shutdown
- [ ] CI想定（lint/test/build）
- [ ] ミニプロジェクト完成・README整備

---

## 進捗ログ

### 2026-02-07
- 開始: ロードマップ作成
- 完了:
  - [ ] （ここに完了項目を追記）
- メモ:
  -

### 2026-02-11
- 完了:
  - [x] 開発環境セットアップ（Go, fmt, lint, test実行）
  - [x] 変数・定数・基本型
  - [x] 制御構文（if/switch/for）
  - [x] 関数（多値返却、defer）
  - [x] ポインタの基本
- 成果:
  - `go run ./cmd/hello Youta`
  - `go test ./...`
  - `go fmt ./...`
  - `go vet ./...`
  - `go run ./cmd/basics`
  - `go run ./cmd/control 72`
  - `go run ./cmd/functions 12 3`
  - `go run ./cmd/functions 12 0`
  - `go run ./cmd/pointer`

## 次週の目標
- [ ] `go run` / `go test` を迷わず実行できる
- [ ] 基本構文（変数、if、for、関数）で小さなCLIを1本作る
- [ ] エラーを `errors.Is` で判定するコードを1回書く
