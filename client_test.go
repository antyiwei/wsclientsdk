package wsclientsdk

import (
	"log"
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
)

func TestAxb(t *testing.T) {

	wsc := New("axbservice-service:30005", http.Header{
		"APP": []string{
			"sipagent",
		}, "MER": []string{
			"weiliao",
		},
	})

	wsc.Ready= func() {
		err :=wsc.WriteMessage(websocket.PingMessage,[]byte{})

		log.Println(err)

	}


	select {

	}
}

func TestNew(t *testing.T) {

	wsc := New("localhost:9999", http.Header{
		"APP": []string{
			"sipagent",
		}, "MER": []string{
			"weiliao",
		},
	})

	wsc.Ready= func() {
		err :=wsc.WriteMessage(websocket.PingMessage,[]byte{})

		log.Println(err)

	}


	select {

	}
}
