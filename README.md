## go-book-simple-api
* Go言語で作成したREST API (CRUD) サンプルです。
* フレームワーク
  * Gin
* ORM
  * GORM
* DB
  * MySQL

### テーブル構成

##### booksテーブル
| Id | Name | Price |
| -- | -- | -- |
| integer | string | integer |

### API定義

| URL | Method | Description |
| -- | -- | -- |
| /books | GET | book情報取得 |
| /books | POST | book情報作成 |
| /books/:id | PUT | book情報更新 |
| /books/:id | DELETE | book情報削除 |

### 設計

* <a href="/app">app</a>配下にGoアプリケーション、<a href="/build">build</a>配下にDockerfileを格納しています。

```
.
├── README.md
├── /app
│   ├── /controller
│   ├── /model
│   ├── /repository
│   ├── /server
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── /build
│   └── /app
│       └── Dockerfile
└── docker-compose.yml
```

* アプリケーションの依存関係は下記のようになっています。<br><img src="https://github.com/plasmo310/go-book-simple-api/assets/77447256/8b960c9a-8114-4214-a126-f4fc925326a8" width=200 />

### 使用方法

* <a href="env_rename_me">.env_rename_me</a>ファイルをコピーして.envファイルを作成します。
* <code>docker compose up -d</code>を実行して localhost:8080 で各APIを実行します。
  * もしくは、<a href="/build/app/Dockerfile">Dockerfile</a>の<code>CMD ["go", "run", "main.go"]</code>をコメントアウトして、下記コマンドで直接実行します。
     ```
     docker compose exec app go run main.go
     ```
