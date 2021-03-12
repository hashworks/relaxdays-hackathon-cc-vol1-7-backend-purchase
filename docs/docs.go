// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Justin Kromlinger",
            "url": "https://hashworks.net",
            "email": "justin.kromlinger@stud.htwk-leipzig.de"
        },
        "license": {
            "name": "GNU Affero General Public License v3",
            "url": "https://gnu.org/licenses/agpl.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/purchase": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Purchase"
                ],
                "summary": "Returns all saved purchases",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Purchase"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Purchase"
                ],
                "summary": "Save a purchase",
                "parameters": [
                    {
                        "description": "Purchase to save",
                        "name": "purchase",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Purchase"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Invalid purchase",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/storage/": {
            "get": {
                "tags": [
                    "Storage"
                ],
                "summary": "Returns all storages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Storage"
                            }
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "Storage"
                ],
                "summary": "Save or update a storage",
                "parameters": [
                    {
                        "description": "Storage to save",
                        "name": "storage",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Storage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Invalid storage",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/storage/{name}": {
            "delete": {
                "tags": [
                    "Storage"
                ],
                "summary": "Delete a storage by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of storage to delete",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Purchase": {
            "type": "object",
            "properties": {
                "articleID": {
                    "type": "integer"
                },
                "lieferant": {
                    "type": "string"
                },
                "menge": {
                    "type": "integer"
                }
            }
        },
        "models.Storage": {
            "type": "object",
            "properties": {
                "articleID": {
                    "type": "integer"
                },
                "bestand": {
                    "type": "integer"
                },
                "name": {
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
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Backend Tasks",
	Description: "Solution for backend task of https://sites.google.com/relaxdays.de/hackathon-relaxdays/startseite",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}