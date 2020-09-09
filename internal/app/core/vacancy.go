package core

// ShortVacancy ...
type ShortVacancy struct {
	ID  string `json:"id"`
	URL string `json:"url"`
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
	ShortVacancy
	Salary    *salary    `json:"salary"`
	Keyskills []KeySkill `json:"key_skills"`
}

// KeySkill ...
type KeySkill = map[string]string

// VacancyPage ...
type VacancyPage struct {
	Items     []ShortVacancy `json:"items"`
	Found     int            `json:"found"`
	Pages     int            `json:"pages"`
	PerPages  int            `json:"per_page"`
	Clusters  *[]Cluster     `json:"clusters"`
	Arguments interface{}    `json:"arguments"`
	AltURL    string         `json:"alternate_url"`
}

// Cluster ...
type Cluster struct {
	ID    string        `json:"id"`
	Name  string        `json:"name"`
	Items []interface{} `json:"items"`
}

// ShortVacancyList ...
type ShortVacancyList struct {
	Items *[]ShortVacancy `json:"items"`
}

// VacancyQueryParams ...
type VacancyQueryParams struct {
	Text     string `query:"text"`
	Area     string `query:"area"`
	NoMagic  bool   `query:"no_magic"`
	Page     int    `query:"page"`
	PerPage  int    `query:"per_page"`
	OrderBy  string `query:"order_by"`
	Clusters bool   `query:"clusters"`
}
