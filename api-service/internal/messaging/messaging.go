package messaging

import (
	"errors"
	
	"api-service/pkg/domain"

	"github.com/moooll/company-manager/rabbit"
	"go.uber.org/zap"
)

// messaging package implements call to message broker for message transfer to repi-service
// it hides the call to specific message broker so that later it can be changed

//
// func ListenAndRoute(Function string, msg chan(rabbit.Msg), er chan(error)) {
// 	err := rabbit.Receive("from-db", msg)
// 	if err != nil {
// 		er <- err
// 		return
// 	}

// 	wait := make(chan bool)
// 	<- wait
// }

func AddEmployee(emp domain.Employee) error {
	body := rabbit.Msg{
		Function: "AddEmployee",
		Entity:   emp,
	}
	err := rabbit.Send(body, "to-db")
	if err != nil {
		return err
	}

	return nil
}

func UpdEmployee(emp domain.Employee) error {
	body := rabbit.Msg{
		Function: "UpdEmployee",
		Entity:   emp,
	}
	err := rabbit.Send(body, "to-db")
	if err != nil {
		return err
	}

	return nil
}

func FindEmployee(id int64) (domain.Employee, error) {
	body := rabbit.Msg{
		Function: "FindEmployee",
		Entity:   id,
	}
	err := rabbit.Send(body, "to-db")
	if err != nil {
		return err
	}
	zap.L().Info("message is sent")
	// go func() {

	// }
	// err := rabbit.
	return domain.Employee{}, nil
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
		ID:        0,
		Name:      "Qulix Systems",
		LegalForm: "OOO",
	}, nil
}

func DelCompany(id int64) error {
	return nil
}

func GetEmlpoyeeList(id int64) ([]domain.Employee, error) {
	return []domain.Employee{}, nil
}
