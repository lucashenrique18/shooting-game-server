package usecases

import (
	"time"

	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
	"github.com/lucashenrique18/shooting-game-server/internal/infra/websocket"
)

const EVENT_MATCH_UPDATE = "GAMESTATE"

type gameLoopUseCase struct {
	tickRate          uint16
	matchesRepository interfaces.MatchesRepositoryInterface
}

func NewGameLoopUseCase(tickRate uint16, matchesRepository interfaces.MatchesRepositoryInterface) interfaces.GameLoopUseCaseInterface {
	return &gameLoopUseCase{
		tickRate:          tickRate,
		matchesRepository: matchesRepository,
	}
}

func (g *gameLoopUseCase) Execute(gameId string) error {
	tickRate := g.tickRate
	tickDuration := 1000 / tickRate
	ticker := time.NewTicker(time.Duration(tickDuration) * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		match, _ := g.matchesRepository.GetMatchByID(gameId)
		match.Mu.Lock()
		playersId := make([]string, 0, len(match.Players))
		for playerId := range match.Players {
			playersId = append(playersId, playerId)
		}
		websocket.Broadcast(playersId, EVENT_MATCH_UPDATE, match)
		match.Mu.Unlock()
	}

	return nil
}
