package cmd

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/moooll/company-manager/api-service/pkg/domain"
	"github.com/moooll/company-manager/repo-service/dao"
)

// Add a new employee to the company, sends employee object that needs to be added to the company
// If something goes wrong 405 - invalid input
func addEmployee(c *fiber.Ctx) error {
	if !c.Is("application/json") {
		c.SendStatus(400)
		c.Context().SetContentType("application/json")
		c.SendString("Content type is application/json only")
		return errors.New("Wrong content type")
	}
	emp := domain.Employee{}
	err := c.BodyParser(emp)
	if err != nil {
		c.Context().SetContentType("application/json")
		c.SendStatus(400)
		c.SendString("Invalid input")
		return err
	}

	err = dao.AddEmployee(emp)
	if err != nil {
		c.SendString("Could not add employee")
		return err
	}

	return nil
}

// app.Put("/employee", updEmployee())
func updEmployee(c *fiber.Ctx) error {

	return nil
}

// app.Get("/employee/:employeeID", findEmployee())
func findEmployee(c *fiber.Ctx) error {

	return nil
}

// app.Post("/employee/:employeeID", updEmployeeFormData())
func updEmployeeFormData(c *fiber.Ctx) error {

	return nil
}

// app.Delete("/employee/:employeeID", delEmplaoyee())
func delEmplaoyee(c *fiber.Ctx) error {

	return nil
}

// app.Put("/company", addCompany())
func addCompany(c *fiber.Ctx) error {

	return nil
}

// app.Post("/company", updCompany())
func updCompany(c *fiber.Ctx) error {

	return nil
}

// app.Get("/company/:companyID", findCompany())
func findCompany(c *fiber.Ctx) error {

	return nil
}

// app.Post("/company/:companyID", updCompanyFormData())
func updCompanyFormData(c *fiber.Ctx) error {

	return nil
}

// app.Delete("/company/:companyID", delCompany())
func delCompany(c *fiber.Ctx) error {

	return nil
}

// app.Get("/company/:companyID/employees", getEmployeeList())
func getEmployeeList(c *fiber.Ctx) error {

	return nil
}
