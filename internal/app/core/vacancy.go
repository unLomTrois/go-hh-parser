package core

// Vacancy ...
type Vacancy struct {
	ID        string              `json:"id"`
	Keyskills []map[string]string `json:"key_skills"`
	Salary    *interface{}        `json:"salary"`
}
