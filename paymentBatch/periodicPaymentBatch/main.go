package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	common "github.com/jokerkwu/backend_common/batch_common"
)

var globalErr error

func handler(ctx context.Context, request events.CloudWatchEvent) error {
	if globalErr != nil {
		return globalErr
	}
	fmt.Println("핸들러 탄다.")

	return nil
}

func main() {
	fmt.Println("main hello")
	if err := InitEnv(); err != nil {
		globalErr = err
	}
	if err := common.InitAws(Env.Region); err != nil {
		globalErr = err
	}
	if err := common.InitMongoDB(); err != nil {
		globalErr = err
	}
	lambda.Start(handler)
}
