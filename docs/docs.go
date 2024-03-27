// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://stayease.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.stayease.io/support",
            "email": "support@stayease.io"
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
        "/rooms": {
            "post": {
                "description": "Add a room to the listing of rooms",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rooms"
                ],
                "summary": "Add a room.",
                "responses": {
                    "201": {
                        "description": "Room added successfully",
                        "schema": {
                            "$ref": "#/definitions/model.Room"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Booking": {
            "type": "object",
            "properties": {
                "check_in_date": {
                    "description": "Date and time of check-in",
                    "type": "string"
                },
                "check_out_date": {
                    "description": "Date and time of check-out",
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "room_id": {
                    "description": "ID of the room being booked                                 // ID of the user making the booking",
                    "type": "integer"
                },
                "total_cost": {
                    "description": "Total cost of the booking",
                    "type": "number"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.Room": {
            "type": "object",
            "properties": {
                "booked_dates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Booking"
                    }
                },
                "capacity": {
                    "description": "Maximum number of occupants the room can accommodate",
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "description": {
                    "description": "Brief description of the room (features, amenities, etc.)",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_booked": {
                    "description": "Indicates whether the room is currently booked",
                    "type": "boolean"
                },
                "price": {
                    "description": "Price per night for booking the room",
                    "type": "number"
                },
                "room_number": {
                    "description": "A unique identifier for the room (e.g., room number or code)",
                    "type": "string"
                },
                "type": {
                    "description": "Type of room (e.g., Single, Double, Suite)",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "StayEase API",
	Description:      "StayEase is a comprehensive room management API that simplifies the process of booking and managing rooms. With StayEase, users can easily search for available rooms, make reservations, and perform various management tasks related to room bookings.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}