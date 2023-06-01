package View

import "github.com/swaggo/swag"

const docTemplate = `{
  "openapi": "3.0.3",
  "info": {
    "title": "Golang MVC Structure",
    "version": "1.0.11"
  },
  "tags": [
    {
      "name": "Fruit"
    }
  ],
  "paths": {
    "/getAllFruits": {
      "get": {
        "tags": [
          "Fruit"
        ],
        "summary": "get Fruit",
        "responses": {
          "201": {
            "description": "get Successful",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/Fruit"
                }
              }
            }
          }
        }
      }
    }
  },
  "components": {
    "schemas": {
		"Fruit": {
			"type": "object",
			"properties": {
			"f_id": {
				"type": "integer",
				"example": "123"
			}
			}
     	}
    }
  }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Golang MVC",
	Description:      "This is a sample server todo server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
