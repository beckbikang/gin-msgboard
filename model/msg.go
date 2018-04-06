package model

import (
	"fmt"

	"gin-msgboard/storage"
	"gin-msgboard/util"
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

func NewMsgBoard(name string, msg string, mtime string) *Msgboard {
	if len(mtime) == 0 {
		mtime = util.GetCurrentTime()
	}
	msgboard := &Msgboard{
		Name:  name,
		Msg:   msg,
		Mtime: mtime,
	}
	return msgboard
}

func (msgboard *Msgboard) Add() error {
	db := storage.GetWriteDB()
	db.Table(MSGBOARD_TAB)
	return db.Save(msgboard).Error
}

func (msgboard *Msgboard) Del(id int) error {
	db := storage.GetWriteDB()
	return db.Where("id = ?", id).Delete(&Msgboard{}).Error
}

func (Msgboard) TableName() string {
	return MSGBOARD_TAB
}

func (msgboard *Msgboard) List(start int, pagesize int) []*Msgboard {
	db := storage.GetReadDB()
	list := make([]*Msgboard, 0, pagesize)
	rawSql := fmt.Sprintf("SELECT * FROM %s ORDER BY mtime DESC LIMIT %d,%d", MSGBOARD_TAB, start, pagesize)
	err := db.Raw(rawSql).Scan(&list).Error
	if err == nil {
		return list
	}
	return nil
}
