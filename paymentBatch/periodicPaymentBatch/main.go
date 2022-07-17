package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	commonSsm "github.com/jokerkwu/backend_common/batch_common/aws/ssm"
	commonDB "github.com/jokerkwu/backend_common/batch_common/db"
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
	if err := commonSsm.InitAws(Env.Region); err != nil {
		globalErr = err
	}
	if err := commonDB.InitMongoDB(); err != nil {
		globalErr = err
	}
	lambda.Start(handler)
}
