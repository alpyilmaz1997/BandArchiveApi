// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample Music Archive server.",
    "title": "Music Archive",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "name": "Alp Eren Yılmaz",
      "url": "http://alperenyilmaz.com",
      "email": "alpyilmaz1997@gmail.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/v2",
  "paths": {
    "/bands": {
      "get": {
        "description": "return all bands",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Bands"
        ],
        "summary": "Find All Band",
        "operationId": "getBands",
        "responses": {
          "200": {
            "description": "successfull operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Band"
              }
            }
          },
          "404": {
            "description": "Bands cannot find"
          },
          "500": {
            "description": "Internal Server"
          }
        }
      }
    },
    "/bands/{id}": {
      "get": {
        "description": "Return band details by id",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Bands"
        ],
        "summary": "Find Band by id",
        "operationId": "getBandById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Id of band to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successfull operation",
            "schema": {
              "$ref": "#/definitions/Band"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Band Id cannot find"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/newBand": {
      "post": {
        "description": "New Band is added to the Database",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Bands"
        ],
        "summary": "Add a new Band to the Database",
        "operationId": "addBand",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/newBand"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/Band"
            }
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    }
  },
  "definitions": {
    "Band": {
      "type": "object",
      "properties": {
        "discography": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Oceanborn",
              "Wishmaster"
            ]
          }
        },
        "formed": {
          "type": "string",
          "pattern": "^[0-9]{4}$",
          "example": "1996"
        },
        "genre": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Symphonic Metal"
            ]
          }
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "example": 1
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Tarja Turunen",
              "Marco Hietala"
            ]
          }
        },
        "name": {
          "type": "string",
          "example": "Nightwish"
        },
        "status": {
          "type": "boolean",
          "example": true
        }
      }
    },
    "newBand": {
      "type": "object",
      "properties": {
        "discography": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Oceanborn",
              "Wishmaster"
            ]
          }
        },
        "formed": {
          "type": "string",
          "pattern": "^[0-9]{4}$",
          "example": "1996"
        },
        "genre": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Symphonic Metal"
            ]
          }
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Tarja Turunen",
              "Marco Hietala"
            ]
          }
        },
        "name": {
          "type": "string",
          "example": "Nightwish"
        },
        "status": {
          "type": "boolean",
          "example": true
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything about Bands",
      "name": "Bands"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample Music Archive server.",
    "title": "Music Archive",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "name": "Alp Eren Yılmaz",
      "url": "http://alperenyilmaz.com",
      "email": "alpyilmaz1997@gmail.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/v2",
  "paths": {
    "/bands": {
      "get": {
        "description": "return all bands",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Bands"
        ],
        "summary": "Find All Band",
        "operationId": "getBands",
        "responses": {
          "200": {
            "description": "successfull operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Band"
              }
            }
          },
          "404": {
            "description": "Bands cannot find"
          },
          "500": {
            "description": "Internal Server"
          }
        }
      }
    },
    "/bands/{id}": {
      "get": {
        "description": "Return band details by id",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Bands"
        ],
        "summary": "Find Band by id",
        "operationId": "getBandById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Id of band to return",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successfull operation",
            "schema": {
              "$ref": "#/definitions/Band"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Band Id cannot find"
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    },
    "/newBand": {
      "post": {
        "description": "New Band is added to the Database",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Bands"
        ],
        "summary": "Add a new Band to the Database",
        "operationId": "addBand",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/newBand"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/Band"
            }
          },
          "500": {
            "description": "Internal Error"
          }
        }
      }
    }
  },
  "definitions": {
    "Band": {
      "type": "object",
      "properties": {
        "discography": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Oceanborn",
              "Wishmaster"
            ]
          }
        },
        "formed": {
          "type": "string",
          "pattern": "^[0-9]{4}$",
          "example": "1996"
        },
        "genre": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Symphonic Metal"
            ]
          }
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "example": 1
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Tarja Turunen",
              "Marco Hietala"
            ]
          }
        },
        "name": {
          "type": "string",
          "example": "Nightwish"
        },
        "status": {
          "type": "boolean",
          "example": true
        }
      }
    },
    "newBand": {
      "type": "object",
      "properties": {
        "discography": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Oceanborn",
              "Wishmaster"
            ]
          }
        },
        "formed": {
          "type": "string",
          "pattern": "^[0-9]{4}$",
          "example": "1996"
        },
        "genre": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Symphonic Metal"
            ]
          }
        },
        "members": {
          "type": "array",
          "items": {
            "type": "string",
            "example": [
              "Tarja Turunen",
              "Marco Hietala"
            ]
          }
        },
        "name": {
          "type": "string",
          "example": "Nightwish"
        },
        "status": {
          "type": "boolean",
          "example": true
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything about Bands",
      "name": "Bands"
    }
  ]
}`))
}
