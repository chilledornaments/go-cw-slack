package main

/*
SNS Message documentation: https://docs.aws.amazon.com/lambda/latest/dg/with-sns.html

*/

// EventRecord is the entire message received from CloudWatch
type EventRecord struct {
	Records []struct {
		Source string `json:"EventSource"`
		SNS    struct {
			Message   string `json:"Message"`
			Type      string `json:"Type"`
			Subject   string `json:"Subject"`
			Timestamp string `json:"Timestamp"`
		}
	} `json:"Records"`
}

// AlarmDetails is the reason for the alarm
type AlarmDetails struct {
	Name           string       `json:"AlarmName"`
	NewStateValue  string       `json:"NewStateValue"`
	NewStateReason string       `json:"NewStateReason"`
	Trigger        AlarmTrigger `json:"Trigger"`
}

// AlarmTrigger contains details about the alarm that was breached
type AlarmTrigger struct {
	MetricName string                   `json:"MetricName"`
	Namespace  string                   `json:"Namespace"`
	Dimensions []AlarmTriggerDimensions `json:"Dimensions"`
}

// AlarmTriggerDimensions gives more details about the alarm
type AlarmTriggerDimensions struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// SlackMessage is the message we send to Slack
type SlackMessage struct {
	Text        string             `json:"text"`
	Channel     string             `json:"channel"`
	Icon        string             `json:"icon_emoji"`
	Username    string             `json:"username"`
	Attachments []SlackAttachments `json:"attachments"`
}

// SlackAttachments is the attachments we send in our message
type SlackAttachments struct {
	Text   string        `json:"text"`
	Color  string        `json:"color"`
	Title  string        `json:"title"`
	Fields []SlackFields `json:"fields"`
}

// SlackFields are additional fields in the attachment
type SlackFields struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}
