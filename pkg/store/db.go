package store

import (
	"log"
	"os"

	"github.com/amila-ku/botbackend/pkg/entity"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var region = "eu-central-1"
var endpoint = "http://localhost:8000"

//Database defines a interface for handling persistance
type Database interface {
	NewTable()
	AddItem()
	DeleteItem()
	UpdateItem()
}

//
type DynamoDBDatabase struct {
	itemTable *dynamodb.DynamoDB
	tableName string
}

