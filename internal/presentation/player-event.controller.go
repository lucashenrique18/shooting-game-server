package controllers

import (
	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
	"github.com/lucashenrique18/shooting-game-server/internal/domain/models"
	"github.com/mitchellh/mapstructure"
)

type playerEventPayload struct {
	PlayerID  string      `json:"playerId"`
	EventName string      `json:"eventName"`
	Payload   interface{} `json:"payload"`
}

type playerEventController struct {
	playerEventUseCase interfaces.PlayerEventUseCaseInterface
}

func NewPlayerEventController(playerEventUseCase interfaces.PlayerEventUseCaseInterface) ControllerInterface {
	return &playerEventController{
		playerEventUseCase: playerEventUseCase,
	}
}

func (a *playerEventController) Handle(request interface{}) (int, interface{}) {
	payload := playerEventPayload{}
	if err := mapstructure.Decode(request, &payload); err != nil {
		return HttpBadRequest(err.Error())
	}
	playerEvent := models.PlayerEvent{
		PlayerID:  payload.PlayerID,
		EventName: payload.EventName,
		Payload:   payload.Payload,
	}
	if err := a.playerEventUseCase.Execute(playerEvent); err != nil {
		return HttpInternalError(err)
	}
	return HttpOkNoContent()
}
