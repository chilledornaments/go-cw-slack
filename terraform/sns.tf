resource "aws_sns_topic" "slack_topic" {
  name   = format("%s-%s-topic", var.environment, var.app_name)
  policy = <<EOF
{
  "Version": "2008-10-17",
  "Id": "__default_policy_ID",
  "Statement": [
    {
      "Sid": "CloudWatchPublish",
      "Effect": "Allow",
      "Principal": {
        "Service": "cloudwatch.amazonaws.com"
      },
      "Action": "SNS:Publish",
      "Resource": "arn:aws:sns:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${format("%s-%s-topic", var.environment, var.app_name)}"
    }
  ]
}
EOF
}

resource "aws_sns_topic_subscription" "lambda_subscription" {
  topic_arn = aws_sns_topic.slack_topic.arn
  protocol = "lambda"
  endpoint_auto_confirms = true
  endpoint = aws_lambda_function.lambda.arn
}