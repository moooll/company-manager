package cmd

import (
	"github.com/gofiber/fiber/v2"
)

func Serve() {
	app := fiber.New()
	app.Post("/employee", addEmployee)
	app.Put("/employee", updEmployee)
	app.Get("/employee/:id", findEmployee)
	app.Post("/employee/:id", updEmployeeFormData)
	app.Delete("/employee/:id", delEmployee)

	app.Post("/company", addCompany)
	app.Put("/company", updCompany)
	app.Get("/company/:id", findCompany)
	app.Post("/company/:id", updCompanyFormData)
	app.Delete("/company/:id", delCompany)
	app.Get("/company/:id/employees", getEmployeeList)

	// todo: were to launch server????
	app.Listen(":3000")
}
