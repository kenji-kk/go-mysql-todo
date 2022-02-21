package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"github.com/google/uuid"
)

var Db *sql.DB

var err error

const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

func init() {
	// [ユーザ名]:[パスワード]@tcp([ホスト名]:[ポート番号])/[データベース名]?charset=[文字コード]
	Db, err = sql.Open("mysql","go_test:password@tcp(db:3306)/go_database?charset=utf8mb4")
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(255) NOT NULL UNIQUE,
		name VARCHAR(255),
		email VARCHAR(255),
		password VARCHAR(255),
		created_at TIMESTAMP)`, tableNameUser)

	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT PRIMARY KEY AUTO_INCREMENT,
		content TEXT,
		user_id INT,
		created_at TIMESTAMP)`, tableNameTodo)

	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(255) NOT NULL UNIQUE,
		email VARCHAR(255),
		user_id INT,
		created_at TIMESTAMP)`, tableNameSession)

	Db.Exec(cmdS)

	fmt.Println("データ初期化完了")
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
