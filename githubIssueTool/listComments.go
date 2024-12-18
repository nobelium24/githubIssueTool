package githubissuetool

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func ListComments() (*CommentsResponse, error) {
	owner := os.Args[2]
	repo := os.Args[3]
	accessToken := os.Args[4]
	issueNumber := os.Args[5]
	comments := os.Args[6]

	url := "https://api.github.com/repos/" + owner + "/" + repo + "/issues/" + issueNumber + "/" + comments

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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to update i: %s", resp.Status)
	}

	var commentsResp CommentsResponse
	err = json.NewDecoder(resp.Body).Decode(&commentsResp)
	if err != nil {
		return nil, err
	}

	return &commentsResp, nil

}

// curl -H "Authorization: token YOUR_PERSONAL_ACCESS_TOKEN" \
//      "https://api.github.com/repos/{owner}/{repo}/issues/{issue_number}/comments"
