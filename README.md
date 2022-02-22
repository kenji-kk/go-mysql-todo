## 概要
docker-composeを使用し、Go,MySQL,のTODOアプリの環境を構築した。

## 構築コマンド
1. `git clone git@github.com:kenji-kk/go-mysql-todo.git`
2. `cd go-mysql-todo`
3. `docker-compose up -d`
### 注意
一回目でエラーが起きてしまうときはもう一度`docker-compose up -d`を実行する
## 確認URL
- http://localhost:80 (nginx経由)
- http://localhost:8080
