package githubissuetool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func UpdateIssue() error {
	owner := os.Args[2]
	repo := os.Args[3]
	accessToken := os.Args[4]
	issueNumber := os.Args[5]

	title := os.Args[6]
	body := os.Args[7]
	assignees := os.Args[8]
	milestone := os.Args[9]
	labels := os.Args[10]
	state := os.Args[11]

	var userNames []string
	if assignees != "[]" {
		err := json.Unmarshal([]byte(assignees), &userNames)
		if err != nil {
			return err
		}
	}

	var tags []string
	if labels != "[]" {
		err := json.Unmarshal([]byte(labels), &tags)
		if err != nil {
			return err
		}
	}

	milestoneInt := 0
	if milestone != "0" {
		var err error
		milestoneInt, err = strconv.Atoi(milestone)
		if err != nil {
			return err
		}
	}

	requestBody := make(map[string]interface{})

	if title != "" {
		requestBody["title"] = title
	}

	if body != "" {
		requestBody["body"] = body
	}

	if state != "" {
		requestBody["state"] = state
	}

	if len(userNames) > 0 {
		requestBody["assignees"] = userNames
	}

	if len(tags) > 0 {
		requestBody["labels"] = tags
	}

	if milestoneInt != 0 {
		requestBody["milestone"] = milestoneInt
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/" + "issues" + "/" + issueNumber
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to update issue: %s", resp.Status)
	}

	fmt.Println("Issue updated successfully!")
	return nil

}

// curl -X PATCH -H "Authorization: token YOUR_PERSONAL_ACCESS_TOKEN" \
//      -H "Content-Type: application/json" \
//      -d '{
//            "title": "Updated issue title",
//            "body": "Updated description",
//            "state": "closed"
//          }' \
//      "https://api.github.com/repos/{owner}/{repo}/issues/{issue_number}"
