package db

import (
	"blog_api/log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var once sync.Once
var mu sync.Mutex
var db *sql.DB
var connectionString string

func init() {
	dbSource := "root:yakun,0622@tcp(106.12.85.26:3306)/cc_bitfinex_top10_hd_15m_2019_07_12?parseTime=true"
	initMysql(dbSource)
}

func initMysql(connStr string) {
	if db == nil || connStr != connectionString {
		connectionString = connStr
		once.Do(openDB)
	}
}

func openDB() {
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		errMsg := fmt.Sprintf("数据库连接失败1: %v", err)
		panic(errMsg)
	}
	db.SetConnMaxLifetime(60)
	if err = db.Ping(); err != nil {
		errMsg := fmt.Sprintf("数据库连接失败2: %v", err)
		panic(errMsg)
	}
}

func Close() {
	mu.Lock()
	defer mu.Unlock()
	log.Logger.Info("default db closing")
	err := db.Close()
	if err != nil {
		log.Logger.Info("default db close exception")
	} else {
		log.Logger.Info("default db close success")
	}

}

func getDb() *sql.DB {
	if db == nil {
		openDB()
	}
	return db
}

func FetchRows(sql string, args ...interface{}) (*sql.Rows, error) {
	if db == nil {
		return nil, DB_NIL
	}
	return db.Query(sql, args...)
}
