package service

import "jwt/internal/domain"

type Places struct {
	repo PlacesRepository
}

type PlacesRepository interface {
	GetPlaces(limit, offset int) (domain.Answer, int, error)
	GetToken() domain.Token
}

func NewServices(repo PlacesRepository) *Places {
	return &Places{
		repo: repo,
	}
}

func (p *Places) GetPlaces(limit, offset int) (domain.Answer, int, error) {
	return p.repo.GetPlaces(limit, offset)
}

func (p *Places) GetToken() domain.Token {
	return p.repo.GetToken()
}
