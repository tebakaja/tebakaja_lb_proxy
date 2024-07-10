// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://swagger.io/terms/",
        "contact": {
            "name": "Si Mimin",
            "url": "https://www.tebakaja.com",
            "email": "tebakaja@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/crypto/lists": {
            "get": {
                "description": "Cryptocurrency List",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cryptocurrency"
                ],
                "summary": "Cryptocurrency List",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/crypto.ApiResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Cryptocurrency Prediction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cryptocurrency"
                ],
                "summary": "Cryptocurrency Prediction",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/crypto.PredictionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/crypto.PredictionResponse"
                        }
                    }
                }
            }
        },
        "/national-currency/lists": {
            "get": {
                "description": "National Currency List",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "National Currency"
                ],
                "summary": "National Currency List",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/national_currency.ApiResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "National Currency Prediction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "National Currency"
                ],
                "summary": "National Currency Prediction",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/national_currency.PredictionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/national_currency.PredictionResponse"
                        }
                    }
                }
            }
        },
        "/stock/lists": {
            "get": {
                "description": "Stock List",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stock"
                ],
                "summary": "Stock List",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/stock.ApiResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Stock Prediction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stock"
                ],
                "summary": "Stock Prediction",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/stock.PredictionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/stock.PredictionResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "crypto.ApiResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "crypto.PredictionRequest": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "days": {
                    "type": "integer"
                }
            }
        },
        "crypto.PredictionResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "currency": {
                            "type": "string"
                        },
                        "predictions": {
                            "type": "object",
                            "properties": {
                                "actuals": {
                                    "type": "array",
                                    "items": {
                                        "type": "object",
                                        "properties": {
                                            "date": {
                                                "type": "string"
                                            },
                                            "price": {
                                                "type": "number"
                                            }
                                        }
                                    }
                                },
                                "predictions": {
                                    "type": "array",
                                    "items": {
                                        "type": "object",
                                        "properties": {
                                            "date": {
                                                "type": "string"
                                            },
                                            "price": {
                                                "type": "number"
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "national_currency.ApiResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "national_currency.PredictionRequest": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "days": {
                    "type": "integer"
                }
            }
        },
        "national_currency.PredictionResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "currency": {
                            "type": "string"
                        },
                        "predictions": {
                            "type": "object",
                            "properties": {
                                "actuals": {
                                    "type": "array",
                                    "items": {
                                        "type": "object",
                                        "properties": {
                                            "date": {
                                                "type": "string"
                                            },
                                            "price": {
                                                "type": "number"
                                            }
                                        }
                                    }
                                },
                                "predictions": {
                                    "type": "array",
                                    "items": {
                                        "type": "object",
                                        "properties": {
                                            "date": {
                                                "type": "string"
                                            },
                                            "price": {
                                                "type": "number"
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "stock.ApiResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "stock.PredictionRequest": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "days": {
                    "type": "integer"
                }
            }
        },
        "stock.PredictionResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "currency": {
                            "type": "string"
                        },
                        "predictions": {
                            "type": "object",
                            "properties": {
                                "actuals": {
                                    "type": "array",
                                    "items": {
                                        "type": "object",
                                        "properties": {
                                            "date": {
                                                "type": "string"
                                            },
                                            "price": {
                                                "type": "number"
                                            }
                                        }
                                    }
                                },
                                "predictions": {
                                    "type": "array",
                                    "items": {
                                        "type": "object",
                                        "properties": {
                                            "date": {
                                                "type": "string"
                                            },
                                            "price": {
                                                "type": "number"
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "192.168.137.1:7860",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "TebakAja",
	Description:      "TebakAja REST API Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}