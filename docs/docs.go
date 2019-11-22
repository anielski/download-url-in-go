// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-10-31 01:15:01.003939575 +0100 CET m=+1.428139338

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/api/fetcher": {
            "get": {
                "description": "get all fetcher - not delete",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API"
                ],
                "summary": "Get All fetcher",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/renderings.Fetchers"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "insert or update fetcher",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API"
                ],
                "summary": "Save fetcher",
                "parameters": [
                    {
                        "description": "Transaction Data",
                        "name": "verification",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/bindings.Fetcher"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/renderings.Fetcher"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "413": {
                        "description": "Request Entity Too Large",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/fetcher/{id}": {
            "delete": {
                "description": "set fetcher as delete",
                "tags": [
                    "API"
                ],
                "summary": "Delete fetcher",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/renderings.Fetcher"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/fetcher/{id}/history": {
            "get": {
                "description": "get all fetcher - not delete",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "API"
                ],
                "summary": "Get All fetcher",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/renderings.History"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "bindings.Fetcher": {
            "type": "object",
            "required": [
                "interval",
                "url"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "interval": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "internal": {
                    "type": "error"
                },
                "message": {
                    "type": "object"
                }
            }
        },
        "renderings.Fetcher": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "renderings.Fetchers": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "interval": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "renderings.History": {
            "type": "object",
            "properties": {
                "response": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/renderings.Response"
                    }
                }
            }
        },
        "renderings.Response": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "duration": {
                    "type": "number"
                },
                "response": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
