package models

import (
	"errors"
	"sync"
)

type Match struct {
	ID         string            `json:"id"`
	Players    map[string]Player `json:"players"`
	Winner     *Player           `json:"winner"`
	Status     string            `json:"status"`
	MaxPlayers uint16            `json:"maxPlayers"`
	MaxTime    uint16            `json:"maxTime"`
	MaxScore   uint16            `json:"max_score"`
	MatchMap   MatchMap          `json:"matchMap"`
	Mu         sync.Mutex
}

func NewMatch(id string, maxPlayers, maxTime, maxScore uint16, gameMap MatchMap) *Match {
	return &Match{
		ID:         id,
		Players:    make(map[string]Player),
		Status:     "waiting",
		MaxPlayers: maxPlayers,
		MaxTime:    maxTime,
		MaxScore:   maxScore,
		MatchMap:   gameMap,
	}
}

func (m *Match) IsFull() bool {
	return len(m.Players) >= int(m.MaxPlayers)
}

func (m *Match) GetPlayer(playerID string) (Player, error) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	player, exists := m.Players[playerID]
	if !exists {
		return Player{}, errors.New("player not found")
	}
	return player, nil
}

func (m *Match) AddPlayer(player Player) error {
	if m.IsFull() {
		return errors.New("match is full")
	}
	m.Mu.Lock()
	defer m.Mu.Unlock()
	m.Players[player.ID] = player
	return nil
}

func (m *Match) RemovePlayer(playerID string) {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	delete(m.Players, playerID)
}

func (m *Match) UpdatePlayer(player Player) error {
	m.Mu.Lock()
	defer m.Mu.Unlock()
	_, exists := m.Players[player.ID]
	if !exists {
		return errors.New("player not found in match")
	}
	m.Players[player.ID] = player
	return nil
}
