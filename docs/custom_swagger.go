package docs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

/*
 * swagger 문서 형식 struct
 */
type SwaggerSpec struct {
	OpenAPI string            `json:"openapi"`
	Info    Info              `json:"info"`
	Paths   map[string]Method `json:"paths"`
}

type Info struct {
	Title       string `json:"title"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

type Method map[string]any

type Operation struct {
	Tags      []string            `json:"tags"`
	Summary   string              `json:"summary"`
	Responses map[string]Response `json:"responses"`
}

type Operation2 struct {
	Tags       []string            `json:"tags"`
	Summary    string              `json:"summary"`
	Parameters []Parameter         `json:"parameters"`
	Responses  map[string]Response `json:"responses"`
}

type Parameter struct {
	Name     string `json:"name"`
	In       string `json:"in"`
	Required bool   `json:"required"`
	Schema   Schema `json:"schema"`
}

type Schema struct {
	Type string `json:"type"`
}

type Response struct {
	Description string `json:"description"`
}

/*
 * 데이터 정리 struct
 */
type Router struct {
	Tags   []string
	Method string
	Path   string
	Params []string
	Des    string
}

func PathConvert(str string) string {
	parts := strings.Split(str, "/")

	for i, part := range parts {
		if strings.Contains(part, ":") {
			params := strings.Split(part, ":")
			param := fmt.Sprint("{", params[1], "}")
			parts[i] = param
		}
	}
	joined := strings.Join(parts, "/")
	return joined
}

func GenSwaggerSpec(app *fiber.App, cfg ...Info) {
	info := Info{
		Title:       "undefined",
		Version:     "undefined",
		Description: "undefined",
	}
	if len(cfg) > 0 {
		info = cfg[0]
	}

	openapi := &SwaggerSpec{
		OpenAPI: "3.0.0",
		Info:    info,
		Paths:   make(map[string]Method),
	}

	// 전달 받은 값 정리
	router := []Router{}
	routes := app.GetRoutes(true)

	for _, route := range routes {
		if route.Name == "" {
			continue
		}

		if route.Method == "HEAD" {
			continue
		}

		if last := route.Path[len(route.Path)-1]; last == '/' {
			route.Path = route.Path[:len(route.Path)-1]
		}

		// fmt.Printf("Method: %v, Name: %v, Path: %v, Params: %v\n", route.Method, route.Name, route.Path, route.Params)

		routeName := strings.Split(route.Name, ".")

		var tags []string

		tags = append(tags, routeName[0])

		router = append(router, Router{
			Tags:   tags,
			Method: strings.ToLower(route.Method),
			Path:   PathConvert(route.Path),
			Params: route.Params,
			Des:    routeName[1],
		})

	}
	// fmt.Printf("%v\n", router)

	// openapi 값 채우기
	if len(cfg) > 0 {
		openapi.Info = cfg[0]
	}

	res := Response{
		Description: "Successful response",
	}

	resAll := make(map[string]Response)
	resAll["200"] = res

	for _, route := range router {
		var op any
		if len(route.Params) == 0 {
			op = Operation{
				Tags:      route.Tags,
				Summary:   route.Des,
				Responses: resAll,
			}
		} else {
			var parameters []Parameter
			parameters = append(parameters, Parameter{
				Name:     "id",
				In:       "path",
				Required: true,
				Schema: Schema{
					Type: "integer",
				},
			})

			op = Operation2{
				Tags:       route.Tags,
				Summary:    route.Des,
				Parameters: parameters,
				Responses:  resAll,
			}
		}

		if openapi.Paths[route.Path] == nil {
			openapi.Paths[route.Path] = Method{}
		}

		openapi.Paths[route.Path][route.Method] = op
	}

	// // JSON 인코딩
	// jsonData, err := json.Marshal(openapi)
	// if err != nil {
	// 	fmt.Println("JSON 인코딩 에러:", err)
	// 	return
	// }
	// fmt.Printf("\n\n\n%v\n", string(jsonData))

	// 파일로 저장
	file, err := os.Create("./docs/swagger.json")
	if err != nil {
		fmt.Println("파일 생성 오류:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(openapi)
	if err != nil {
		fmt.Println("JSON 인코딩 및 파일 쓰기 오류:", err)
		return
	}
}
