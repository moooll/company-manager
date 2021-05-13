package domain

import (
	"context"

	"github.com/go-pg/pg/v10"
)

// Employee type describes company employee with id, full name, hire date, position in the company (developer, manager, quality assurance, business analyst) and company id.
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

func UpdEmployee(ctx *context.Context, db *pg.DB, emp Employee) error {
	// _, err := db.Model(emp).Insert()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func FindEmployee(ctx *context.Context, db *pg.DB, emp Employee) error {
	// _, err := db.Model(emp).Insert()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func UpdEmployeeFormData(ctx *context.Context, db *pg.DB, emp Employee) error {
	// _, err := db.Model(emp).Insert()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func DelEmployee(ctx *context.Context, db *pg.DB, emp Employee) error {
	// _, err := db.Model(emp).Insert()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func AddCompany(ctx *context.Context, db *pg.DB, emp Employee) error {
	// _, err := db.Model(emp).Insert()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func UpdCompany(ctx *context.Context, db *pg.DB, emp Employee) error {
	// _, err := db.Model(emp).Insert()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func FindCompany(ctx *context.Context, db *pg.DB, emp Employee) error {
	// _, err := db.Model(emp).Insert()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func DelCompany(ctx *context.Context, db *pg.DB, emp Employee) error {
	// _, err := db.Model(emp).Insert()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func GetEmlpoyeeList(ctx *context.Context, db *pg.DB, emp Employee) error {
	// _, err := db.Model(emp).Insert()
	// if err != nil {
	// 	return err
	// }

	return nil
}
