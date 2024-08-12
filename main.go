package main

import (
	"lambda-cychore/api"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(api.Handler)
}
