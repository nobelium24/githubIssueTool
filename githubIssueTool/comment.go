package githubissuetool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func CommentIssue() error {
	owner := os.Args[2]
	repo := os.Args[2]
	accessToken := os.Args[4]
	issueNumber := os.Args[5]
	body := os.Args[6]

	url := "https://api.github.com/repos/" + owner + repo + "/issues/" + issueNumber + "/comments"

	reqBody := CommentRequestBody{
		Body: body,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("Post", url, bytes.NewBuffer(jsonData))
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
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create issue: %s", resp.Status)
	}

	fmt.Println("Issue updated successfully!")
	return nil
}

// curl -X POST -H "Authorization: token YOUR_PERSONAL_ACCESS_TOKEN" \
//      -H "Content-Type: application/json" \
//      -d '{
//            "body": "This is a comment on the issue."
//          }' \
//      "https://api.github.com/repos/{owner}/{repo}/issues/{issue_number}/comments"
