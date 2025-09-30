package interfaces

type GameLoopUseCaseInterface interface {
	Execute(gameId string) error
}
