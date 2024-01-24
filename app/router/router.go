package router

import (
	"testAPI/app/handlers"

	"github.com/gofiber/fiber/v2"
)

type RouterInterface interface {
	BoardRouter(fiber.Router)
}

type Router struct {
	boardHandler handlers.BoardHandlerInterface
}

func NewRouter() (RouterInterface, error) {
	br, err := handlers.NewBoardHandler()
	if err != nil {
		return nil, err
	}

	return &Router{
		boardHandler: br,
	}, nil
}

func (r *Router) BoardRouter(api fiber.Router) {

	api.Get("/", r.boardHandler.Board)

	api.Get("/:id", r.boardHandler.BoardById)

	api.Post("/", r.boardHandler.AddBoard)
}
