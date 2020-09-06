package core

// Vacancy ...
type Vacancy struct {
	ID        string       `json:"id"`
	Keyskills []KeySkill   `json:"key_skills"`
	Salary    *interface{} `json:"salary"`
}

// KeySkill ...
type KeySkill = map[string]string
