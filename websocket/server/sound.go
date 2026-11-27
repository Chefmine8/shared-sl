package server

import (
	"math/rand/v2"
	"time"

	"github.com/gorilla/websocket"
)

func soundLoop() {
	for {
		time.Sleep(time.Duration(15+rand.IntN(16)) * time.Second)

		activeMutex.Lock()
		conn := activeConn
		activeMutex.Unlock()

		if conn != nil {
			_ = writeWS(conn, websocket.TextMessage, []byte("Sound"))
		}
	}
}
