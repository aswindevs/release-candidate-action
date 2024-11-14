package rc

import (
	"encoding/json"
	"fmt"
	"strings"
)


func SlackPayloadBuilder(rcVersion string,prList []map[string]interface{} ) (slackPayload string,err error) {

	var prDetails strings.Builder
	for _, pr := range prList {

		if pr["url"] != "" {
		prDetails.WriteString(fmt.Sprintf("• *%s:* <%s|PR Link> - %s\n", pr["repo"], pr["url"], pr["error"]))
	
		} else {
			prDetails.WriteString(fmt.Sprintf("• *%s:* %s\n", pr["repo"], pr["error"]))
		}
	}

	// Constructing the Slack message payload
	payload := map[string]interface{}{
		"blocks": []interface{}{
			map[string]interface{}{
				"type": "header",
				"text": map[string]string{
					"type": "plain_text",
					"text": fmt.Sprintf("🚀 Release Candidate Branches for %s", rcVersion),
				},
			},
			map[string]interface{}{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": "Below is a compact list of RC branch PR details for review. 📋",
				},
			},
			map[string]interface{}{
				"type": "divider",
			},
			map[string]interface{}{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": fmt.Sprintf("*PRs by Repository:* \n\n%s", prDetails.String()),
				},
			},
			map[string]interface{}{
				"type": "divider",
			},
			map[string]interface{}{
				"type": "context",
				"elements": []interface{}{
					map[string]string{
						"type": "mrkdwn",
						"text": ":infinity: Generated by the *ReleaseWave*.",
					},
					map[string]string{
						"type": "mrkdwn",
						"text": ":rocket: *ReleaseWave* platform is under development.",
					},
				},
			},
		},
	}

	// Convert payload to JSON to send to Slack
	payloadJSON, _ := json.MarshalIndent(payload, "", "  ")
	slackPayload = string(payloadJSON)

	return slackPayload, nil
}