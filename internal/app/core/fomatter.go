package core

import (
	"fmt"
	"net/url"
	"reflect"
)

type formatter struct {
	baselink string
}

func (f *formatter) buildQueryParams(params VacancyQueryParams) (*url.URL, error) {

	// parse baselink
	u, err := url.Parse(f.baselink)
	if err != nil {
		return nil, err
	}

	// fill queries
	q := u.Query()
	v := reflect.ValueOf(params)
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Tag.Get("query")
		value := fmt.Sprintf("%v", v.Field(i).Interface())

		q.Set(key, value)
	}

	u.RawQuery = q.Encode()

	return u, nil
}
