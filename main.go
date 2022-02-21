package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
    dbconf := "go_test:password@tcp(db:3306)/go_database?charset=utf8mb4"

    db, err := sql.Open("mysql", dbconf)

    // 接続が終了したらクローズする
    defer db.Close()  

    if err != nil {
        fmt.Println(err.Error())
    }

    err = db.Ping()

    if err != nil {
        fmt.Println("データベース接続失敗")
        return
    } else {
        fmt.Println("データベース接続成功")
    }
}
