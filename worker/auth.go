package worker

import (
	"errors"
	"time"
)

const apikey = "23567b218376f79d9415"

func getBearerToken() (string, error) {
	ds := newDS()

	// first check if token was already saved in cache
	token, found := ds.Get("token")
	if found {
		return token.(string), nil
	}

	rest := newHttpClient()
	path := "/auth"
	var result AuthResponse
	body := AuthRequest{ApiKey: apikey}

	resp, err := rest.
		R().
		SetBody(body).
		SetResult(&result).
		Post(baseUrl + path)
	if err != nil {
		return "", err
	}
	if resp.IsError() {
		return "", errors.New(resp.String())
	}

	// save token to cache
	ds.Set("token", result.Token, 3600*time.Second)

	return result.Token, nil
}
