package models

type PlayerEvent struct {
	PlayerID  string      `json:"playerId"`
	EventName string      `json:"eventName"`
	Payload   interface{} `json:"payload"`
}
