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

// HTTPResponse ...
type HTTPResponse struct {
	URL      string
	Response *http.Response
}

// fetch ...
func (r *Requests) fetch(url string) (*HTTPResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return &HTTPResponse{
		URL:      url,
		Response: resp,
	}, nil
}

// GetVacancy ...
func (r *Requests) GetVacancy(url string) (*Vacancy, error) {
	httpresp, err := r.fetch(url)
	if err != nil {
		return nil, err
	}

	resp := httpresp.Response

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		bodyString := string(bodyBytes)

		return &Vacancy{
			Data: bodyString,
		}, nil
	}

	// status code != ok
	return nil, err
}
