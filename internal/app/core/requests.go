package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

	// to save &
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	_ = enc.Encode(data)

	var vacancy Vacancy
	if err := json.Unmarshal(buf.Bytes(), &vacancy); err != nil {
		return nil, err
	}

	return &vacancy, nil
}

// SearchVacancies ...
func (r *Requests) SearchVacancies(params VacancyQueryParams) (*VacancyPage, error) {

	// with clusters
	if params.Clusters {
		return r.searchByClusters(params)
	}

	// without clusters
	params.Clusters = false

	// build url for searching
	url, err := r.buildQueryParams(params)
	if err != nil {
		return nil, err
	}

	data, err := r.fetch(url.String())

	var vacancyPage VacancyPage
	if err := json.Unmarshal(*data, &vacancyPage); err != nil {
		return nil, err
	}

	return &vacancyPage, nil
}

func (r *Requests) searchByClusters(params VacancyQueryParams) (*VacancyPage, error) {
	// build url for searching
	url, err := r.buildQueryParams(params)
	if err != nil {
		return nil, err
	}

	data, err := r.fetch(url.String())

	var vacancyPage VacancyPage
	if err := json.Unmarshal(*data, &vacancyPage); err != nil {
		return nil, err
	}

	return &vacancyPage, nil
}
