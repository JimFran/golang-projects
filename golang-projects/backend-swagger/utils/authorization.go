package utils

import (
	"github.com/go-openapi/errors"
)

func BearerTokenAuth(token string) (interface{}, error) {
	if token == "Secret" {
		return token, nil
	}

	return nil, errors.New(401, "Incorrect API key auth")
}
