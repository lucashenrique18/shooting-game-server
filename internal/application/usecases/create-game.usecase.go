package usecases

import (
	"github.com/google/uuid"
	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
	"github.com/lucashenrique18/shooting-game-server/internal/domain/models"
)

type createMatchUseCase struct {
	gameLoopUseCase   interfaces.GameLoopUseCaseInterface
	matchesRepository interfaces.MatchesRepositoryInterface
}

func NewCreateMatchUseCase(gameLoopUseCase interfaces.GameLoopUseCaseInterface, matchesRepository interfaces.MatchesRepositoryInterface) interfaces.CreateMatchUseCaseInterface {
	return &createMatchUseCase{
		gameLoopUseCase:   gameLoopUseCase,
		matchesRepository: matchesRepository,
	}
}

func (c *createMatchUseCase) Execute(maxPlayers uint16, maxTime uint16, maxScore uint16) (string, error) {
	gameId := uuid.NewString()
	matchMap := models.MatchMap{
		Name:   "Default Map",
		ID:     "default-map",
		Width:  800,
		Height: 600,
	}
	match := models.NewMatch(gameId, maxPlayers, maxTime, maxScore, matchMap)
	c.matchesRepository.Save(match)
	go c.gameLoopUseCase.Execute(gameId)
	return gameId, nil
}
