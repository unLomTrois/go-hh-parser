package core

import (
	"io/ioutil"
	"net/http"
)

// Requests ...
type Requests struct{}

// Vacancy ...
type Vacancy struct {
	Data string
}

// fetch ...
func (r *Requests) fetch(url string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		bodyString := string(bodyBytes)

		return &bodyString, nil
	}
	// status != ok
	return &resp.Status, nil
}

// GetVacancy ...
func (r *Requests) GetVacancy(url string) (*Vacancy, error) {
	bodyString, err := r.fetch(url)
	if err != nil {
		return nil, err
	}

	return &Vacancy{
		Data: *bodyString,
	}, nil
}
