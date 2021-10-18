package ports

import "github.com/alibaraneser/go-link-shortener-hexagonal/internal/core/domain"

type RedirectSerializer interface {
	Decode(input []byte) (*domain.Redirect, error)
	Encode(input *domain.Redirect) ([]byte, error)
}