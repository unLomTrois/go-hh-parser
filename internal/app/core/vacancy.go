package core

// Vacancy ...
type Vacancy struct {
	ID     string  `json:"id"`
	Salary *salary `json:"salary"`
	URL    string  `json:"url"`
}

// salary ...
type salary struct {
	Currency string `json:"currency"`
	From     *int   `json:"from"`
	To       *int   `json:"to"`
	Gross    bool   `json:"gross"`
}

// FullVacancy ...
type FullVacancy struct {
	Vacancy
	Keyskills []KeySkill `json:"key_skills"`
}

// VacancyPage ...
type VacancyPage struct {
	Items     []Vacancy   `json:"items"`
	Found     int         `json:"found"`
	Pages     int         `json:"pages"`
	PerPages  int         `json:"per_page"`
	Clusters  interface{} `json:"clusters"`
	Arguments interface{} `json:"arguments"`
	AltURL    string      `json:"alternate_url"`
}

// KeySkill ...
type KeySkill = map[string]string

// VacancyQueryParams ...
type VacancyQueryParams struct {
	Text    string `query:"text"`
	Area    string `query:"area"`
	NoMagic bool   `query:"no_magic"`
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
	OrderBy string `query:"order_by"`
}
