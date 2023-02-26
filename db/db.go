package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB_PATH string
var sdb *sql.DB = nil

func Init(path string) {
	DB_PATH = path
	db, err := sql.Open("sqlite3", DB_PATH)
	if err != nil {
		log.Fatalln("Error:", err)
	}
    sdb = db

	sqls := []string{
		"CREATE TABLE IF NOT EXISTS `accessTokens` (`uid` INTEGER PRIMARY KEY AUTOINCREMENT, `token` VARCHAR(64) NOT NULL UNIQUE)",
	}

	for _, s := range sqls {
		if stmt, err := sdb.Prepare(s); err != nil {
			log.Fatalln("Error:", err)
		} else {
			if _, err := stmt.Exec(); err != nil {
				log.Fatalln("Error:", err)
			}
		}
	}

	// 创建Token表
	if tokens, err := QueryAccessTokens(); err == nil {
        if len(tokens) <= 0 {
            t := "testToken"
		    tokens = append(tokens, t )
            AddAccessToken(t)
        }

		for _, token := range tokens {
			s := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (`uid` INTEGER PRIMARY KEY AUTOINCREMENT, `name` VARCHAR(256) NOT NULL UNIQUE, `value` VARCHAR(65535) NOT NULL)", token)
			if stmt, err := sdb.Prepare(s); err != nil {
				log.Fatalln("Error:", err)
			} else {
				if _, err := stmt.Exec(); err != nil {
					log.Fatalln("Error:", err)
				}
			}

		}
	}
}

func Uninit () {
    if sdb != nil {
        sdb.Close()
    }
}
