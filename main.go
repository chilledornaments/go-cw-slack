package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func handler(r EventRecord) error {
	var details AlarmDetails

	eventRecordBytes := []byte(r.Records[0].SNS.Message)
	err := json.Unmarshal(eventRecordBytes, &details)

	if err != nil {
		log.Printf("Hit error unmarshalling alarm %s", err.Error())
		return err
	}

	log.Println(string(eventRecordBytes))

	log.Printf("New alarm %s - Reason %s - Value %s", details.Name, details.NewStateReason, details.NewStateValue)

	err = sendSlack(details)

	if err != nil {
		log.Printf("Error sending message to Slack %s", err.Error())
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
