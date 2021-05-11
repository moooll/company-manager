package test

import (
	"context"
	"os"
	"testing"

	"go.uber.org/zap"

	"github.com/go-pg/pg/v10"

	"repo-service/cmd/dao"
	"repo-service/constants"
)

func TestAddEmployee(t *testing.T) {
	ctx := context.Background()
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv(constants.DatabaseHost),
		User:     os.Getenv(constants.DatabaseUser),
		Password: os.Getenv(constants.DatabasePassword),
		Database: os.Getenv(constants.DatabaseName),
	})

	defer db.Close()

	// if err := db.Ping(ctx); err != nil {
	// 	zap.L().Error("could not launch the db:" + err.Error())
	// 	return
	// }
	emp := dao.Employee{
		ID:         0,
		Name:       "Aaron",
		SecondName: "Glen",
		Surname:    "Smith",
		HireDate:   "2020-09-01",
		Position:   "developer",
		CompanyID:  1,
	}
	err := dao.AddEmployee(&ctx, db, emp)
	if err != nil {
		zap.L().Error(err.Error())
		t.Fail()
	}
}
