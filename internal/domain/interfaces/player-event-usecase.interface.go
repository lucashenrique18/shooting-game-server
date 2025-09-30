package interfaces

import "github.com/lucashenrique18/shooting-game-server/internal/domain/models"

type PlayerEventUseCaseInterface interface {
	Execute(event models.PlayerEvent) error
}
