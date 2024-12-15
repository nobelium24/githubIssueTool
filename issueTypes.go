package githubissuetool

import "time"

type GitHubIssue struct {
	ID        int        `json:"id"`         // Unique identifier for the issue
	NodeID    string     `json:"node_id"`    // GraphQL node ID
	URL       string     `json:"url"`        // API URL for the issue
	HTMLURL   string     `json:"html_url"`   // Web URL for the issue
	Title     string     `json:"title"`      // Title of the issue
	Body      string     `json:"body"`       // Body/description of the issue
	State     string     `json:"state"`      // Current state: "open" or "closed"
	Locked    bool       `json:"locked"`     // Indicates if the issue is locked
	Comments  int        `json:"comments"`   // Number of comments on the issue
	CreatedAt time.Time  `json:"created_at"` // Timestamp of issue creation
	UpdatedAt time.Time  `json:"updated_at"` // Timestamp of the last update
	ClosedAt  *time.Time `json:"closed_at"`  // Timestamp when the issue was closed (nullable)
	Assignees []User     `json:"assignees"`  // List of assigned users
	Labels    []Label    `json:"labels"`     // List of labels assigned to the issue
	Milestone *Milestone `json:"milestone"`  // Associated milestone (nullable)
	User      User       `json:"user"`       // User who created the issue
}

type RequestBody struct {
	title     string
	body      string
	assignees []string
	milestone int
	labels    []string
}

// Supporting types

type User struct {
	Login     string `json:"login"`      // GitHub username
	ID        int    `json:"id"`         // Unique user ID
	AvatarURL string `json:"avatar_url"` // URL of the user's avatar
	HTMLURL   string `json:"html_url"`   // Web URL of the user's profile
}

type Label struct {
	ID          int    `json:"id"`          // Unique identifier for the label
	NodeID      string `json:"node_id"`     // GraphQL node ID
	Name        string `json:"name"`        // Name of the label
	Description string `json:"description"` // Description of the label
	Color       string `json:"color"`       // Hex color code for the label
	Default     bool   `json:"default"`     // Indicates if it is a default label
}

type Milestone struct {
	ID           int        `json:"id"`            // Unique identifier for the milestone
	NodeID       string     `json:"node_id"`       // GraphQL node ID
	Title        string     `json:"title"`         // Title of the milestone
	Description  string     `json:"description"`   // Description of the milestone
	State        string     `json:"state"`         // Current state: "open" or "closed"
	OpenIssues   int        `json:"open_issues"`   // Number of open issues in this milestone
	ClosedIssues int        `json:"closed_issues"` // Number of closed issues in this milestone
	CreatedAt    time.Time  `json:"created_at"`    // Timestamp of milestone creation
	UpdatedAt    time.Time  `json:"updated_at"`    // Timestamp of the last update
	DueOn        *time.Time `json:"due_on"`        // Due date (nullable)
}

type CommentRequestBody struct {
	Body string `json:"body"`
}

type Comment struct {
	ID        int    `json:"id"`         // Unique ID of the comment
	Body      string `json:"body"`       // The content of the comment
	User      User   `json:"user"`       // User who created the comment
	CreatedAt string `json:"created_at"` // Timestamp of creation
	UpdatedAt string `json:"updated_at"` // Timestamp of last update
	URL       string `json:"url"`        // API URL of the comment
	HTMLURL   string `json:"html_url"`   // HTML URL of the comment
}

type CommentsResponse []Comment
