package interfaces

type JoinMatchUseCaseInterface interface {
	Execute(matchID, playerName string) (string, error)
}
