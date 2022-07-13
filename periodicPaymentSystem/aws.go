package main

import (
	"context"
	"fmt"
	AwsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var awsClientSsm *ssm.Client

func InitAws() error {

	awsConfig, err := AwsConfig.LoadDefaultConfig(context.TODO(),
		AwsConfig.WithRegion(Env.Region))
	if err != nil {
		fmt.Println("athena client init ")
		return err
	}
	awsClientSsm = ssm.NewFromConfig(awsConfig)

	return nil
}
