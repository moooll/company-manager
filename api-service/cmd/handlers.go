package cmd

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/moooll/company-manager/api-service/internal/messaging"
	"github.com/moooll/company-manager/api-service/pkg/domain"
)

// Add a new employee to the company, sends employee object that needs to be added to the company
// If something goes wrong 405 - invalid input
func addEmployee(c *fiber.Ctx) error {
	// if !c.Is("application/json") {
	// 	c.SendStatus(400)
	// 	c.Context().SetContentType("application/json")
	// 	c.SendString("Content type is application/json only")
	// 	return errors.New("Wrong content type")
	// }
	emp := domain.Employee{}
	err := c.BodyParser(&emp)
	if err != nil {
		c.Context().SetContentType("application/json")
		c.SendStatus(400)
		c.SendString("Invalid input")
		return err
	}

	err = messaging.AddEmployee(emp)
	if err != nil {
		c.SendString("Could not add employee: " + err.Error())
		return err
	}

	return nil
}

// updates employee with put method
func updEmployee(c *fiber.Ctx) error {
	emp := domain.Employee{}
	err := c.BodyParser(&emp)
	if err != nil {
		c.Context().SetStatusCode(400)
		c.SendString("Could not update employee: " + err.Error())
		return err
	}

	err = messaging.UpdEmployee(emp)
	if err != nil {
		c.Context().SetStatusCode(400)
		c.SendString("Could not update employee: " + err.Error())
		return err
	}

	c.Context().SetStatusCode(200)
	return nil
}

// finds employee by id
func findEmployee(c *fiber.Ctx) error {
	var id int64
	c.Locals("id", id)
	emp, err := messaging.FindEmployee(id)
	if err != nil {
		c.Context().SetStatusCode(404)
		c.SendString("User not found")
		return err
	}

	err = c.JSON(emp)
	if err != nil {
		c.SendString("error marshalling")
		return err
	}

	return nil
}

// update employee info read from multipart form
func updEmployeeFormData(c *fiber.Ctx) error {
	var id int64
	c.Locals("id", id)
	cID, err := strconv.ParseInt(c.FormValue("companyId"), 10, 64)
	if err != nil {
		return err
	}

	emp := domain.Employee{
		ID: id,
		Name: c.FormValue("name"),
		SecondName: c.FormValue("secondName"),
		Surname: c.FormValue("surname"),
		HireDate: c.FormValue("hireDate"),
		Position: c.FormValue("position"),
		CompanyID: cID,
	}

	err = messaging.UpdEmployeeFormData(emp)
	if err != nil {
		return err
	}

	return nil
}

// app.Delete("/employee/:employeeID", delEmployee())
func delEmployee(c *fiber.Ctx) error {
	var id int64
	c.Locals("id", id)
	err := messaging.DelEmployee(id)
	if err != nil {
		return err
	}
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
