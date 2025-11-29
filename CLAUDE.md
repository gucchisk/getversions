# CLAUDE.md

このファイルは、このリポジトリで作業する際にClaude Code (claude.ai/code) へのガイダンスを提供します。

## 概要

`getversions`は、様々なソフトウェアプロジェクトのウェブサイトからバージョン情報を抽出するCLIツールです。ビルトインコマンドと拡張性のためのプラグインシステムの両方をサポートしています。

## ビルドコマンド

```bash
# メインアプリケーションのビルド
go build -o getversions

# プラグインのビルド（例: gradle）
go build -o getversions-gradle actions/gradle/main.go

# Task を使用する場合（インストールされている場合）
task build          # メインアプリケーションのビルド
task build-gradle   # gradleプラグインのビルド
task test           # すべてのテストを実行
task clean          # バイナリを削除
```

## テストの実行

```bash
# すべてのテストを実行
go test ./...

# 特定のパッケージのテストを実行
go test ./pkg/utils/version
go test ./pkg/latest/apache
```

## アーキテクチャ

### プラグインシステム

このアプリケーションは、RPCベースのプラグイン用にHashiCorp's go-pluginフレームワークを使用しています：

**主要コンポーネント:**
- `pkg/action/action.go` - プラグインインターフェースとRPCレイヤーを定義
- `actions/*/main.go` - プラグインの実装
- `cmd/root.go` - プラグインの検出とコマンド登録

**プラグインインターフェース (`GetVersionsAction`):**
```go
type GetVersionsAction interface {
    Version() string
    Short() string
    Long() string
    GetVersions(reader io.Reader) []string
}
```

**プラグインの重要な詳細:**
1. プラグインは`$GOPATH/bin`で`getversions-`で始まるバイナリを検索して発見される
2. プラグインコマンドは実行時に`addPluginCommands()`を介して動的に登録される
3. RPC通信: `io.Reader`はシリアライゼーションのために`[]byte`に変換する必要がある（インターフェースをRPC経由で直接渡すことはできない）
4. ハンドシェイク設定: `GETVERSIONS_PLUGIN="hello"`、プロトコルバージョン1を使用

**新しいプラグインの作成:**
1. `GetVersionsAction`インターフェースを実装した`actions/<name>/main.go`を作成
2. ビルド: `go build -o getversions-<name> actions/<name>/main.go`
3. インストール: バイナリを`$GOPATH/bin/`に配置
4. プラグインは自動的にサブコマンドとして登録される

### ビルトインコマンド

**Apacheコマンド (`cmd/apache.go`):**
- Apache AutoIndexページからバージョンを抽出
- HTMLパースに`pkg/latest/apache/apache.go`を使用
- `-v`フラグによるバージョンフィルタリングをサポート
- 最新の一致するバージョンを返す

### コアユーティリティ

**`pkg/utils/version/`** - バージョンの解析と比較:
- `SearchVersion()` - テキストからバージョン文字列を抽出
- `ToSemver()` / `FromSemver()` - Semverの正規化
- `IsBig()` - バージョン文字列の比較

**`pkg/utils/htmlparser/`** - HTMLトラバーサルヘルパー:
- `FindAll()` - atom型ですべてのノードを検索
- `FindAllTexts()` - テキストノードを抽出

**`pkg/utils/utils.go`** - プラグイン検出:
- `SearchPlugins()` - `$GOPATH/bin`で`getversions-*`バイナリをスキャン

### ロギング

アプリケーションはzapバックエンドで`go-logr`を使用：
- レベルは`--log`フラグで制御（0=warn, 1=info, 2=debug）
- プラグインはgo-pluginとの互換性のために`hclog`を使用

## 一般的なパターン

**新しいデータソースの追加:**
1. シンプルなスクレイパーの場合: `actions/<name>/`に新しいプラグインを作成
2. 複雑なロジックの場合: `cmd/<name>.go`にビルトインコマンドを追加し、`pkg/latest/<name>/`にパッケージを追加

**RPCシリアライゼーションルール:**
プラグインRPC経由でデータを渡す際は、クライアント側で`io.ReadAll()`を使用して`io.Reader`を`[]byte`に変換し、サーバー側で`bytes.NewReader()`を使用して再構築します。Goのgobエンコーディングは、エクスポートされていないフィールドを持つインターフェース型をシリアライズできません。
