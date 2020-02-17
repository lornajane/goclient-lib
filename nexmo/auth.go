package nexmo

import "fmt"

type Auth interface {
	getCreds() []string
}

type KeySecretAuth struct {
	apiKey    string
	apiSecret string
}

func (auth *KeySecretAuth) getCreds() []string {
	creds := []string{auth.apiKey, auth.apiSecret}
	return creds
}

func CreateAuthFromKeySecret(apiKey string, apiSecret string) *KeySecretAuth {
	auth := new(KeySecretAuth)
	auth.apiKey = apiKey
	auth.apiSecret = apiSecret
	return auth
}
