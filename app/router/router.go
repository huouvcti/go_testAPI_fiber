package router

import (
	"log"
	"testAPI/app/dal"
	"testAPI/app/handlers"

	"github.com/gofiber/fiber/v2"
)

type Router struct {
}

func NewRouter() (*Router, error) {
	return &Router{}, nil
}

func (r *Router) RunRouter(group fiber.Router, model dal.ModelInterface) {
	h, err := handlers.NewHandler(model)

	if err != nil {
		log.Println(err)
	}

	group.Get("/", h.ReadAll).Name("ReadAll")
	group.Get("/:id", h.ReadById).Name("ReadById")

	group.Post("/", h.Create).Name("Create")
	group.Put("/:id", h.UpdateById).Name("UpdateById")
	group.Delete("/:id", h.DeleteById).Name("DeleteById")
}
