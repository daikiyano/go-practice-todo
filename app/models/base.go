package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"go-todo-practice/config"
	"log"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						uuid STRING NOT NULL UNIQUE,
						name STRING,email STRING,
						password STRING, 
						created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)
}

// UUID作成
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// パスワードを暗号化
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
