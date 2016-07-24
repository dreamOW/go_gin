package main

import (
	"fmt"
	"io"
	"math/rand"
	"realtime-chat/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetHTMLTemplate(lib.Html)  //显示模板界面

	router.GET("/room/:roomid", roomGET) //根据get到的roomid决定开启的roomid
	router.POST("/room/:roomid", roomPOST) //把信息获取并发送出去
	router.DELETE("/room/:roomid", roomDELETE) //删除那个组播点
	router.GET("/stream/:roomid", stream)

	router.Run(":8080") 
}

func stream(c *gin.Context) {
	roomid := c.Param("roomid")
	fmt.Printf("roomid is:  ",roomid)
	listener := lib.OpenListener(roomid)
	defer lib.CloseListener(roomid, listener)

	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", <-listener)
		return true
	})
}

func roomGET(c *gin.Context) {
	roomid := c.Param("roomid")  //获得客户端传来的roomid 如localhost：80/room/：1 roomid就等于1
	userid := fmt.Sprint(rand.Int31()) //随机获得一个userid
	c.HTML(200, "chat_room", gin.H{      //type H map[string]interface{}
		"roomid": roomid,
		"userid": userid,
	})
}

func roomPOST(c *gin.Context) {
	roomid := c.Param("roomid") 
	userid := c.PostForm("user")
	message := c.PostForm("message")
	//从表单中获取用户ID和消息信息
	lib.Room(roomid).Submit(userid + ": " + message)
	//以roomid为键值来创建一个组播，并把信息发送出去
	c.JSON(200, gin.H{
		"status":  "success",
		"message": message,
	})
}

func roomDELETE(c *gin.Context) {
	roomid := c.Param("roomid")
	lib.DeleteBroadcast(roomid)
}
