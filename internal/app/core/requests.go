package core

import (
	"bytes"
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
func (r *Requests) GetFullVacancy(url string) (*ShortVacancy, error) {
	data, err := r.fetch(url)
	if err != nil {
		return nil, err
	}

	// to save &
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	_ = enc.Encode(data)

	var vacancy ShortVacancy
	if err := json.Unmarshal(buf.Bytes(), &vacancy); err != nil {
		return nil, err
	}

	return &vacancy, nil
}

// SearchVacancies ...
func (r *Requests) SearchVacancies(params VacancyQueryParams) (*VacancyPage, error) {

	// get advance info
	info, err := r.getAdvanceVacancyInfo(params)
	if err != nil {
		return nil, err
	}

	if info.Found >= 2000 { // with clusters

		if params.Clusters {
			return r.searchByClusters(params)
		}
	}

	// without clusters

	var pages int = info.Pages / 100

	vacancyPage := *info

	// channel
	ch := make(chan *[]ShortVacancy, pages)

	// add goroutins
	for i := 0; i < pages; i++ {

		go func(page int) {

			list, err := r.getVacanciesFromPage(params, page)
			if err != nil {
				panic(err)
			}

			ch <- list
		}(i)
	}

	// get data from goroutins
	for {
		vacancyPage.Items = append(vacancyPage.Items, *<-ch...)

		if len(vacancyPage.Items) == info.Pages {
			break
		}
	}

	return &vacancyPage, nil
}

func (r *Requests) getVacanciesFromPage(params VacancyQueryParams, page int) (*[]ShortVacancy, error) {

	params.Page = page

	// build url for searching
	url, err := r.buildQueryParams(params)
	if err != nil {
		return nil, err
	}

	data, err := r.fetch(url.String())

	var vacancyList ShortVacancyList
	if err := json.Unmarshal(*data, &vacancyList); err != nil {
		return nil, err
	}

	return vacancyList.Items, nil
}

func (r *Requests) getAdvanceVacancyInfo(params VacancyQueryParams) (*VacancyPage, error) {
	params.PerPage = 0

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

	log.Println("LOL")
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
