// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-06-26 16:13:31.204563394 +0200 CEST m=+0.051967123

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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/models": {
            "get": {
                "description": "Query all available room models",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "models"
                ],
                "summary": "Query models",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.RoomModel"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/models/{id}": {
            "get": {
                "description": "Query a single room model by id with containing sensors",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "models"
                ],
                "summary": "Query room model",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "RoomModel ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.RoomModel"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sensors": {
            "get": {
                "description": "Query all available sensors.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sensors"
                ],
                "summary": "Query sensors",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Sensor"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sensors/{id}": {
            "get": {
                "description": "Query a single sensor by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sensors"
                ],
                "summary": "Query sensor",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Sensor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Sensor"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates the mesh id and anomaly preferences of a single sensor.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sensors"
                ],
                "summary": "Update sensor preferences",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "SensorId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateSensor",
                        "name": "update_sensor",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.UpdateSensor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Sensor"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sensors/{id}/anomalies": {
            "get": {
                "description": "Query anomalies for a specific sensor",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sensors"
                ],
                "summary": "Query anomalies",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Sensor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Start Date",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End Date",
                        "name": "end_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Anomaly"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sensors/{id}/data": {
            "get": {
                "description": "Query data for a specific sensor",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sensors"
                ],
                "summary": "Query sensor data",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Sensor ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Data Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Include only every nth element [1-16]",
                        "name": "density",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Start Date",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End Date",
                        "name": "end_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Data"
                            }
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Anomaly": {
            "type": "object",
            "properties": {
                "end_data": {
                    "type": "Data"
                },
                "peak_data": {
                    "type": "Data"
                },
                "start_data": {
                    "type": "Data"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "api.UpdateSensor": {
            "type": "object",
            "properties": {
                "gradient_bound": {
                    "type": "number"
                },
                "lower_bound": {
                    "type": "number"
                },
                "mesh_id": {
                    "type": "string"
                },
                "upper_bound": {
                    "type": "number"
                }
            }
        },
        "model.Data": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "object",
                    "$ref": "#/definitions/model.Date"
                },
                "gradient": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "sensor_id": {
                    "type": "integer"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "model.Date": {
            "type": "object"
        },
        "model.RoomModel": {
            "type": "object",
            "properties": {
                "floors": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "sensors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Sensor"
                    }
                },
                "type": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "model.Sensor": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "gradient_bound": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "import_name": {
                    "type": "string"
                },
                "latest_data": {
                    "type": "object",
                    "$ref": "#/definitions/model.Data"
                },
                "lower_bound": {
                    "type": "number"
                },
                "measurement_unit": {
                    "type": "string"
                },
                "mesh_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "range": {
                    "type": "string"
                },
                "room_model_id": {
                    "type": "integer"
                },
                "upper_bound": {
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
	Version:     "0.1.8",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "vi-sense BIM API",
	Description: "This API provides information about 3D room models with associated sensors and their data.",
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
