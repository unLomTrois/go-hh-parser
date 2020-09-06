package core

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Requests ...
type Requests struct{}

// Vacancy ...
type Vacancy struct {
	Data *Intermap
}

// Intermap is an alias type
// map of interfaces with string as a key
type Intermap = map[string]interface{}

// Interarr is an alias type
type Interarr = []interface{}

// fetch ...
func (r *Requests) fetch(url string) (*Intermap, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// body as bytes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// body as map of interfaces
	var bodyintermap Intermap
	if err = json.Unmarshal(body, &bodyintermap); err != nil {
		return nil, err
	}

	return &bodyintermap, nil
}

// GetVacancy ...
func (r *Requests) GetVacancy(url string) (*Vacancy, error) {
	respintermap, err := r.fetch(url)
	if err != nil {
		return nil, err
	}

	return &Vacancy{
		Data: respintermap,
	}, nil
}
