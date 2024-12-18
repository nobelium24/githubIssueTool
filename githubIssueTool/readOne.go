package githubissuetool

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetIssue() (*GitHubIssue, error) {
	//https://api.github.com/repos/{owner}/{repo}/issues/{issue_number}
	owner := os.Args[2]
	repo := os.Args[3]
	accessToken := os.Args[4]
	issueNumber := os.Args[5]

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/issues/" + issueNumber
	req, err := http.NewRequest("GET", url, nil)
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create issue: %s", resp.Status)
	}

	var issue GitHubIssue
	err = json.NewDecoder(resp.Body).Decode(&issue)
	if err != nil {
		return nil, err
	}

	return &issue, nil
}
