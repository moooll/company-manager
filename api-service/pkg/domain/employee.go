package domain

// Employee type describes company employee with id, full name, hire date, position in the company (developer, manager, quality assurance, business analyst) and company id.
type Employee struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
	Surname    string `json:"surname"`
	HireDate   string `json:"hireDate"`
	Position   string `json:"position"`
	CompanyID  int64  `json:"companyId"`
}
