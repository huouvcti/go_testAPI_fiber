package handlers

import "github.com/gofiber/fiber/v2"

type HandlerInterface interface {
	Create(c *fiber.Ctx) error

	ReadAll(c *fiber.Ctx) error
	ReadById(c *fiber.Ctx) error

	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type Handler struct {
	ORM interface{}
}

func (h Handler) Create(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddd",
	})
}

func (h Handler) ReadAll(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddd",
	})
}

func (h Handler) ReadById(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddd",
	})
}

func (h Handler) Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddd",
	})
}

func (h Handler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddd",
	})
}
