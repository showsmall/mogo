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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v2/base/instances": {
            "get": {
                "description": "gets all instances, databases, and table nested data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "base"
                ],
                "summary": "Gets all instance database and table data filtered by permissions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/view.RespInstanceSimple"
                            }
                        }
                    }
                }
            }
        },
        "/api/v2/pandas/nodes-results/{result-id}": {
            "patch": {
                "description": "only support excelProcess update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Updates the action on the execution result",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "result id",
                        "name": "result-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "excelProcess",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/pandas/nodes/{node-id}/crontab": {
            "post": {
                "description": "isRetry: 0 no 1 yes\nretryInterval: the unit is in seconds, 100 means 100s",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Creating a scheduled node scheduling task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "node id",
                        "name": "node-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "params",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.ReqCreateCrontab"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "isRetry: 0 no 1 yes\nretryInterval: the unit is in seconds, 100 means 100s",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Updating a scheduled node scheduling task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "node id",
                        "name": "node-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "params",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.ReqUpdateCrontab"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/pandas/nodes/{node-id}/lock-acquire": {
            "post": {
                "description": "Force the file edit lock to be acquired",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Force the file edit lock to be acquired",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "node id",
                        "name": "node-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/pandas/nodes/{node-id}/results": {
            "get": {
                "description": "Obtain the node execution result record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pandas"
                ],
                "summary": "Obtain the node execution result record",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "node id",
                        "name": "node-id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "current",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "isExcludeCrontabResult",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.RespNodeResultList"
                        }
                    }
                }
            }
        },
        "/api/v2/storage": {
            "post": {
                "description": "Creating a log library",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "storage"
                ],
                "summary": "Creating a log library",
                "parameters": [
                    {
                        "type": "string",
                        "name": "brokers",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "consumers",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "databaseId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "days",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "desc",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "rawLogField",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Raw JSON data",
                        "name": "source",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "tableName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "timeField",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "topics",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "1 string 2 float",
                        "name": "typ",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v2/storage/mapping-json": {
            "post": {
                "description": "Kafka JSON field mapping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "storage"
                ],
                "summary": "Kafka JSON field mapping",
                "parameters": [
                    {
                        "type": "string",
                        "name": "data",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.MappingStruct"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "view.MappingStruct": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.MappingStructItem"
                    }
                }
            }
        },
        "view.MappingStructItem": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "view.ReqCreateCrontab": {
            "type": "object",
            "properties": {
                "args": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.ReqCrontabArg"
                    }
                },
                "cron": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "dutyUid": {
                    "type": "integer"
                },
                "isRetry": {
                    "type": "integer"
                },
                "retryInterval": {
                    "type": "integer"
                },
                "retryTimes": {
                    "type": "integer"
                },
                "typ": {
                    "type": "integer"
                }
            }
        },
        "view.ReqCrontabArg": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "val": {
                    "type": "string"
                }
            }
        },
        "view.ReqUpdateCrontab": {
            "type": "object",
            "properties": {
                "args": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.ReqCrontabArg"
                    }
                },
                "cron": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "dutyUid": {
                    "type": "integer"
                },
                "isRetry": {
                    "type": "integer"
                },
                "retryInterval": {
                    "type": "integer"
                },
                "retryTimes": {
                    "type": "integer"
                },
                "typ": {
                    "type": "integer"
                }
            }
        },
        "view.RespDatabaseSimple": {
            "type": "object",
            "properties": {
                "cluster": {
                    "type": "string"
                },
                "databaseName": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "iid": {
                    "type": "integer"
                },
                "isCreateByCV": {
                    "type": "integer"
                },
                "tables": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.RespTableSimple"
                    }
                }
            }
        },
        "view.RespInstanceSimple": {
            "type": "object",
            "properties": {
                "databases": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.RespDatabaseSimple"
                    }
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "instanceName": {
                    "type": "string"
                }
            }
        },
        "view.RespNodeResult": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "cost": {
                    "type": "integer"
                },
                "ctime": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nickname": {
                    "type": "string"
                },
                "nodeId": {
                    "type": "integer"
                },
                "result": {
                    "type": "string"
                },
                "uid": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "view.RespNodeResultList": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/view.RespNodeResult"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "view.RespTableSimple": {
            "type": "object",
            "properties": {
                "createType": {
                    "type": "integer"
                },
                "desc": {
                    "type": "string"
                },
                "did": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "tableName": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.4.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "ClickVisual API",
	Description:      "Defines interface prefixes in terms of module overrides：\n- base : the global basic readable information module\n- storage : the log module\n- alarm : the alarm module\n- pandas : the data analysis module\n- cmdb : the configuration module\n- sysop : the system management module",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}