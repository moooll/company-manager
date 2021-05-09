package cmd

import (
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	app := fiber.New()
	app.Post("/employee", addEmployee)
	app.Put("/employee", updEmployee)
	app.Get("/employee/:employeeID", findEmployee)
	app.Post("/employee/:employeeID", updEmployeeFormData)
	app.Delete("/employee/:employeeID", delEmployee)

	app.Put("/company", addCompany)
	app.Post("/company", updCompany)
	app.Get("/company/:companyID", findCompany)
	app.Post("/company/:companyID", updCompanyFormData)
	app.Delete("/company/:companyID", delCompany)
	app.Get("/company/:companyID/employees", getEmployeeList)

	// todo: were to launch server????
	app.Listen(":3000")
}
