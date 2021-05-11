package cmd

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

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
		c.Status(400).SendString("Invalid input")
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
		c.Status(400).SendString("Could not update employee: " + err.Error())
		return err
	}

	err = messaging.UpdEmployee(emp)
	if err != nil {
		c.Status(400).SendString("Could not update employee: " + err.Error())
		return err
	}

	c.Status(200)
	return nil
}

// finds employee by id
func findEmployee(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(400)
		return err
	}

	emp, err := messaging.FindEmployee(id)
	if err != nil {
		c.Status(404).SendString("User not found")
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
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(400)
		return err
	}

	cID, err := strconv.ParseInt(c.FormValue("companyId"), 10, 64)
	if err != nil {
		return err
	}

	emp := domain.Employee{
		ID:         id,
		Name:       c.FormValue("name"),
		SecondName: c.FormValue("secondName"),
		Surname:    c.FormValue("surname"),
		HireDate:   c.FormValue("hireDate"),
		Position:   c.FormValue("position"),
		CompanyID:  cID,
	}

	err = messaging.UpdEmployeeFormData(emp)
	if err != nil {
		return err
	}

	return nil
}

// app.Delete("/employee/:employeeID", delEmployee())
func delEmployee(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(400)
		return err
	}

	err = messaging.DelEmployee(id)
	if err != nil {
		c.Status(400)
		return err
	}

	return nil
}

// app.Post("/company", addCompany())
func addCompany(c *fiber.Ctx) error {
	comp := domain.Company{}
	err := c.BodyParser(&comp)
	if err != nil {
		c.Status(400)
		return err
	}

	err = messaging.AddCompany(comp)
	if err != nil {
		c.Status(400)
		return err
	}

	c.JSON(comp)
	return nil
}

// app.Put("/company", updCompany())
func updCompany(c *fiber.Ctx) error {
	comp := domain.Company{}
	err := c.BodyParser(&comp)
	if err != nil {
		zap.L().Error(err.Error())
		c.Status(400).SendString(err.Error())
		return err
	}

	err = messaging.UpdCompany(comp)
	if err != nil {
		c.Status(400)
		return err
	}

	return nil
}

// app.Get("/company/:companyID", findCompany())
func findCompany(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(400)
		return err
	}
	if id < 1 {
		c.Status(400)
		return errors.New("company id can't be under 1")
	}

	comp, err := messaging.FindCompany(id)
	if err != nil {
		c.Status(400)
		return err
	}

	c.JSON(comp)
	return nil
}

// app.Post("/company/:companyID", updCompanyFormData())
func updCompanyFormData(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(500)
		return err
	}

	comp := domain.Company{
		ID: id,
		Name: c.FormValue("name"),
		LegalForm: c.FormValue("status"),
	}

	messaging.UpdCompany(comp)

	return nil
}

// app.Delete("/company/:companyID", delCompany())
func delCompany(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(500)
		return err
	}
	
	err = messaging.DelCompany(id) 
	if err != nil {
		c.Status(400)
	}

	return nil
}

// app.Get("/company/:companyID/employees", getEmployeeList())
func getEmployeeList(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(500)
		return err
	}
	emps, err := messaging.GetEmlpoyeeList(id)
	if err != nil {
		c.Status(400)
		return err
	}

	c.JSON(emps)
	return nil
}
