package handlers

import (
	"log"
	"testAPI/app/dal"

	"github.com/gofiber/fiber/v2"
)

type HandlerInterface interface {
	Create(c *fiber.Ctx) error

	ReadAll(c *fiber.Ctx) error
	ReadById(c *fiber.Ctx) error

	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type Handler struct {
	Model  dal.ModelInterface
	Models []dal.ModelInterface
	ORM    dal.ORMInterface
}

func NewHandler() (HandlerInterface, error) {
	orm, err := dal.NewORM()
	if err != nil {
		return nil, err
	}

	model := dal.Board{}

	return &Handler{
		ORM:    orm,
		Model:  &model,
		Models: []dal.ModelInterface{&model},
	}, nil
}

func (h *Handler) Create(c *fiber.Ctx) error {
	// req := &dal.Board{}

	// if err := c.BodyParser(req); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"msg": err,
	// 	})
	// }

	// if err := h.ORM.Create(req); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"msg": err,
	// 	})
	// }

	// fmt.Printf("%v\n", req)

	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"msg": req,
	// })

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddd",
	})
}

func (h *Handler) ReadAll(c *fiber.Ctx) error {
	data, err := h.ORM.ReadAll(h.Models)

	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": data,
	})
}

func (h *Handler) ReadById(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddddd",
	})
}

func (h *Handler) Update(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddd",
	})
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddd",
	})
}
