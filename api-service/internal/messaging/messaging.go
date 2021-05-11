package messaging

import (
	"errors"

	"github.com/moooll/company-manager/api-service/pkg/domain"
)

// messaging package implements call to message broker for message transfer to repi-service
// it hides the call to specific message broker so that later it can be changed

//
func AddEmployee(emp domain.Employee) error {
	return nil
}

func UpdEmployee(emp domain.Employee) error {
	return nil
}

func FindEmployee(id int64) (domain.Employee, error) {
	if id != 0 {
		return domain.Employee{}, errors.New("wrong key")
	}

	return domain.Employee{
		ID:         0,
		Name:       "Aaron",
		SecondName: "Glen",
		Surname:    "Smith",
		HireDate:   "2020-09-01",
		Position:   "developer",
		CompanyID:  1,
	}, nil
}

func UpdEmployeeFormData(emp domain.Employee) error {

	return nil
}

func DelEmployee(id int64) error {
	if id != 0 {
		return errors.New("not found")
	}

	return nil
}


func AddCompany(emp domain.Company) error {
	return nil
}

func UpdCompany(emp domain.Company) error {
	return nil
}

func FindCompany(id int64) (domain.Company, error) {
	return domain.Company{
		ID: 0,
		Name: "Qulix Systems",
		LegalForm: "OOO",
	  }, nil
}

func DelCompany(id int64) error {
	return nil
}

func GetEmlpoyeeList(id int64) ([]domain.Employee, error) {
	return []domain.Employee{}, nil
}