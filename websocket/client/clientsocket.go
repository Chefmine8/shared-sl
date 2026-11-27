package client

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gorilla/websocket"
)

func StartWebSocket(addr string, order string) {
	url := fmt.Sprintf("ws://%s/ws", addr)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Dial error : ", err)
	}
	defer conn.Close()

	fmt.Println("Client connected !")
	err = conn.WriteMessage(websocket.TextMessage, []byte(order))
	if err != nil {
		log.Fatal("Send error : ", err)
		return
	}

	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Fatal("Connection close by host : ", err)
		return
	}

	if string(msg) != "Your are " {
		log.Fatal("Wrong msg : ", msg)
		return
	}
	fmt.Printf("Received msg : %s\n", msg)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("Connection close by host : ", err)
			return
		}
		fmt.Printf("Received msg : %s\n", msg)
		cmd := exec.Command("../sl/sl")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Println("Exec error :", err)
		}
	}
}
