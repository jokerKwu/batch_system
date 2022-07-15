package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	AwsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"strings"
)

var awsClientSsm *ssm.Client

func InitAws() error {

	awsConfig, err := AwsConfig.LoadDefaultConfig(context.TODO(),
		AwsConfig.WithRegion(Env.Region))
	if err != nil {
		return err
	}
	awsClientSsm = ssm.NewFromConfig(awsConfig)

	return nil
}

func AwsGetParams(paths []string) ([]string, error) {
	ctx := context.TODO()
	// get ssm param
	params, err := awsClientSsm.GetParameters(ctx, &ssm.GetParametersInput{
		Names:          paths,
		WithDecryption: true,
	})
	if err != nil {
		return nil, err
	}
	result := make([]string, len(paths))
	for i, path := range paths {
		val := ""
		for _, parameter := range params.Parameters {
			if strings.Contains(aws.ToString(parameter.ARN), path) {
				val = aws.ToString(parameter.Value)
				break
			}
		}
		result[i] = val
	}
	return result, nil
}
