package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"github.com/google/uuid"

	"todo/config"
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
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbUser + ":" + config.Config.DbPassword + config.Config.DbPort + "/" + config.Config.DbName + "?" + "charset=" + config.Config.DbCharSet )
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

		_, err = Db.Exec(cmdT)
		if err != nil {
			log.Fatalln(err)
		}

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(255) NOT NULL UNIQUE,
		email VARCHAR(255),
		user_id INT,
		created_at TIMESTAMP)`, tableNameSession)

		_, err = Db.Exec(cmdS)
		if err != nil {
			log.Fatalln(err)
		}

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
