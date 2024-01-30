package handlers

import (
	"log"
	"strconv"
	"testAPI/app/dal"

	"github.com/gofiber/fiber/v2"
)

type HandlerInterface interface {
	Create(c *fiber.Ctx) error

	ReadAll(c *fiber.Ctx) error
	ReadById(c *fiber.Ctx) error

	UpdateById(c *fiber.Ctx) error
	DeleteById(c *fiber.Ctx) error
}

type Handler struct {
	Model  dal.ModelInterface
	Models []dal.ModelInterface
	ORM    dal.ORMInterface
}

func NewHandler(model dal.ModelInterface) (HandlerInterface, error) {
	orm, err := dal.NewORM()
	if err != nil {
		return nil, err
	}

	return &Handler{
		ORM:    orm,
		Model:  model,
		Models: []dal.ModelInterface{model},
	}, nil
}

func (h *Handler) Create(c *fiber.Ctx) error {
	req := h.Model.GetModel()

	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	if err := h.ORM.Create(req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "OK",
	})
}

func (h *Handler) ReadAll(c *fiber.Ctx) error {
	data, err := h.ORM.ReadAll(h.Model)

	if err != nil {
		log.Print("[handler] ReadAll: ", err)
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

func (h *Handler) UpdateById(c *fiber.Ctx) error {
	req := h.Model.GetModel()

	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	req.SetID(c.Params("id"))

	if err := h.ORM.Update(req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "OK",
	})
}

func (h *Handler) DeleteById(c *fiber.Ctx) error {

	reqId := c.Params("id")

	uint64_id, err := strconv.ParseUint(reqId, 10, 0)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	id := uint(uint64_id)

	if err := h.ORM.Delete(h.Model, id); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "OK",
	})
}
