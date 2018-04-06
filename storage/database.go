package storage

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"gin-msgboard/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var Inited bool //是否初始化

func InitDB() error {
	var err error

	conf := config.GetDatabase()

	fmt.Println(len(conf.DbServers))

	if len(conf.DbServers) < 1 {
		return fmt.Errorf("InitDB: NO configuration for database server")
	}
	server := conf.DbServers[0]
	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		server.User,
		server.Password,
		server.Host,
		server.Port,
		server.Schema,
	)
	//打开orm
	db, err = gorm.Open(conf.Type, args)
	if err != nil {
		return err
	}

	Inited = true
	return nil
}

func CloseDB() error {
	return db.Close()
}

//获取db
func GetReadDB() *gorm.DB {
	return db
}

//获取db
func GetWriteDB() *gorm.DB {
	return db
}
