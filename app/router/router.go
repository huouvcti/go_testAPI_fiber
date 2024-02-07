package router

import (
	"fmt"
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

/*
 * gName - 그룹명 string
 * group - 그룹된 Router 객체
 * model - 사용할 DB 테이블 객체
 */

func (r *Router) RunRouter(gName string, group fiber.Router, model dal.ModelInterface) {
	h, err := handlers.NewHandler(model)

	if err != nil {
		log.Println(err)
	}

	group.Get("/", h.ReadAll).Name(fmt.Sprint("Read ALL ", model.TableName()))
	group.Get("/:id", h.ReadById).Name("ReadById")

	group.Post("/", h.Create).Name("Create")
	group.Put("/:id", h.UpdateById).Name("UpdateById")

	// 특정 그룹 적용
	if gName == "board" {
		group.Delete("/:id", h.DeleteById).Name("DeleteById")
	}

}
