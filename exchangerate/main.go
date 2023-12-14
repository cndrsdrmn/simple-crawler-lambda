package main

import (
	"exchangerate/app"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if err := app.RunScraper(); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error running scraper with: %v", err.Error()),
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       "Success scraping data",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
