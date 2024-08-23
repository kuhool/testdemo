//// server.go
//package main
//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/websocket"
//	"log"
//	"net/http"
//)
//
//var upgrader = websocket.Upgrader{
//	CheckOrigin: func(r *http.Request) bool {
//		return true // 简化示例，实际项目中应进行判断
//	},
//}
//
//// 用于管理所有连接
//var connections = make(map[*websocket.Conn]struct{})
//
//func wsHandler(c *gin.Context) {
//	// 升级连接为 WebSocket
//	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		log.Println("upgrade:", err)
//		return
//	}
//	defer ws.Close()
//
//	// 将连接加入管理
//	connections[ws] = struct{}{}
//
//	// 处理消息
//	for {
//		_, msg, err := ws.ReadMessage()
//		if err != nil {
//			log.Println("read:", err)
//			break
//		}
//		log.Println("write msg:", msg)
//		// 如果收到 Ping，回复 Pong
//		if string(msg) == "ping" {
//			if err := ws.WriteMessage(websocket.PongMessage, msg); err != nil {
//				log.Println("write pong:", err)
//				break
//			}
//
//			log.Println("write msg:", msg)
//		}
//	}
//
//	// 断开连接
//	delete(connections, ws)
//}
//
//func main() {
//	router := gin.Default()
//	router.GET("/ws", wsHandler)
//
//	log.Fatal(router.Run(":8080"))
//}
//// server.go
//package main
//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/websocket"
//	"log"
//	"net/http"
//)
//
//var upgrader = websocket.Upgrader{
//	CheckOrigin: func(r *http.Request) bool {
//		return true // 简化示例，实际项目中应进行判断
//	},
//}
//
//// 用于管理所有连接
//var connections = make(map[*websocket.Conn]struct{})
//
//func wsHandler(c *gin.Context) {
//	// 升级连接为 WebSocket
//	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		log.Println("upgrade:", err)
//		return
//	}
//	defer ws.Close()
//
//	// 将连接加入管理
//	connections[ws] = struct{}{}
//
//	// 处理消息
//	for {
//		_, msg, err := ws.ReadMessage()
//		if err != nil {
//			log.Println("read:", err)
//			break
//		}
//		log.Println("write msg:", msg)
//		// 如果收到 Ping，回复 Pong
//		if string(msg) == "ping" {
//			if err := ws.WriteMessage(websocket.PongMessage, msg); err != nil {
//				log.Println("write pong:", err)
//				break
//			}
//
//			log.Println("write msg:", msg)
//		}
//	}
//
//	// 断开连接
//	delete(connections, ws)
//}
//
//func main() {
//	router := gin.Default()
//	router.GET("/ws", wsHandler)
//
//	log.Fatal(router.Run(":8080"))
//}
//// server.go
//package main
//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/websocket"
//	"log"
//	"net/http"
//)
//
//var upgrader = websocket.Upgrader{
//	CheckOrigin: func(r *http.Request) bool {
//		return true // 简化示例，实际项目中应进行判断
//	},
//}
//
//// 用于管理所有连接
//var connections = make(map[*websocket.Conn]struct{})
//
//func wsHandler(c *gin.Context) {
//	// 升级连接为 WebSocket
//	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		log.Println("upgrade:", err)
//		return
//	}
//	defer ws.Close()
//
//	// 将连接加入管理
//	connections[ws] = struct{}{}
//
//	// 处理消息
//	for {
//		_, msg, err := ws.ReadMessage()
//		if err != nil {
//			log.Println("read:", err)
//			break
//		}
//		log.Println("write msg:", msg)
//		// 如果收到 Ping，回复 Pong
//		if string(msg) == "ping" {
//			if err := ws.WriteMessage(websocket.PongMessage, msg); err != nil {
//				log.Println("write pong:", err)
//				break
//			}
//
//			log.Println("write msg:", msg)
//		}
//	}
//
//	// 断开连接
//	delete(connections, ws)
//}
//
//func main() {
//	router := gin.Default()
//	router.GET("/ws", wsHandler)
//
//	log.Fatal(router.Run(":8080"))
//}
