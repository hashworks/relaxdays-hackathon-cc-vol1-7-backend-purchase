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
        "/articlesForLieferant": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Purchase"
                ],
                "summary": "Returns articleIDs by vendor",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Vendor",
                        "name": "x",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        },
        "/plot": {
            "get": {
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "Purchase"
                ],
                "summary": "Returns a plot of the price of an article over time",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Article ID",
                        "name": "x",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/purchase": {
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
        "/purchases": {
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Purchase"
                            }
                        }
                    }
                }
            }
        },
        "/purchasesBetween": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Purchase"
                ],
                "summary": "Returns all saved purchases between two points in time",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Starting point in time in the format 13.03.2021 13:59:58",
                        "name": "x",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ending point in time in the format 20.03.2021 15:59:58",
                        "name": "y",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Purchase"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid points in time",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/purchasesForArticle": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Purchase"
                ],
                "summary": "Returns all saved purchases for a given article",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of article to query",
                        "name": "x",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Purchase"
                            }
                        }
                    }
                }
            }
        },
        "/searchLieferant": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Purchase"
                ],
                "summary": "Returns a list of vendors similar to a query",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Vendor query",
                        "name": "x",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
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
                },
                "preis": {
                    "type": "number"
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
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Purchase Backend Task",
	Description: "Solution for 'Einkauf' backend task of https://sites.google.com/relaxdays.de/hackathon-relaxdays/startseite#h.klg8hathdmsn",
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
