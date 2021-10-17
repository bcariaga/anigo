package fetching

import "github.com/bcariaga/anigo/src/animes/internal/animes"

//interface of fetching service
type Service interface {
	FindByTerm(term string) ([]*animes.Anime, error)
}

//private
type service struct {
	animes animes.Respository
}

//constructor
func NewService(aR animes.Respository) Service {
	return &service{animes: aR}
}

//implementation of Service
func (s *service) FindByTerm(term string) ([]*animes.Anime, error) {
	animes, err := s.animes.Find(term)

	if err != nil {
		return nil, err
	}

	return animes, nil
}
