package business

import (
	"github.com/alibaraneser/go-link-shortener-hexagonal/internal/core/domain"
	"github.com/alibaraneser/go-link-shortener-hexagonal/internal/core/ports"
	"github.com/alibaraneser/go-link-shortener-hexagonal/pkg/generator"
)

type redirectService struct {
	redirectRepo ports.RedirectRepository
}

func NewRedirectService(redirectRepo ports.RedirectRepository) ports.RedirectService {
	return &redirectService{redirectRepo: redirectRepo}
}

func (r redirectService) Find(code string) (*domain.Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r redirectService) Save(redirect *domain.Redirect) error {
	if redirect.Code == "" {
		redirect.Code = generator.RandomGenerate(10)
	}
	return r.redirectRepo.Save(redirect)
}
