package core

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Requests ...
type Requests struct{}

// fetch ...
func (r *Requests) fetch(url string) (*[]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// get body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

// GetVacancy ...
func (r *Requests) GetVacancy(url string) (*Vacancy, error) {
	data, err := r.fetch(url)
	if err != nil {
		return nil, err
	}

	var vacancy Vacancy
	if err := json.Unmarshal(*data, &vacancy); err != nil {
		return nil, err
	}

	return &vacancy, nil
}
