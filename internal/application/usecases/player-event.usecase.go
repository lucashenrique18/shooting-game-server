package usecases

import (
	"log"

	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
	"github.com/lucashenrique18/shooting-game-server/internal/domain/models"
)

type PlayerEventUseCase struct {
	matchesRepository interfaces.MatchesRepositoryInterface
}

func NewPlayerEventUseCase(matchesRepository interfaces.MatchesRepositoryInterface) interfaces.PlayerEventUseCaseInterface {
	return &PlayerEventUseCase{
		matchesRepository: matchesRepository,
	}
}

func (h *PlayerEventUseCase) Execute(event models.PlayerEvent) error {
	match, exists := h.matchesRepository.GetMatchByPlayerId(event.PlayerID)
	if !exists {
		return nil
	}
	switch event.EventName {
	case "move":
		payload, ok := event.Payload.(map[string]interface{})
		if !ok {
			return nil
		}
		direction, dirOk := payload["direction"].(string)
		if dirOk {
			player, playerExists := match.Players[event.PlayerID]
			if playerExists {
				movePlayer(&player, direction, match.MatchMap)
				match.Players[event.PlayerID] = player
			}
		}
	default:
		log.Println("Unhandled event:", event.EventName)
	}
	return nil
}

func movePlayer(player *models.Player, direction string, matchMap models.MatchMap) {
	switch direction {
	case "up":
		if player.Y > 0 {
			player.Y -= 1
		}
	case "down":
		if player.Y < matchMap.Height-1 {
			player.Y += 1
		}
	case "left":
		if player.X > 0 {
			player.X -= 1
		}
	case "right":
		if player.X < matchMap.Width-1 {
			player.X += 1
		}
	}
}
