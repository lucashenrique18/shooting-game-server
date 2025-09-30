package controllers

import (
	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
)

type getAllPossibleMatchesController struct {
	getAllPossibleMatchesUseCase interfaces.GetAllPossibleMatchesUseCaseInterface
}

func NewGetAllPossibleMatchesController(getAllPossibleMatchesUseCase interfaces.GetAllPossibleMatchesUseCaseInterface) ControllerInterface {
	return &getAllPossibleMatchesController{
		getAllPossibleMatchesUseCase: getAllPossibleMatchesUseCase,
	}
}

func (g *getAllPossibleMatchesController) Handle(request interface{}) (int, interface{}) {

	response, err := g.getAllPossibleMatchesUseCase.Execute()
	if err != nil {
		return HttpBadRequest(err.Error())
	}
	return HttpOk(response)
}
