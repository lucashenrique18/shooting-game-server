package modules

import (
	"github.com/lucashenrique18/shooting-game-server/internal/application/usecases"
	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
)

type UseCasesModules struct {
	CreateMatchUseCase           interfaces.CreateMatchUseCaseInterface
	JoinMatchUseCase             interfaces.JoinMatchUseCaseInterface
	GameLoopUseCase              interfaces.GameLoopUseCaseInterface
	PlayerEventUseCase           interfaces.PlayerEventUseCaseInterface
	GetAllPossibleMatchesUseCase interfaces.GetAllPossibleMatchesUseCaseInterface
}

func NewUseCaseModules(repositoriesModules RepositoriesModules, tickRate uint16) UseCasesModules {
	gameloopUseCase := usecases.NewGameLoopUseCase(tickRate, repositoriesModules.MatchesRepository)
	createMatchUseCase := usecases.NewCreateMatchUseCase(gameloopUseCase, repositoriesModules.MatchesRepository)
	joinMatchUseCase := usecases.NewJoinMatchUseCase(repositoriesModules.MatchesRepository)
	playerEventUseCase := usecases.NewPlayerEventUseCase(repositoriesModules.MatchesRepository)
	getAllPossibleMatchesUseCase := usecases.NewGetAllPossibleMatchesUseCase(repositoriesModules.MatchesRepository)
	return UseCasesModules{
		CreateMatchUseCase:           createMatchUseCase,
		JoinMatchUseCase:             joinMatchUseCase,
		GameLoopUseCase:              gameloopUseCase,
		PlayerEventUseCase:           playerEventUseCase,
		GetAllPossibleMatchesUseCase: getAllPossibleMatchesUseCase,
	}
}
