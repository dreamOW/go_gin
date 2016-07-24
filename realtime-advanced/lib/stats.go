package lib

import (
	"runtime"
	"sync"
	"time"

	"github.com/manucorporat/stats"
)

//STATS这个包的基本作用就是封装了一个锁机制，在增删改查的过程中用锁机制来操作异步等情况
var ips = stats.New()
var messages = stats.New()
var users = stats.New()
var mutexStats sync.RWMutex
var savedStats map[string]uint64

func StatsWorker() {
	c := time.Tick(1 * time.Second) //Tick可以重复的每隔一段时间向C（c是一个channel）发送时间
	var lastMallocs uint64
	var lastFrees uint64
	for range c {
		var stats runtime.MemStats
		runtime.ReadMemStats(&stats)

		mutexStats.Lock()
		savedStats = map[string]uint64{
			"timestamp":  uint64(time.Now().Unix()),
			"HeapInuse":  stats.HeapInuse,
			"StackInuse": stats.StackInuse,
			"Mallocs":    (stats.Mallocs - lastMallocs),
			"Frees":      (stats.Frees - lastFrees),
			"Inbound":    uint64(messages.Get("inbound")),
			"Outbound":   uint64(messages.Get("outbound")),
			"Connected":  connectedUsers(),
		}
		lastMallocs = stats.Mallocs
		lastFrees = stats.Frees
		messages.Reset()
		mutexStats.Unlock()
	}
	//本段代码相当于每隔一段时间提取一次CPU的情况
}

func connectedUsers() uint64 {
	connected := users.Get("connected") - users.Get("disconnected")
	if connected < 0 {
		return 0
	}
	return uint64(connected)
}

func Stats() map[string]uint64 {
	mutexStats.RLock()
	defer mutexStats.RUnlock()

	return savedStats
}
