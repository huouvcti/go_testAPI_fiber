package app

import (
	"testAPI/app/dal"
	"testAPI/app/router"
	"testAPI/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type API struct {
	Router *fiber.App
}

func RunAPI() error {
	config.OpenDB()

	r, err := router.NewRouter()
	if err != nil {
		return err
	}

	return RunAPIWithRouter(r)
}

func RunAPIWithRouter(r *router.Router) error {

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		TimeFormat:    "2006-01-02 15:04:05",
		DisableColors: false,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	/* RequestId 미들웨어 */
	// app.Use(requestid.New())
	// app.Use(func(c *fiber.Ctx) error {
	// 	requestID := c.GetRespHeader("X-Request-ID")
	// 	log.Printf("RequestID: %s - %s %s\n", requestID, c.Method(), c.OriginalURL())
	// 	return c.Next()
	// })

	api := app.Group("/api")
	v1 := api.Group("/v1")

	boardGroup := v1.Group("/board")
	r.RunRouter(boardGroup, &dal.Board{})

	listGroup := v1.Group("/list")
	r.RunRouter(listGroup, &dal.List{})

	return app.Listen(":1234")
}
