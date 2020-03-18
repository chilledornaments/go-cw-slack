# go-cw-slack

## AWS Resources

You will need:

- A lambda function

- An SNS topic that the lambda function is subscribed to

- A CloudWatch alarm that sends breach and OK events to your SNS topic

## CloudWatch

### Sample Breach Event

**Note:** This is the `EventRecord.Records[0].SNS.Message` content

```json
{
    "AlarmName": "test-cw-go-mem",
    "AlarmDescription": null,
    "AWSAccountId": "123456789",
    "NewStateValue": "ALARM",
    "NewStateReason": "Threshold Crossed: 2 out of the last 2 datapoints [14.0 (12/03/20 14:31:00), 14.0 (12/03/20 14:26:00)] were greater than the threshold (10.0) (minimum 2 datapoints for OK -> ALARM transition).",
    "StateChangeTime": "2020-03-12T14:36:48.821+0000",
    "Region": "US East (Ohio)",
    "OldStateValue": "INSUFFICIENT_DATA",
    "Trigger": {
        "MetricName": "MemoryUtilized",
        "Namespace": "ECS/ContainerInsights",
        "StatisticType": "Statistic",
        "Statistic": "AVERAGE",
        "Unit": null,
        "Dimensions": [
            {
                "value": "dev-twoheavy-svc",
                "name": "ServiceName"
            },
            {
                "value": "dev-twoheavy",
                "name": "ClusterName"
            }
        ],
        "Period": 300,
        "EvaluationPeriods": 2,
        "ComparisonOperator": "GreaterThanThreshold",
        "Threshold": 10,
        "TreatMissingData": "- TreatMissingData:                    missing",
        "EvaluateLowSampleCountPercentile": ""
    }
}
```

### Sample Recovery Event

```json
{
    "AlarmName": "test-cw-go-mem",
    "AlarmDescription": null,
    "AWSAccountId": "123456789",
    "NewStateValue": "OK",
    "NewStateReason": "Threshold Crossed: 2 out of the last 2 datapoints [14.0 (12/03/20 14:45:00), 14.0 (12/03/20 14:40:00)] were not greater than the threshold (20.0) (minimum 1 datapoint for ALARM -> OK transition).",
    "StateChangeTime": "2020-03-12T14:50:13.507+0000",
    "Region": "US East (Ohio)",
    "OldStateValue": "ALARM",
    "Trigger": {
        "MetricName": "MemoryUtilized",
        "Namespace": "ECS/ContainerInsights",
        "StatisticType": "Statistic",
        "Statistic": "AVERAGE",
        "Unit": null,
        "Dimensions": [
            {
                "value": "dev-twoheavy-svc",
                "name": "ServiceName"
            },
            {
                "value": "dev-twoheavy",
                "name": "ClusterName"
            }
        ],
        "Period": 300,
        "EvaluationPeriods": 2,
        "ComparisonOperator": "GreaterThanThreshold",
        "Threshold": 20,
        "TreatMissingData": "- TreatMissingData:                    missing",
        "EvaluateLowSampleCountPercentile": ""
    }
}
```

## Lambda

### ENV Vars

- `SLACK_CHANNEL`
- `SLACK_ICON`
- `SLACK_USERNAME`
- `SLACK_WEBHOOK`

### Using

You can retrieve this from `https://bhi-oss.s3.us-east-2.amazonaws.com/cw-go-slack/latest.zip`
