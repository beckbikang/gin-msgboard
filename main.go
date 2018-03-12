package main

import (
	"github.com/gin-msgboard/config"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadConfig("./conf/config.json")
}

func main() {
	/*
		log.Println("start-run")
		runHi()
	*/
	log.Println(GetIsInit())

}

func runHi() {
	router := gin.Default()
	router.GET("/hi", sayHi)
	router.Run(":8080")
}

func sayHi(c *gin.Context) {
	c.String(200, "hello")
}
