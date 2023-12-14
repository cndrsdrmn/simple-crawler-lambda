# Simple Crawler Lambda

## Overview

This project is a web scraping application built with Go and the Colly library. It uses AWS Serverless Application Model (AWS SAM) for deployment. The application is designed to crawl and scrape data from specific websites.

## Prerequisites

- [Go](https://go.dev/doc/install) installed on your local machine.
- [AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html) CLI installed.
- [Docker](https://www.docker.com/get-started) installed.

## Available Spiders

- [Bank Indonesia](https://www.bi.go.id/en/statistik/informasi-kurs/transaksi-bi/Default.aspx)
- [Ortax](https://ortax.org/ortax/?mod=kurs)

## Getting Started

1. Clone the repository:

   ```shell
   git clone git@github.com:cndrsdrmn/simple-crawler-lambda.git
   cd simple-crawler-lambda
   ```

2. Copy `env.example.json` to `env.json` and please make sure the environment variable of `DB_` same as your local settings.
3. Build and running invoke function:

   ```shell
   make && make invoke
   ```

4. Running on HTTP:

   ```shell
   make http && curl -X POST http://localhost:3000/crawler
   ```

## Packaging and deployment

AWS Lambda Golang runtime requires a flat folder with the executable generated on build step. SAM will use `CodeUri` property to know where to look up for the application:

```yaml
...
    FirstFunction:
        Type: AWS::Serverless::Function
        Properties:
            CodeUri: hello_world/
            ...
```

To deploy your application for the first time, run the following in your shell:

```bash
sam deploy --guided
```

The command will package and deploy your application to AWS, with a series of prompts:

- **Stack Name**: The name of the stack to deploy to CloudFormation. This should be unique to your account and region, and a good starting point would be something matching your project name.
- **AWS Region**: The AWS region you want to deploy your app to.
- **Confirm changes before deploy**: If set to yes, any change sets will be shown to you before execution for manual review. If set to no, the AWS SAM CLI will automatically deploy application changes.
- **Allow SAM CLI IAM role creation**: Many AWS SAM templates, including this example, create AWS IAM roles required for the AWS Lambda function(s) included to access AWS services. By default, these are scoped down to minimum required permissions. To deploy an AWS CloudFormation stack which creates or modifies IAM roles, the `CAPABILITY_IAM` value for `capabilities` must be provided. If permission isn't provided through this prompt, to deploy this example you must explicitly pass `--capabilities CAPABILITY_IAM` to the `sam deploy` command.
- **Save arguments to samconfig.toml**: If set to yes, your choices will be saved to a configuration file inside the project, so that in the future you can just re-run `sam deploy` without parameters to deploy changes to your application.

You can find your API Gateway Endpoint URL in the output values displayed after deployment.

## Contributing

Feel free to contribute to this project by submitting issues or pull requests. Your feedback and contributions are highly appreciated.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
