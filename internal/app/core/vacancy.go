package core

// Vacancy ...
type Vacancy struct {
	ID        string       `json:"id"`
	Keyskills []KeySkill   `json:"key_skills"`
	Salary    *interface{} `json:"salary"`
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
