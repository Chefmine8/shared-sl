package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"

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
}

func remove(i int) {
	clientsMutex.Lock()
	clients = append(clients[:i], clients[i+1:]...)
	clientsMutex.Unlock()
}

func back() {
	for {
		clientsMutex.Lock()
		for i := 0; i < len(clients); i++ {
			conn := clients[i]
			clientsMutex.Unlock()
			log.Println("Is", i, " turn.")
			send := fmt.Sprintf("Your turn")
			err := conn.WriteMessage(1, []byte(send))
			if err != nil {
				log.Println("Send error : ", err)
				remove(i)
				conn.Close()
				i--
				clientsMutex.Lock()
				continue
			}
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Client disconnected :", err)
				remove(i)
				conn.Close()
				i--
				clientsMutex.Lock()
				continue
			}
			if string(msg) != "Next" {
				log.Println("Send error : ", string(msg))
			}

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
