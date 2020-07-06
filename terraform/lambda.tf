resource "aws_lambda_function" "lambda" {
  function_name = format("%s-%s", var.environment, var.app_name)
  s3_bucket     = "bhi-oss"
  s3_key        = "cw-go-slack/latest.zip"
  role                           = aws_iam_role.role.arn
  handler                        = "main"
  memory_size                    = 128
  runtime                        = "go1.x"
  reserved_concurrent_executions = -1 // Disable concurrency limits
  timeout                        = 3
  environment {
    variables = {
      SLACK_CHANNEL  = "#channel-name"
      SLACK_ICON     = ":exclamation:"
      SLACK_USERNAME = "CloudWatch Bot"
      SLACK_WEBHOOK  = var.slack_webhook
    }
  }

  tags = {
    env      = var.environment
    app_name = var.app_name
  }
}