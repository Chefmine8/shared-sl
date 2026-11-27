package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
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

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected or read message error :", err)
			break
		}

		fmt.Printf("Message reçu : %s\n", message)

		reponse := fmt.Sprintf("Serveur Go a bien reçu : %s", message)
		err = conn.WriteMessage(messageType, []byte(reponse))
		if err != nil {
			log.Println("Erreur lors de l'envoi:", err)
			break
		}
	}
}

func StartWebSocket(port int) {
	http.HandleFunc("/ws", handleWebSocket)

	p := fmt.Sprintf(":%d", port)
	fmt.Println("Serveur start at http://localhost:" + p)
	log.Fatal(http.ListenAndServe(p, nil))
}
