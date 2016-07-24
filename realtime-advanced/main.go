package main

import (
	"fmt"
	"runtime"
	"realtime-advanced/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()//返回执行当前进程的CPU数目
	runtime.GOMAXPROCS(nuCPU)//系统最大CPU数目
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartWorkers() {
	//持续提取CPU的情况
	go lib.StatsWorker()
}

func StartGin() {
	gin.SetMode(gin.ReleaseMode)  //额，分出了三个不同的版本，不然就是panic的，虽然不知道有什么用

	router := gin.Default() //注册路由
	router.Use(lib.RateLimit, gin.Recovery())//这里的Use主要用于调用中间件，比如这里就是为了调用中间件来限制rate
	router.LoadHTMLGlob("resources/room_login.templ.html")
	//router.Static("/static", "resources/static")
	router.GET("/", lib.Index)
	router.GET("/room/:roomid", lib.RoomGET)
	router.POST("/room-post/:roomid", lib.RoomPOST)
	router.GET("/stream/:roomid", lib.StreamRoom)

	router.Run(":80")
}
