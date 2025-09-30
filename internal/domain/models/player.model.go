package models

type Player struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	X      uint16 `json:"x"`
	Y      uint16 `json:"y"`
	Health uint16 `json:"health"`
	Score  uint16 `json:"score"`
}
