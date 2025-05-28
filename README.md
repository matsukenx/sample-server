# sample-server

## 概要

Go 言語製の API サーバーです。Echo フレームワークを利用しています。
go-server-template の雛形を作成

## 前提条件

- Go 1.22.5 以上
- (任意) .envrc ファイルで環境変数を設定可能

## セットアップ

1. リポジトリのクローン

```sh
git clone <リポジトリURL>
cd sample-server
```

2. 依存パッケージのインストール

```sh
go mod tidy
```

3. 環境変数の設定

`.envrc` ファイルを作成し、必要に応じて `PORT` や `API_KEY` などを設定してください。

例:

```
PORT=8080
API_KEY=test
```

4. サーバーの起動

```sh
go run cmd/sample-server/main.go
```

## API の利用例

### Hello エンドポイント

```sh
curl -H "X-API-KEY: test" "http://localhost:8080/hello?name=Ken"
```

### ユーザー取得エンドポイント

```sh
curl -H "X-API-KEY: test" "http://localhost:8080/users/1"
```

## 注意事項

- ポートが既に使用中の場合は、他のプロセスを停止するか、`.env` で別のポートを指定してください。
- API_KEY は適宜変更してください。
