package usecases

import (
	"errors"

	"github.com/lucashenrique18/shooting-game-server/internal/domain/interfaces"
)

type getAllPossibleMatchesUsecase struct {
	matchesRepository interfaces.MatchesRepositoryInterface
}

func NewGetAllPossibleMatchesUseCase(matchesRepository interfaces.MatchesRepositoryInterface) interfaces.GetAllPossibleMatchesUseCaseInterface {
	return &getAllPossibleMatchesUsecase{
		matchesRepository: matchesRepository,
	}
}
func (g *getAllPossibleMatchesUsecase) Execute() ([]string, error) {
	var availableMatches []string
	matches := g.matchesRepository.GetAllMatches()
	for _, match := range matches {
		if uint16(len(match.Players)) < match.MaxPlayers {
			availableMatches = append(availableMatches, match.ID)
		}
	}
	if len(availableMatches) == 0 {
		return nil, errors.New("no available matches")
	}
	return availableMatches, nil
}
