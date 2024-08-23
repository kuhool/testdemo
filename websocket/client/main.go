// client.go
package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	// 连接到 WebSocket 服务端
	conn, _, err := websocket.DefaultDialer.Dial("ws://"+*addr+"/ws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// 发送 Ping 消息并接收 Pong 响应
	pingTicker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-pingTicker.C:
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("write ping:", err)
				return
			}
			messageType, _, err := conn.ReadMessage()
			if err != nil {
				log.Println("read pong:", err)
				return
			}
			if messageType != websocket.PongMessage {
				log.Println("expected pong message, got:", messageType)
				return
			}
			log.Println("Received pong from server")
		}
	}
}
