package repositories

import (
	"errors"

	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
	"github.com/lucashenrique18/shooting-game-server/internal/domain/models"
)

type matchesRepository struct {
	matches            map[string]*models.Match
	playersIdMatchesId map[string]string
}

func NewMatchesRepository() interfaces.MatchesRepositoryInterface {
	return &matchesRepository{
		matches:            make(map[string]*models.Match),
		playersIdMatchesId: make(map[string]string),
	}
}

func (m *matchesRepository) GetMatchByID(id string) (*models.Match, bool) {
	match, exists := m.matches[id]
	return match, exists
}

func (m *matchesRepository) Save(match *models.Match) {
	m.matches[match.ID] = match
}

func (m *matchesRepository) GetMatchByPlayerId(playerId string) (*models.Match, bool) {
	matchId, exists := m.playersIdMatchesId[playerId]
	if !exists {
		return nil, false
	}
	match, exists := m.matches[matchId]
	if !exists {
		return nil, false
	}
	return match, true
}

func (m *matchesRepository) linkPlayerToMatch(playerId, matchId string) {
	m.playersIdMatchesId[playerId] = matchId
}

func (m *matchesRepository) PushPlayerInMatch(player models.Player, matchId string) error {
	match, exists := m.GetMatchByID(matchId)
	if !exists {
		return errors.New("match not found")
	}
	match.Players[player.ID] = player
	m.linkPlayerToMatch(player.ID, matchId)
	return nil
}

func (m *matchesRepository) GetAllMatches() []*models.Match {
	var allMatches []*models.Match
	for _, match := range m.matches {
		allMatches = append(allMatches, match)
	}
	return allMatches
}
