## 概要
docker-composeを使用し、Go,MySQL,のTODOアプリの環境を構築した。

## 構築コマンド
1. クローン `git clone git@github.com:kenji-kk/go-mysql-todo.git`
2. プロジェクトフォルダに移動 `cd go-mysql-todo`
3. docker関連管理ディレクトリに移動 `cd docker`
4. .env.exampleをコピー `cp .env.example .env`
5. **COMPOSE_PROJECT_NAME=**を変更<sup>*1<sup>
6. バックグラウンドでコンテナ起動 `docker-compose up -d`
7. 停止 `docker-compose down`

<sup>*1</sup>
各コンテナ名にprefixをつけられる
nginxの場合docker-go-sample_nginxというようになる。

### 注意
一回目でエラーが起きてしまうときはもう一度`docker-compose up -d`を実行する
## 確認URL
- http://localhost:80 (nginx経由)
- http://localhost:8080
