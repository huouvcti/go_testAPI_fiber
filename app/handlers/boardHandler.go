package handlers

import (
	"fmt"
	"testAPI/app/dal"

	"github.com/gofiber/fiber/v2"
)

type BoardHandlerInterface interface {
	HandlerInterface

	// Board(c *fiber.Ctx) error
	// BoardById(c *fiber.Ctx) error
	// AddBoard(c *fiber.Ctx) error
}

type BoardHandler struct {
	Handler
	ORM dal.BoardORMInterface
}

func NewBoardHandler() (BoardHandlerInterface, error) {
	orm, err := dal.NewBoardORM()
	if err != nil {
		return nil, err
	}

	return &BoardHandler{
		ORM: orm,
	}, nil
}

func (h *BoardHandler) ReadAll(c *fiber.Ctx) error {
	data, err := h.ORM.GetBoard()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": data,
	})
}

func (h *BoardHandler) ReadById(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddddd",
	})
}

func (h *BoardHandler) Create(c *fiber.Ctx) error {
	req := &dal.Board{}

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err,
		})
	}

	if err := h.ORM.AddBoard(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": err,
		})
	}

	fmt.Printf("%v\n", req)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": req,
	})
}

func (h *BoardHandler) Update(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddd",
	})
}

func (h *BoardHandler) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg": "ddd",
	})
}
