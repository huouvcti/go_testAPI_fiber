package app

import (
	"testAPI/app/dal"
	"testAPI/app/router"
	"testAPI/config"
	"testAPI/docs"

	_ "testAPI/docs"

	"github.com/gofiber/contrib/swagger"
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

	config.DB.AutoMigrate(&dal.Board{})

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
	// api.Name("api.")

	v1 := api.Group("/v1")
	// v1.Name("v1.")

	boardGroup := v1.Group("/board")
	boardGroup.Name("board.")
	r.RunRouter("board", boardGroup, &dal.Board{})

	listGroup := v1.Group("/list")
	listGroup.Name("list.")
	r.RunRouter("list", listGroup, &dal.List{})

	// aaaGroup := v1.Group("/aaa")
	// aaaGroup.Name("aaa.")
	// r.RunRouter(aaaGroup, &dal.List{})

	dalArr := []dal.ModelInterface{}
	dalArr = append(dalArr, &dal.Board{})
	dalArr = append(dalArr, &dal.List{})

	docs.GenSwaggerSpec(app, dalArr, docs.Info{
		Title:       "TEST API",
		Version:     "1.0",
		Description: "Go Fiber REST API server TEST !",
	})

	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "",
		Title:    "Swagger API Docs",
	}
	app.Use(swagger.New(cfg))

	// data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	// fmt.Println(string(data))

	// var data []fiber.Route
	// data = app.GetRoutes(true)
	// for _, value := range data {
	// 	fmt.Printf("Method: %v, Name: %v, Path: %v, Params: %v\n", value.Method, value.Name, value.Path, value.Params)
	// }

	// fmt.Println(data)

	return app.Listen(":1234")
}
