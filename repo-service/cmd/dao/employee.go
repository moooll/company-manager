package dao

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type Employee struct {
	tableName  struct{} `pg:"genres"`
	ID         int64    `pg:"id"`
	Name       string   `pg:"name"`
	SecondName string   `pg:"second_name"`
	Surname    string   `pg:"surname"`
	HireDate   string   `pg:"hire_date"`
	Position   string   `pg:"position"`
	CompanyID  int64    `pg:"company_id"`
}

// AddEmployee function adds a new employee to the database
func AddEmployee(ctx *context.Context, db *pg.DB, emp Employee) error {
	_, err := db.Model(emp).Insert()
	if err != nil {
		return err
	}

	return nil
}
