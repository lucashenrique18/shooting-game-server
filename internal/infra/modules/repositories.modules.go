package modules

import (
	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
	"github.com/lucashenrique18/shooting-game-server/internal/infra/repositories"
)

type RepositoriesModules struct {
	MatchesRepository interfaces.MatchesRepositoryInterface
}

func NewRepositoriesModules() RepositoriesModules {
	matchesRepository := repositories.NewMatchesRepository()
	return RepositoriesModules{
		MatchesRepository: matchesRepository,
	}
}
