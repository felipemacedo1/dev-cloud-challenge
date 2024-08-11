// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Felipe Macedo",
            "email": "felipealexandrej@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/alunos": {
            "get": {
                "description": "Obtém a lista de TODOS os alunos cadastrados no sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Alunos"
                ],
                "summary": "Retorna a lista de alunos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Aluno"
                            }
                        }
                    },
                    "500": {
                        "description": "Erro interno no servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Adiciona um novo aluno ao sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Alunos"
                ],
                "summary": "Cria um novo aluno",
                "parameters": [
                    {
                        "description": "Dados do Aluno",
                        "name": "aluno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Aluno"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Aluno"
                        }
                    },
                    "400": {
                        "description": "Dados do aluno inválidos",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Erro interno no servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/alunos/{id}": {
            "get": {
                "description": "Obtém os dados de um aluno específico pelo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Alunos"
                ],
                "summary": "Retorna um aluno pelo ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do Aluno",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Dados do Aluno",
                        "schema": {
                            "$ref": "#/definitions/models.Aluno"
                        }
                    },
                    "400": {
                        "description": "ID inválido",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Aluno não encontrado",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Erro interno no servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Atualiza as informações de um aluno específico pelo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Alunos"
                ],
                "summary": "Atualiza os dados de um aluno",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do Aluno",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dados do Aluno",
                        "name": "aluno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Aluno"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Aluno"
                        }
                    },
                    "400": {
                        "description": "Dados inválidos ou ID inválido",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Aluno não encontrado",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Erro interno no servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove um aluno específico pelo ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Alunos"
                ],
                "summary": "Deleta um aluno pelo ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID do Aluno",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Aluno não encontrado",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Erro interno no servidor",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Aluno": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "idade": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "nome_professor": {
                    "type": "string"
                },
                "nota_primeiro_semestre": {
                    "type": "number"
                },
                "nota_segundo_semestre": {
                    "type": "number"
                },
                "numero_sala": {
                    "type": "integer"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "dev-cloud-challenge.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "API de Gestão de Alunos",
	Description:      "Esta é a documentação da API de Gestão de Alunos.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
