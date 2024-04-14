package main

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

func reader(conn *websocket.Conn) {
	for {
		messages, p, err := conn.ReadMessage()
		if err != nil {
			log.Print(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messages, p); err != nil {
			log.Print(err)
			return
		}
	}
}
