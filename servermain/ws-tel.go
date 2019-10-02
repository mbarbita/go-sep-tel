package servermain

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// wsMessage handles browser requests to /msg/
func wsMessage(w http.ResponseWriter, r *http.Request) {

	// upgrade to websocket
	c, err := upgrader.Upgrade(w, r, nil)
	defer c.Close()
	if err != nil {
		log.Println("upgrade err:", err)
		return
	}

	// handle websocket incoming browser messages
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read err:", err)
			// cc <- true
			return
		}
		// log.Printf("recv: %s", message)
		// log.Printf("sending reply to: %s\n", message)

		req := strings.ToLower(string(message))
		if req == "" {
			continue
		}
		response := ""

		// process map
		for k, v := range telMap {
			k1 := strings.ToLower(k)
			v1 := strings.ToLower(v)
			if strings.Contains(k1, req) || strings.Contains(v1, req) ||
				req == "*" {
				response += k + " " + v + "\n"
				// log.Println(req, k, v, response)
			}
		}

		// mesage type = 1
		err = c.WriteMessage(1, []byte(response))
		if err != nil {
			log.Println("ws write err:", err)
			break
			// return
		}
	}
}
