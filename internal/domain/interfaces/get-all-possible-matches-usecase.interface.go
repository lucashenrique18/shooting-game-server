package interfaces

type GetAllPossibleMatchesUseCaseInterface interface {
	Execute() ([]string, error)
}
