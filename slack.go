package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func sendSlack(m AlarmDetails) error {

	var c string
	var t string
	if m.NewStateValue == "ALARM" {
		c = "danger"
		t = fmt.Sprintf("CloudWatch Alarm %s in alarm state", m.Name)
	} else {
		c = "good"
		t = fmt.Sprintf("CloudWatch Alarm %s has recovered", m.Name)
	}

	var f []SlackFields

	for _, v := range m.Trigger.Dimensions {
		//x := SlackFields{Title: v.Name, Value: v.Value, Short: true}
		f = append(f, SlackFields{Title: v.Name, Value: v.Value, Short: true})
	}

	sm := SlackMessage{
		Text:     m.Name,
		Channel:  os.Getenv("SLACK_CHANNEL"),
		Icon:     os.Getenv("SLACK_ICON"),
		Username: os.Getenv("SLACK_USERNAME"),
		Attachments: []SlackAttachments{
			SlackAttachments{
				Text:   m.NewStateReason,
				Color:  c,
				Title:  t,
				Fields: f,
			},
		},
	}

	client := &http.Client{Timeout: 4 * time.Second}

	d, err := json.Marshal(sm)

	if err != nil {
		log.Printf("Hit error marshalling Slack JSON - %s", err.Error())
		return err
	}

	req, err := http.NewRequest("POST", os.Getenv("SLACK_WEBHOOK"), bytes.NewBuffer(d))

	if err != nil {
		log.Printf("Hit error creating http request - %s", err.Error())
		return err
	}

	resp, err := client.Do(req)

	rb, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Hit error making request to Slack - %s", err.Error())
		return err
	}

	if resp.StatusCode != 200 {
		log.Printf("Hit error sending Slack request - status code: %d - body: %s", resp.StatusCode, string(rb))
		return errors.New("Error sending Slack request")
	}

	return nil
}
