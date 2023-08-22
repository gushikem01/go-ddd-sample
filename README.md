## go-handson

golang練習用のリポジトリになります。
このリポジトリではgolangの基本的なDDDが学べます。

handler (APIのリクエストを受ける)
↓
usecase (handlerから受けたリクエストをmodelに変換)
↓
repository (datasourceの関数指定)
↓
datasource (modelの永続化)

## 流れ

1. postgresqlとgolangのコンテナをdockerで起動します。
2. curlコマンドを叩いて、ユーザーの作成を行います。
3. postgresqlのコンテナに入ってSQLを実行してテーブルにユーザーが追加されたかどうか確認する。

## 動作確認

リポジトリをクローンします。

```shell
gh repo clone github.com/gushikem01/go-handson
```

環境変数をコピーします。

```code
cp .env.sample .env
```

postgresqlとgolangのコンテナをコマンド一発で起動します。

```code
docker-compose up -d
```

curlコマンドを実行してユーザー作成APIを叩きます。

```code
curl -X POST -H "Content-Type: application/json" -d '{"name":"テスト0", "email":"test@test.com", "password":"123456"}' localhost:8080/api/v1/users
```

データが追加されたことを確認します。

```sql
psql "postgresql://user:password@localhost:54321/user_app?sslmode=disable"
```

pgcliがインストールされていない場合、DBEaverにて下記設定を入れてDBに接続します。

ユーザー名: user
パスワード： password
DB： user_app
ポート: 54321
ホスト名: localhost

```sql
select * from users;
```
