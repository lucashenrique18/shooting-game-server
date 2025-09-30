package interfaces

type CreateMatchUseCaseInterface interface {
	Execute(maxPlayers uint16, maxTime uint16, maxScore uint16) (string, error)
}
