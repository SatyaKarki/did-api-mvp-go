package domain

import "errors"

var (
	ErrInvalidDID         = errors.New("invalid DID")
	ErrDIDNotFound        = errors.New("DID not found")
	ErrCredentialNotValid = errors.New("credential is not valid")
)
