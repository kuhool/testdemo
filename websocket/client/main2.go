//// client.go
//package main
//
//import (
//	"github.com/gorilla/websocket"
//	"log"
//	"time"
//)
//
//func main() {
//	// 设置服务端地址
//	address := "ws://localhost:8080/ws"
//
//	// 创建一个新的 WebSocket 连接
//	conn, _, err := websocket.DefaultDialer.Dial(address, nil)
//	if err != nil {
//		log.Fatal("dial:", err)
//	}
//	defer conn.Close()
//	log.Println("111111111111111")
//	// 发送 Ping 消息并接收 Pong 响应
//	pingTicker := time.NewTicker(2 * time.Second)
//	defer pingTicker.Stop()
//	log.Println("2222222222222")
//	for {
//		select {
//		case <-pingTicker.C:
//			log.Println("33333333333333333333")
//			// 发送 Ping
//			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
//				log.Println("write ping:", err)
//				return
//			}
//			//log.Println("44444444444444444444444")
//			//// 等待 Pong
//			//messageType, _, err := conn.ReadMessage()
//			//if err != nil {
//			//	log.Println("read message:", err)
//			//	return
//			//}
//			//log.Println("55555555555555")
//			//if messageType != websocket.PongMessage {
//			//	log.Printf("expected pong message, got %v", messageType)
//			//	return
//			//}
//			//log.Println("Received pong from server")
//		}
//	}
//}
