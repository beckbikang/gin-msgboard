package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"gin-msgboard/model"
	"gin-msgboard/util"
	"github.com/gin-gonic/gin"
)

type Msglist struct{}

func (msg *Msglist) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "add.tmpl", gin.H{
		"title": "留言",
	})
}

func (msg *Msglist) AddDone(c *gin.Context) {

	name := strings.TrimSpace(c.PostForm("name"))
	if name == "" || len(name) == 0 {
		jumpToMsgAdd(c)
	}

	msgStr := strings.TrimSpace(c.PostForm("msg"))
	if msgStr == "" || len(msgStr) == 0 {
		jumpToMsgAdd(c)
	}

	msgboard := model.NewMsgBoard(name, msgStr, util.GetCurrentTime())

	//添加数据
	err := msgboard.Add()
	if err != nil {
		log.Printf("%v", err)
	}
	jumpToMsgList(c)
}

//列表
func (msg *Msglist) List(c *gin.Context) {

	//页码
	page, err := strconv.Atoi(c.Query("page"))
	if page <= 0 || err != nil {
		page = 1
	}

	//每页条数
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if pagesize <= 0 || err != nil {
		pagesize = 20
	}

	start := (page - 1) * pagesize

	msgboard := new(model.Msgboard)
	msglist := msgboard.List(start, pagesize)

	title := "留言列表"

	prePage := page - 1

	lastPage := page + 1
	if len(msglist) < pagesize {
		lastPage = 0
	}

	showData := make(map[string]interface{})

	showData["title"] = title
	showData["list"] = msglist
	showData["prePage"] = prePage
	showData["lastPage"] = lastPage
	c.HTML(http.StatusOK, "list.tmpl", showData)
}

func (msg *Msglist) Del(c *gin.Context) {
	//页码
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		jumpToMsgAdd(c)
	}
	err = (&model.Msgboard{}).Del(id)
	jumpToMsgList(c)
}

func jumpToMsgAdd(c *gin.Context) {
	c.Redirect(http.StatusFound, "/msg/add")
}
func jumpToMsgList(c *gin.Context) {
	c.Redirect(http.StatusFound, "/msg/list")
}

func jumpUrl(c *gin.Context, url string) {
	c.Redirect(http.StatusFound, url)
}
