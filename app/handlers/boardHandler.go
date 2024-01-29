package handlers

import (
	"testAPI/app/dal"
)

type BoardHandlerInterface interface {
	HandlerInterface

	// Board(c *fiber.Ctx) error
	// BoardById(c *fiber.Ctx) error
	// AddBoard(c *fiber.Ctx) error
}

type BoardHandler struct {
	Handler
	ORM dal.ORMInterface
}

// func NewBoardHandler() (BoardHandlerInterface, error) {
// 	orm, err := dal.NewBoardORM()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &BoardHandler{
// 		ORM: orm,
// 	}, nil
// }
