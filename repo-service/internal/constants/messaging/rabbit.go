package rabbit

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/streadway/amqp"

	"github.com/moooll/company-manager/repo-service/cmd/domain"
	"github.com/moooll/company-manager/rabbit"
)


type dbInfo struct {
	ctx *context.Context
	db *pg.DB
}
func Listen(dbInfo *dbInfo, ch amqp.Channel, mes chan (rabbit.Msg), er chan (error)) {
	err := rabbit.Receive("to-db", mes, ch)
	if err != nil {
		er <- err
	}
	go func() {
		for m := range mes {
			switch m.Function {
			case "AddEmployee":
				domain.AddEmployee(dbInfo.ctx, dbInfo.db, mes.Entity)
			case "UpdEmployee":
				domain.UpdEmployee(dbInfo.ctx, dbInfo.db, mes.Entity)
			case "FindEmployee":
				domain.FindEmployee(dbInfo.ctx, dbInfo.db, mes.Entity)
			case "UpdEmployeeFormData":
				domain.UpdEmployeeFormData(dbInfo.ctx, dbInfo.db, mes.Entity)
			case "DelEmployee":
				domain.DelEmployee(dbInfo.ctx, dbInfo.db, mes.Entity)
			case "AddCompany":
				domain.AddCompany(dbInfo.ctx, dbInfo.db, mes.Entity)
			case "UpdCompany":
				domain.UpdCompany(dbInfo.ctx, dbInfo.db, mes.Entity)
			case "FindCompany":
				domain.FindCompany(dbInfo.ctx, dbInfo.db, mes.Entity)
			case "DelCompany":
				domain.DelCompany(dbInfo.ctx, dbInfo.db, mes.Entity)
			case "GetEmlpoyeeList":
				domain.GetEmlpoyeeList(dbInfo.ctx, dbInfo.db, mes.Entity)
			}
		}
	}()
	wait := make(chan bool)
	<-wait
}
