// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Amirali Kalhor",
            "url": "http://www.swagger.io/support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/shop/add-item": {
            "post": {
                "description": "add item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shop"
                ],
                "summary": "add item",
                "parameters": [
                    {
                        "description": "item",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.AddItem"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer jwtToken",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/shop/get-item": {
            "get": {
                "description": "create item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shop"
                ],
                "summary": "register item",
                "parameters": [
                    {
                        "description": "item",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.RegisterCategory"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer jwtToken",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/shop/register-category": {
            "post": {
                "description": "create category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shop"
                ],
                "summary": "register category",
                "parameters": [
                    {
                        "description": "item",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.RegisterCategory"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer jwtToken",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/shop/register-item": {
            "post": {
                "description": "create item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shop"
                ],
                "summary": "register item",
                "parameters": [
                    {
                        "description": "item",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.RegisterItem"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer jwtToken",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "http.AddItem": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "http.RegisterCategory": {
            "type": "object",
            "properties": {
                "desc": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "http.RegisterItem": {
            "type": "object",
            "properties": {
                "categoryName": {
                    "type": "string"
                },
                "count": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "sex": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "shop Swagger API",
	Description:      "this is document of shop service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
