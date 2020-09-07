package worker

import (
	"errors"
	"strconv"
)

func getPaginated(pageNumber int) (Page, error) {
	rest := newHttpClient()
	path := "/images"
	var result Page

	token, err := getBearerToken()
	if err != nil {
		return Page{}, err
	}

	resp, err := rest.
		R().
		SetAuthToken(token).
		SetResult(&result).
		SetQueryParam("page", strconv.Itoa(pageNumber)).
		Get(baseUrl + path)
	if err != nil {
		return Page{}, err
	}
	if resp.IsError() {
		return Page{}, errors.New(resp.String())
	}

	return result, nil
}

func getById(picId string) (Picture, error) {
	rest := newHttpClient()
	path := "/images/" + picId
	var result Picture

	token, err := getBearerToken()
	if err != nil {
		return Picture{}, err
	}

	resp, err := rest.
		R().
		SetAuthToken(token).
		SetResult(&result).
		Get(baseUrl + path)
	if err != nil {
		return Picture{}, err
	}
	if resp.IsError() {
		return Picture{}, errors.New(resp.String())
	}

	return result, nil
}
