package models

type MatchMap struct {
	ID     string `json:"id"`
	Name   string `json:"players"`
	Width  uint16 `json:"width"`
	Height uint16 `json:"height"`
}
