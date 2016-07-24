package lib

import "github.com/dustin/go-broadcast"  //broadcast 是一个组播，用于

var roomChannels = make(map[string]broadcast.Broadcaster)

func OpenListener(roomid string) chan interface{} {
	//开启一个组播链
	listener := make(chan interface{})
	Room(roomid).Register(listener)
	return listener
}

func CloseListener(roomid string, listener chan interface{}) {
	//关闭组播链条
	Room(roomid).Unregister(listener)
	close(listener)
}

func DeleteBroadcast(roomid string) {
	//删除组播链条
	b, ok := roomChannels[roomid]
	if ok {
		b.Close()
		delete(roomChannels, roomid)
	}
}

func Room(roomid string) broadcast.Broadcaster {
	//创建10个组播的位置
	b, ok := roomChannels[roomid]
	if !ok {
		b = broadcast.NewBroadcaster(10)// new c出10个组播的broadcast
		roomChannels[roomid] = b
	}
	return b
}
