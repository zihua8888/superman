package main
//
//import (
//	"fmt"
//	"github.com/gorilla/websocket"
//	"log"
//	"net/http"
//)
//
////const (
////	// Time allowed to write a message to the peer.时间允许写一条消息给同伴
////	writeWait = 10 * time.Second
////	// Time allowed to read the next pong message from the peer.允许时间读取来自对等方的下一条pong消息
////	pongWait = 60 * time.Second
////	// Send pings to peer with this period. Must be less than pongWait.：发送pings到peer。一定比彭维特小
////	pingPeriod = (pongWait * 9) / 10
////	// Maximum message size allowed from peer.：对等点允许的最大消息大小。
////	maxMessageSize = 512
////)
//
//type message struct {
//	data []byte
//	room string
//}
//
//type subscription struct {
//	conn *connection
//	room string
//}
//
//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//}
//
//type connection struct {
//	// The websocket connection.   websocket连接。
//	ws *websocket.Conn
//	// Buffered channel of outbound messages.   ：出站消息的缓冲通道。
//	//send chan []byte
//}
//
//// hub maintains the set of active connections and broadcasts messages to the  ：hub维护一组活动连接，并将消息广播到
//// connections. //连接
//type hub struct {
//	// Registered connections. //注册连接。
//	rooms map[string]map[*connection]bool
//
//	// Inbound messages from the connections.  来自连接的入站消息。
//	broadcast chan message
//
//	// Register requests from the connections.  注册来自连接的请求。
//	register chan subscription
//
//	// Unregister requests from connections.  注销来自连接的请求。
//	unregister chan subscription
//}
//
//var h = hub{
//	broadcast:  make(chan message),
//	register:   make(chan subscription),
//	unregister: make(chan subscription),
//	rooms:      make(map[string]map[*connection]bool),
//}
//
//
//func (h *hub) run() {
//	for {
//		select {
//		case s := <-h.register:
//			connections := h.rooms[s.room]
//			if connections == nil {
//				connections = make(map[*connection]bool)
//				h.rooms[s.room] = connections
//			}
//			h.rooms[s.room][s.conn] = true
//
//		case s := <-h.unregister:
//			connections := h.rooms[s.room]
//			if connections != nil {
//				if _, ok := connections[s.conn]; ok {
//					delete(connections, s.conn)
//					//close(s.conn.send)
//					if len(connections) == 0 {
//						delete(h.rooms, s.room)
//					}
//				}
//			}
//
//
//		//case m := <-h.broadcast:
//		//	connections := h.rooms[m.room]
//		//	for c := range connections {
//		//		select {
//		//		case c.send <- m.data:
//		//		default:
//		//			close(c.send)
//		//			delete(connections, c)
//		//			if len(connections) == 0 {
//		//				delete(h.rooms, m.room)
//		//			}
//		//		}
//		//	}
//
//		}
//	}
//}
//
////func (c *connection) write(mt int, payload []byte) error {
////	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
////	return c.ws.WriteMessage(mt, payload)
////}
//
//
////func (s subscription) readPump() {
////	c := s.conn
////	defer func() {
////		h.unregister <- s
////		c.ws.Close()
////	}()
////	c.ws.SetReadLimit(maxMessageSize)
////	c.ws.SetReadDeadline(time.Now().Add(pongWait))
////	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
////	for {
////		_, msg, err := c.ws.ReadMessage()
////		if err != nil {
////			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
////				log.Printf("error: %v", err)
////			}
////			break
////		}
////		m := message{msg, s.room}
////		h.broadcast <- m
////	}
////}
//
//
////func (s *subscription) writePump() {
////	c := s.conn
////	ticker := time.NewTicker(pingPeriod)
////	defer func() {
////		ticker.Stop()
////		c.ws.Close()
////	}()
////	for {
////		select {
////		case message, ok := <-c.send:
////			if !ok {
////				c.write(websocket.CloseMessage, []byte{})
////				return
////			}
////			if err := c.write(websocket.TextMessage, message); err != nil {
////				return
////			}
////		case <-ticker.C:
////			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
////				return
////			}
////		}
////	}
////}
//
//func main() {
//	fmt.Println("Starting application...")
//	//开一个goroutine执行开始程序
//	go h.run()
//	//注册默认路由为 /ws ，并使用wsHandler这个方法
//	http.HandleFunc("/ws", serveWs)
//	//监听本地的8011端口
//	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
//}
//
//func serveWs(w http.ResponseWriter, r *http.Request) {
//
//	//upgrader.CheckOrigin = func(r *http.Request) bool {
//	//	return true
//	//}
//	//ws, err := upgrader.Upgrade(w, r, nil)
//	////vars := mux.Vars(r)
//	////log.Println(vars["room"])
//	//if err != nil {
//	//	log.Println(err)
//	//	return
//	//}
//	//r.ParseForm()
//	//roomid := r.Form["roomid"][0]
//
//
//	c := &connection{ ws: ws}
//	s := subscription{c,roomid }
//	h.register <- s
//	//go s.writePump()
//	//go s.readPump()
//
//}