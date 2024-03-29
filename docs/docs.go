// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

// import "github.com/swaggo/swag"

// const docTemplate = `{
//     "schemes": {{ marshal .Schemes }},
//     "swagger": "2.0",
//     "info": {
//         "description": "{{escape .Description}}",
//         "title": "{{.Title}}",
//         "contact": {},
//         "version": "{{.Version}}"
//     },
//     "host": "{{.Host}}",
//     "basePath": "{{.BasePath}}",
//     "paths": {
//         "/users/{id}": {
//             "get": {
//                 "tags": [
//                     "Users",
//                     "Boards"
//                 ],
//                 "summary": "Get a user by ID",
//                 "responses": {
//                     "200": {
//                         "description": "OK"
//                     }
//                 }
//             }
//         }
//     }
// }`

// // SwaggerInfo holds exported Swagger Info so clients can modify it
// var SwaggerInfo = &swag.Spec{
// 	Version:          "",
// 	Host:             "",
// 	BasePath:         "",
// 	Schemes:          []string{},
// 	Title:            "",
// 	Description:      "",
// 	InfoInstanceName: "swagger",
// 	SwaggerTemplate:  docTemplate,
// 	LeftDelim:        "{{",
// 	RightDelim:       "}}",
// }

func init() {
	// swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
