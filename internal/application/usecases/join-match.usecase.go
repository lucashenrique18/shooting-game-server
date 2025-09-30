package usecases

import (
	"errors"

	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
	"github.com/lucashenrique18/shooting-game-server/internal/domain/models"
)

type joinMatchUseCase struct {
	matchesRepository interfaces.MatchesRepositoryInterface
}

func NewJoinMatchUseCase(matchesRepository interfaces.MatchesRepositoryInterface) interfaces.JoinMatchUseCaseInterface {
	return &joinMatchUseCase{
		matchesRepository: matchesRepository,
	}
}

func (c *joinMatchUseCase) Execute(matchId, playerName string) (string, error) {
	match, exists := c.matchesRepository.GetMatchByID(matchId)
	if !exists {
		return "", errors.New("match not found")
	}
	player := models.Player{
		ID:     playerName,
		Name:   playerName,
		Score:  0,
		X:      0,
		Y:      0,
		Health: 100,
	}
	match.Mu.Lock()
	defer match.Mu.Unlock()
	if uint16(len(match.Players)) >= match.MaxPlayers {
		return "", errors.New("match is full")
	}
	match.Players[player.ID] = player
	c.matchesRepository.PushPlayerInMatch(player, match.ID)
	return player.ID, nil
}
