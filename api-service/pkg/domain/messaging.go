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
			ID: 0,
			Name: "Aaron",
			SecondName: "Glen",
			Surname: "Smith",
			HireDate: "2020-09-01",
			Position: "developer",
			CompanyID: 1,
	}, nil
}


func UpdEmployeeFormData(emp domain.Employee) error {

	return nil
}

func DelEmployee()