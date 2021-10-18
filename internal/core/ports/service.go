package ports

import "github.com/alibaraneser/go-link-shortener-hexagonal/internal/core/domain"

type RedirectService interface {
	Find(code string) (*domain.Redirect, error)
	Save(redirect *domain.Redirect) error
}