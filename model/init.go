package model

/**
 * user: ZY
 * Date: 2019/9/26 9:34
 */
import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB
var RD redis.Conn

//数据库连接且初始化
func ModelInit() {
	MysqlInit()
	RedisInit()
}

func ModelClose() {
	var err error
	err = DB.Close()
	err = RD.Close()
	if err != nil {
		log.Println("the database close failed", err)
		return
	}
}

func MysqlInit() {
	var err error

	DB, err = gorm.Open("mysql", "root:@(localhost:3306)/barrage?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("mysql init failed:", err)
		return
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func RedisInit() {
	var err error

	RD, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal("redis init failed", err)
		return
	}
}
