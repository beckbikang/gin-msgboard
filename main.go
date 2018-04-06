package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"gin-msgboard/config"
	"gin-msgboard/controller"
	"gin-msgboard/storage"
	"github.com/gin-gonic/gin"
)

const (
	ErrnoOk int = iota
	ErrorLoadTemplates
)

func init() {
	config.LoadConfig("./conf/config.json")
}

func main() {
	log.Println("start running ...")
	router := gin.Default()

	tmpls, err := filepath.Glob("view/*/*.tmpl")
	if err != nil {
		log.Printf("failed to load templates(1)")
		os.Exit(ErrorLoadTemplates)
	}

	//加载html文件
	router.LoadHTMLFiles(tmpls...)

	//数据库初始化

	router.GET("/hi", sayHi)
	//添加数据
	msgController := new(controller.Msglist)
	log.Println(msgController)
	//router.GET("/add", new(controller.Msglist).Add)
	router.GET("/msg/add", msgController.Add)

	router.POST("/msg/add", msgController.AddDone)

	router.GET("/msg/list", msgController.List)

	router.GET("/msg/del", msgController.Del)

	//加载数据库
	storage.InitDB()

	//添加数据的post

	//列表

	router.Run(config.GetServer().Host + ":" +
		strconv.Itoa(config.GetServer().Port))
}

func sayHi(c *gin.Context) {
	c.String(200, "hello")
}
