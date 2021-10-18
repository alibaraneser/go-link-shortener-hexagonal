package domain

import "errors"

var (
	RedirectNotFound = errors.New("redirect not found")
	RedirectSaveError = errors.New("error occurred while adding redirect")
	EncodingError = errors.New("error occurred while encoding")
	DecodingError = errors.New("error occurred while decoding")
	ConnectionError = errors.New("database connection failed")
)
