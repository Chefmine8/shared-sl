package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/gorilla/websocket"
)

func playTrain(conn *websocket.Conn) {
	ctx, stopAudio := context.WithCancel(context.Background())
	go PlaySound(ctx, "./oxidvideos-train-over-bridge-inside-ambiance-617983.mp3")
	cmd := exec.Command("./sl")
	cmd.Stdout = os.Stdout
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal("Pipe init error : ", err)
	}

	if err := cmd.Start(); err != nil {

		log.Println("Exec error :", err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "END") {
			err = conn.WriteMessage(1, []byte("Next"))
			if err != nil {
				log.Fatal("Write error : ", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Scanner error : ", err)
	}

	if err := cmd.Wait(); err != nil {
		log.Println("Exec error : ", err)
	}
	stopAudio()
}

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
		//fmt.Printf("Received msg : %s\n", msg)

		if string(msg) == "Your turn" {
			go playTrain(conn)
		} else if string(msg) == "Sound" {
			ctx, _ := context.WithCancel(context.Background())
			go PlaySound(ctx, "./caboose_canine-steam-locomotive-fouled-5-chime-375193.mp3")
		}

	}
}
