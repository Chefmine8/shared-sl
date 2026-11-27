package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	clients      []*websocket.Conn
	clientsMutex sync.Mutex
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgraded error : ", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected !")

	messageType, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("Client disconnected or read message error :", err)
		return
	}

	fmt.Printf("Received msg : %s\n", message)

	reponse := fmt.Sprintf("Your are ")
	err = conn.WriteMessage(messageType, []byte(reponse))
	if err != nil {
		log.Println("Send error : ", err)
		return
	}
	clientsMutex.Lock()
	clients = append(clients, conn)
	clientsMutex.Unlock()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected or read message error :", err)
			return
		}
	}
}

func remove(i int) {
	clientsMutex.Lock()
	clients = append(clients[:i], clients[i+1:]...)
	clientsMutex.Unlock()
}

func back() {
	for {
		clientsMutex.Lock()
		if len(clients) == 0 {
			clientsMutex.Unlock()
			time.Sleep(10 * time.Second)
			continue
		}
		for i := 0; i < len(clients); i++ {
			conn := clients[i]
			clientsMutex.Unlock()
			send := fmt.Sprintf("Your turn")
			err := conn.WriteMessage(1, []byte(send))
			if err != nil {
				log.Println("Send error : ", err)
				remove(i)
				i--
				clientsMutex.Lock()
				continue
			}
			time.Sleep(1 * time.Second)
			clientsMutex.Lock()
		}
		clientsMutex.Unlock()
	}
}

func StartWebSocket(port int) {
	go back()
	http.HandleFunc("/ws", handleWebSocket)
	p := fmt.Sprintf(":%d", port)
	fmt.Println("Serveur start at http://localhost:" + p)
	log.Fatal(http.ListenAndServe(p, nil))
}
