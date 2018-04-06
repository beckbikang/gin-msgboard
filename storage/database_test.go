package storage

import (
	"fmt"
	"testing"

	"gin-msgboard/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Msgboard struct {
	Id    uint `gorm:"primary_key"`
	Name  string
	Msg   string
	Mtime string
}

const (
	MSGBOARD_TAB = "msgboard"
)

func TestInitDB(t *testing.T) {

	configer, _ := config.LoadConfigFile("../conf/config.json")
	config.GlobalDefaultConfig = *configer

	conf := config.GetDatabase()

	server := conf.DbServers[0]
	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		server.User,
		server.Password,
		server.Host,
		server.Port,
		server.Schema,
	)
	t.Log(args)

	db, _ = gorm.Open(conf.Type, args)
	for i := 0; i < 20; i++ {
		msg := Msgboard{Name: "Jinzhu", Msg: "tt"}
		db = db.Table(MSGBOARD_TAB).Save(&msg)

		t.Log(db.Error)
	}
}
