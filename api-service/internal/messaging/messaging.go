package messaging

import "github.com/moooll/company-manager/api-service/pkg/domain"

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
	return domain.Employee{}, nil
}