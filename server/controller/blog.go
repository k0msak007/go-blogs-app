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
	return nil
}

func BlogDelete(c *fiber.Ctx) error {
	return nil
}
