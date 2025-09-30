package controllers

import (
	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
	"github.com/mitchellh/mapstructure"
)

type joinMatch struct {
	MatchId    string
	PlayerName string
}

type joinMatchController struct {
	joinMatchUseCase interfaces.JoinMatchUseCaseInterface
}

func NewJoinMatchController(joinMatchUseCase interfaces.JoinMatchUseCaseInterface) ControllerInterface {
	return &joinMatchController{
		joinMatchUseCase: joinMatchUseCase,
	}
}

func (j *joinMatchController) Handle(request interface{}) (int, interface{}) {
	var payload joinMatch
	err := mapstructure.Decode(request, &payload)
	if err != nil {
		return HttpBadRequest(err.Error())
	}
	result, err := j.joinMatchUseCase.Execute(payload.MatchId, payload.PlayerName)
	if err != nil {
		return HttpBadRequest(err.Error())
	}
	return HttpOk(map[string]interface{}{
		"playerId": result,
	})
}
