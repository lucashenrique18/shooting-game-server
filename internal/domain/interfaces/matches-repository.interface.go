package interfaces

import "github.com/lucashenrique18/shooting-game-server/internal/domain/models"

type MatchesRepositoryInterface interface {
	Save(match *models.Match)
	PushPlayerInMatch(player models.Player, matchId string) error
	GetMatchByID(id string) (*models.Match, bool)
	GetMatchByPlayerId(playerId string) (*models.Match, bool)
	GetAllMatches() []*models.Match
}
