package adapters

import (

	  "github.com/aws/aws-sdk-go-v2/service/dynamodb"

    "context"

)

type DynamoAdapter struct {
	client DynamoClient
}

type DynamoClient interface {
  // Remove or update methods from client as needed 

	GetItem(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItem(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	

}

func NewDynamoAdapter(client DynamoClient) *DynamoAdapter {
  return &DynamoAdapter{
    client: client,
  }
}

