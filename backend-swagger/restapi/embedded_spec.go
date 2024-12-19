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
    "description": "API to manage users into a database with CRUD operations",
    "title": "User Management API",
    "version": "1.0.0"
  },
  "host": "localhost:8082",
  "basePath": "/",
  "paths": {
    "/users": {
      "get": {
        "summary": "Get a list of users",
        "responses": {
          "200": {
            "description": "list of users",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          },
          "500": {
            "description": "Internal error server"
          }
        }
      },
      "post": {
        "summary": "Add a user",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User added successfully"
          },
          "400": {
            "description": "Invalid requesta"
          },
          "500": {
            "description": "Internal error server"
          }
        }
      }
    },
    "/users/{id}": {
      "put": {
        "summary": "Update a user",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User updated successfully"
          },
          "400": {
            "description": "Invalid request"
          },
          "500": {
            "description": "Internal error server"
          }
        }
      },
      "delete": {
        "summary": "Delete a user",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Usuario eliminado exitosamente"
          },
          "400": {
            "description": "Solicitud inválida"
          },
          "500": {
            "description": "Error interno del servidor"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "BearerAuth": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "BearerAuth": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API to manage users into a database with CRUD operations",
    "title": "User Management API",
    "version": "1.0.0"
  },
  "host": "localhost:8082",
  "basePath": "/",
  "paths": {
    "/users": {
      "get": {
        "summary": "Get a list of users",
        "responses": {
          "200": {
            "description": "list of users",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          },
          "500": {
            "description": "Internal error server"
          }
        }
      },
      "post": {
        "summary": "Add a user",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "User added successfully"
          },
          "400": {
            "description": "Invalid requesta"
          },
          "500": {
            "description": "Internal error server"
          }
        }
      }
    },
    "/users/{id}": {
      "put": {
        "summary": "Update a user",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "User updated successfully"
          },
          "400": {
            "description": "Invalid request"
          },
          "500": {
            "description": "Internal error server"
          }
        }
      },
      "delete": {
        "summary": "Delete a user",
        "parameters": [
          {
            "type": "integer",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Usuario eliminado exitosamente"
          },
          "400": {
            "description": "Solicitud inválida"
          },
          "500": {
            "description": "Error interno del servidor"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "BearerAuth": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "BearerAuth": []
    }
  ]
}`))
}