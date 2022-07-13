package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var globalErr error

func handler(ctx context.Context, request events.CloudWatchEvent) error {
	if globalErr != nil {
		return globalErr
	}

	return nil
}

func main() {
	fmt.Println("main hello")

	lambda.Start(handler)
}
