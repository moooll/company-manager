package domain

import (
	"context"

	"github.com/go-pg/pg/v10"
)

// Employee type describes company employee with id, full name, hire date, position in the company (developer, manager, quality assurance, business analyst) and company id.
type Employee struct {
	tableName  struct{} `json:"genres"`
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	SecondName string   `json:"second_name"`
	Surname    string   `json:"surname"`
	HireDate   string   `json:"hire_date"`
	Position   string   `json:"position"`
	CompanyID  int64    `json:"company_id"`
}

// AddEmployee function adds a new employee to the database
func AddEmployee(ctx *context.Context, db *pg.DB, emp Employee) error {
	_, err := db.Model(emp).Insert()
	if err != nil {
		return err
	}

	return nil
}
