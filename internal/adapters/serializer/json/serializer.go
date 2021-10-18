package json

import (
	"encoding/json"
	"github.com/alibaraneser/go-link-shortener-hexagonal/internal/core/domain"
	errs "github.com/pkg/errors"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*domain.Redirect, error) {
	redirect := &domain.Redirect{}

	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errs.Wrapf(domain.DecodingError, "serialize.Decode")
	}

	return redirect, nil
}

func (r *Redirect) Encode(input *domain.Redirect) ([]byte, error) {
	output, err := json.Marshal(input)

	if err != nil {
		return nil, errs.Wrapf(domain.EncodingError, "serialize.Encode")
	}

	return output, nil
}
