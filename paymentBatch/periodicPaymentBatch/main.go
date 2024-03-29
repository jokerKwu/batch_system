package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	common "github.com/jokerkwu/backend_common/batch_common"
	"time"
)

var globalErr error

func handler(ctx context.Context, request events.CloudWatchEvent) error {
	if globalErr != nil {
		return globalErr
	}
	//배치 로직 구현
	//TODO NextPayment가 오늘인 결제에 대해서 결제 처리한다.
	KST, _ := time.LoadLocation("Asia/Seoul")
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, KST).UTC() // 자정
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 24, 0, 0, 0, KST).UTC()  // 자정 직전
	fmt.Println(startTime)
	fmt.Println(endTime)

	return nil
}

func main() {
	fmt.Println("main hello11")
	if err := common.InitEnv(); err != nil {
		globalErr = err
	}
	if err := common.InitAws(common.Env.Region); err != nil {
		globalErr = err
	}
	if err := InitMongoDB(); err != nil {
		globalErr = err
	}
	lambda.Start(handler)
}
