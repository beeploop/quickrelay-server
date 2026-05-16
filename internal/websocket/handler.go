package websocket

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	PONG_WAIT            = 60 * time.Second
	WRITE_DEADLINE       = 5 * time.Second
	READ_LIMIT           = 64 * 1024
	PING_TICKER_INTERVAL = 30 * time.Second
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnection() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("upgrade error: ", err.Error())
			return
		}

		done := make(chan struct{})

		conn.SetReadLimit(int64(READ_LIMIT))
		conn.SetPongHandler(func(appData string) error {
			conn.SetReadDeadline(time.Now().Add(PONG_WAIT))
			return nil
		})

		go func() {
			defer close(done)
			defer conn.Close()

			conn.SetReadDeadline(time.Now().Add(PONG_WAIT))

			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Println("unexpected closure: ", err.Error())
					}

					if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
						log.Println("client disconnected", err.Error())
					}

					return
				}

				fmt.Println("received: ", string(msg))

				// TODO: Handle ack/other messages
			}
		}()

		ticker := time.NewTicker(PING_TICKER_INTERVAL)
		defer ticker.Stop()

		go func() {
			select {
			case <-ticker.C:
				conn.SetWriteDeadline(time.Now().Add(WRITE_DEADLINE))

				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					log.Println("ping error: ", err.Error())
					return
				}

			case <-done:
				log.Println("connection closed")
				return
			}
		}()
	}
}
