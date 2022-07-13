package main

import (
	"os"
)

/*
   PROJECT: 'medical_web'
   ENV: 'dev'
   REGION: 'ap-northeast-2'
*/
type ProjectEnv struct {
	Project     string
	Environment string
	Region      string
	IsLocal     bool
}

var Env ProjectEnv

func InitEnv() error {
	project := os.Getenv("PROJECT")
	environment := os.Getenv("ENV")
	region := os.Getenv("REGION")
	Env.Project = project
	Env.Environment = environment
	Env.Region = region

	return nil
}
