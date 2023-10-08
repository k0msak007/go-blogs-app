package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/k0msak007/blog/database"
	"github.com/k0msak007/blog/model"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog List",
	}

	db := database.DBCon
	records := []model.Blog{}

	db.Find(&records)
	context["blog_records"] = &records

	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a Blog",
	}

	record := new(model.Blog)
	if err := c.BodyParser(&record); err != nil {
		context["statusText"] = ""
		context["msg"] = "Something went wrong"
		log.Println("Error in parsing request.")
	}

	if err := database.DBCon.Create(record).Error; err != nil {
		context["statusText"] = ""
		context["msg"] = "Something went wrong"
		log.Println("Error in saving data")
	}

	context["msg"] = "Record is saved successfully"
	context["data"] = record

	return c.Status(fiber.StatusCreated).JSON(context)

	return nil
}

func BlogUpdate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Blog",
	}

	//http://localhost:8000/2

	id := c.Params("id")

	var record model.Blog

	database.DBCon.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")

		context["statusText"] = ""
		context["msg"] = "Record not Found."
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")

		context["msg"] = "Something went wrong."
		c.Status(400)
		return c.JSON(context)
	}

	result := database.DBCon.Save(record)

	if result.Error != nil {
		log.Println("Error in saving data.")

		context["msg"] = "Error in saving data."
		c.Status(400)
		return c.JSON(context)
	}

	context["msg"] = "Record updated successfully."
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

// Delete a Blog
func BlogDelete(c *fiber.Ctx) error {

	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBCon.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."

		return c.JSON(context)
	}

	result := database.DBCon.Delete(record)

	if result.Error != nil {
		context["msg"] = "Something went wrong."
		return c.JSON(context)
	}

	context["statusText"] = "Ok"
	context["msg"] = "Record deleted successfully."
	c.Status(200)
	return c.JSON(context)
}
