package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	controllers "github.com/lucashenrique18/shooting-game-server/internal/presentation"
)

var clientsInstance = make(map[string]*websocket.Conn)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func SendToPlayer(playerId, eventName string, message interface{}) error {
	playerConnection, exists := clientsInstance[playerId]
	if !exists {
		return fmt.Errorf("player %s does not exist", playerId)
	}
	err := playerConnection.WriteJSON(map[string]interface{}{
		"event": eventName,
		"data":  message,
	})
	return err
}

func Broadcast(playersId []string, eventName string, message interface{}) error {
	for _, conn := range playersId {
		SendToPlayer(conn, eventName, message)
	}
	return nil
}

func ServeWs(w http.ResponseWriter, r *http.Request, playerEventController controllers.ControllerInterface) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	playerId := r.URL.Query().Get("playerId")
	err = EnterInRoom(playerId, conn)
	if err != nil {
		conn.Close()
		return
	}
	go ServeRoom(playerId, conn, playerEventController)
}

func ServeRoom(playerId string, conn *websocket.Conn, playerEventController controllers.ControllerInterface) {
	defer delete(clientsInstance, playerId)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		println("Received message:", string(msg))
		var eventPayload map[string]interface{}
		err = json.Unmarshal(msg, &eventPayload)
		if err != nil {
			println("Error unmarshalling message:", err.Error())
			continue
		}
		if eventPayload["eventName"] == "ping" {
			sendPong(conn)
			continue
		}
		eventPayload["playerId"] = playerId
		_, resp := playerEventController.Handle(eventPayload)
		if resp != nil {
			println("Error handling message:", resp)
			continue
		}
	}
}

func EnterInRoom(playerId string, conn *websocket.Conn) error {
	clientsInstance[playerId] = conn
	return nil
}

func sendPong(conn *websocket.Conn) error {
	err := conn.WriteJSON(map[string]interface{}{
		"event": "pong",
	})
	return err
}
