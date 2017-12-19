package db

import (
	"github.com/ihahoo/go-api-lib/config"
	"github.com/ihahoo/go-api-lib/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Client 默认数据库实例
var Client *sqlx.DB

// Init 默认初始化
func Init() {
	Client = Conn()
}

// ConnectDB 用连接字符串连接数据库
func ConnectDB(dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Connect 用配置文件连接数据库
func Connect(prePath string) *sqlx.DB {
	user := config.GetString(prePath + "user")
	password := config.GetString(prePath + "password")
	host := config.GetString(prePath + "host")
	port := config.GetString(prePath + "port")
	dbname := config.GetString(prePath + "dbname")
	connectTimeout := config.GetString(prePath + "connectTimeout")

	db, err := ConnectDB("dbname=" + dbname + " user=" + user + " password=" + password + " host=" + host + " port=" + port + " connect_timeout=" + connectTimeout + " sslmode=disable")
	if err != nil {
		log.GetLog().WithFields(log.Fields{"func": "db.Connect"}).Fatal(err)
	}
	return db
}

// Conn 用配置文件的默认参数连接数据库
func Conn() *sqlx.DB {
	return Connect("db.options.")
}
