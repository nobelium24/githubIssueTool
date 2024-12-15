package githubissuetool

import (
	"fmt"
	"net/http"
	"os"
)

func LockIssue() error {
	owner := os.Args[2]
	repo := os.Args[3]
	accessToken := os.Args[4]
	issueNumber := os.Args[5]
	issueRequest := os.Args[6]

	if issueRequest == "LOCK" {
		url := "https://api.github.com/repos/" + owner + "/" + repo + "/issues/" + issueNumber + "/lock"
		req, err := http.NewRequest("PUT", url, nil)
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
			return fmt.Errorf("Failed to lock issue: " + issueNumber)
		}
		fmt.Print("Issue locked successfully")
		return nil
	} else if issueRequest == "UNLOCK" {
		url := "https://api.github.com/repos/" + owner + "/" + repo + "/issues/" + issueNumber + "/lock"
		req, err := http.NewRequest("DELETE", url, nil)
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
			return fmt.Errorf("Failed to lock issue: " + issueNumber)
		}
		fmt.Print("Issue locked successfully")
		return nil
	}
	return fmt.Errorf("Invalid request type")

}

// curl -X PUT -H "Authorization: token YOUR_PERSONAL_ACCESS_TOKEN" \
//      -H "Content-Type: application/json" \
//      "https://api.github.com/repos/{owner}/{repo}/issues/{issue_number}/lock"

// 	 curl -X DELETE -H "Authorization: token YOUR_PERSONAL_ACCESS_TOKEN" \
//      "https://api.github.com/repos/{owner}/{repo}/issues/{issue_number}/lock"
