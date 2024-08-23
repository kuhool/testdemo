package test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestOne(t *testing.T) {

	url := "ws://localhost:8080/ws"
	log.Println("Connecting to", url)

	// 创建一个新的WebSocket连接
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	fmt.Println("Connected")
	// 发送消息到服务器
	for i := 0; ; i++ {
		conn.WriteMessage(websocket.TextMessage, []byte("Hello "+time.Now().String()))

		// 接收服务器回传的消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		time.Sleep(1 * time.Second)
	}
}

func TestTwo(t *testing.T) {

	url := "ws://localhost:8080/ws"
	log.Println("Connecting to", url)

	// 创建一个新的WebSocket连接
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	fmt.Println("Connected")
	//// 发送消息到服务器
	//for i := 0; ; i++ {
	//	//conn.WriteMessage(websocket.TextMessage, []byte("Hello "+time.Now().String()))
	//	err = conn.WriteMessage(websocket.PingMessage, nil)
	//	fmt.Println(err)
	//	// 接收服务器回传的消息
	//	_, msg, err := conn.ReadMessage()
	//	if err != nil {
	//		log.Println("read:", err)
	//		break
	//	}
	//	log.Printf("recv: %s", msg)
	//
	//	time.Sleep(1 * time.Second)
	//}

	// 发送 Ping 消息并接收 Pong 响应
	pingTicker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-pingTicker.C:
			fmt.Println("Sending ping...")
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
