package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Requests ...
type Requests struct {
	*formatter
}

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

// GetFullVacancy ...
func (r *Requests) GetFullVacancy(url string) (*Vacancy, error) {
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

// SearchVacancies ...
func (r *Requests) SearchVacancies(params VacancyQueryParams) (*[]Vacancy, error) {

	// build url for searching
	u, err := r.buildQueryParams(params)
	if err != nil {
		return nil, err
	}

	log.Println(u)

	return nil, nil
}
