package domain

import "time"

// Employee type describes company employee with id, full name, hire date, position in the company (developer, manager, quality assurance, business analyst) and company id.
type Employee struct {
	ID         int64
	Name       string
	SecondName string
	Surname    string
	HireDate   time.Time
	Position   string
	CompanyID  int64
}
