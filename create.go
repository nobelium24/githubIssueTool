package githubissuetool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func CreateIssue() (*GitHubIssue, error) {
	//POST /repos/{owner}/{repo}/issues
	owner := os.Args[1]
	repo := os.Args[2]
	accessToken := os.Args[3]
	assignees := os.Args[4]
	milestone := os.Args[5]
	labels := os.Args[6]
	title := os.Args[7]
	body := os.Args[8]

	var userNames []string
	err := json.Unmarshal([]byte(assignees), &userNames)
	if err != nil {
		return nil, err
	}

	var tags []string
	err = json.Unmarshal([]byte(labels), &tags)
	if err != nil {
		return nil, err
	}

	milestoneInt, err := strconv.Atoi(milestone)
	if err != nil {
		return nil, err
	}

	requestBody := RequestBody{
		title:     title,
		body:      body,
		assignees: userNames,
		milestone: milestoneInt,
		labels:    tags,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := " https://api.github.com/repos/" + owner + "/" + repo + "/issues"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create issue: %s", resp.Status)
	}

	var issue GitHubIssue
	err = json.NewDecoder(resp.Body).Decode(&issue)
	if err != nil {
		return nil, err
	}

	return &issue, nil

}
