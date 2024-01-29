package router

import (
	"testAPI/app/handlers"

	"github.com/gofiber/fiber/v2"
)

type RouterInterface interface {
	BoardRouter(fiber.Router)
}

type Router struct {
	boardHandler handlers.HandlerInterface
}

func NewRouter() (RouterInterface, error) {
	br, err := handlers.NewHandler()
	if err != nil {
		return nil, err
	}

	return &Router{
		boardHandler: br,
	}, nil
}

func (r *Router) BoardRouter(api fiber.Router) {

	api.Get("/", r.boardHandler.ReadAll)
	api.Get("/:id", r.boardHandler.ReadById)

	api.Post("/", r.boardHandler.Create)
}
