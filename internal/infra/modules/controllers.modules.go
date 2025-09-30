package modules

import controllers "github.com/lucashenrique18/shooting-game-server/internal/presentation"

type ControllersModules struct {
	CreateMatchController           controllers.ControllerInterface
	JoinMatchController             controllers.ControllerInterface
	PlayerEventController           controllers.ControllerInterface
	GetAllPossibleMatchesController controllers.ControllerInterface
}

func NewControllersModules(useCasesModules UseCasesModules) ControllersModules {
	return ControllersModules{
		CreateMatchController:           controllers.NewCreateMatchController(useCasesModules.CreateMatchUseCase),
		JoinMatchController:             controllers.NewJoinMatchController(useCasesModules.JoinMatchUseCase),
		PlayerEventController:           controllers.NewPlayerEventController(useCasesModules.PlayerEventUseCase),
		GetAllPossibleMatchesController: controllers.NewGetAllPossibleMatchesController(useCasesModules.GetAllPossibleMatchesUseCase),
	}
}
