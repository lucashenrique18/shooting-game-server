package controllers

import (
	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
	"github.com/mitchellh/mapstructure"
)

type createGamePayload struct {
	MaxPlayers uint16
	MaxTime    uint16
	MaxScore   uint16
}

type createMatchController struct {
	createMatchUseCase interfaces.CreateMatchUseCaseInterface
}

func NewCreateMatchController(createMatchUseCase interfaces.CreateMatchUseCaseInterface) ControllerInterface {
	return &createMatchController{
		createMatchUseCase: createMatchUseCase,
	}
}

func (c *createMatchController) Handle(request interface{}) (int, interface{}) {
	var payload createGamePayload
	err := mapstructure.Decode(request, &payload)
	if err != nil {
		return HttpBadRequest(err.Error())
	}
	response, err := c.createMatchUseCase.Execute(payload.MaxPlayers, payload.MaxTime, payload.MaxScore)
	if err != nil {
		return HttpBadRequest(err.Error())
	}
	return HttpOk(map[string]interface{}{"matchId": response})
}
