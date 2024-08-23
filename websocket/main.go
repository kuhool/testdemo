package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		handleWebsocket(c.Writer, c.Request)
	})

	r.Run(":8080")
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	//连接成功
	log.Println("连接成功", conn)
	//for {
	//	log.Println("start...............1")
	//	messageType, message, err := conn.ReadMessage()
	//	if err != nil {
	//		return
	//	}
	//	fmt.Println("start...............2")
	//	fmt.Println("messageType", messageType)
	//	if messageType == websocket.PingMessage {
	//		fmt.Println("ping....... pong")
	//		_ = conn.WriteMessage(websocket.PongMessage, nil)
	//	}
	//	fmt.Println("start...............3")
	//	// 如果收到 Ping 消息，回复 Pong 消息
	//	if string(message) == "ping" {
	//		conn.WriteMessage(websocket.TextMessage, []byte("pong"))
	//	}
	//
	//	fmt.Println(messageType)
	//	if err := conn.WriteMessage(+messageType, []byte("server say "+string(message))); err != nil {
	//		return
	//	}
	//}
}
