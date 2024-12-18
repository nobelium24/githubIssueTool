package main

import (
	githubissuetool "GithubIssueTool/githubIssueTool"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	requestType := os.Args[1]
	switch requestType {
	case "CREATE":
		if len(os.Args) != 10 {
			fmt.Println("Usage: ./githubissuetool CREATE <owner> <repo> <accessToken> <[assignee1, assignee2]> <milestone> <[label1, label2]> <title> <body>")
			return
		}
		githubissuetool.CreateIssue()
	case "GET":
		if len(os.Args) != 4 {
			fmt.Println("Usage: ./githubissuetool GET <owner> <repo> <accessToken>")
			return
		}
		githubissuetool.GetIssues()
	case "GET_ONE":
		if len(os.Args) != 5 {
			fmt.Println("Usage: ./githubissuetool GET_ONE <owner> <repo> <accessToken> <issueNumber>")
			return
		}
		githubissuetool.GetIssue()
	case "UPDATE":
		if len(os.Args) != 11 {
			fmt.Println("Usage: ./githubissuetool UPDATE <owner> <repo> <accessToken> <issueNumber> <title> <body> <[assignee1, assignee2]> <milestone> <[label1, label2]> <state>")
			return
		}
		githubissuetool.UpdateIssue()
	case "LOCK":
		if len(os.Args) != 6 {
			fmt.Println("Usage: ./githubissuetool LOCK <owner> <repo> <accessToken> <issueNumber> <LOCK|UNLOCK>")
			return
		}
		githubissuetool.LockIssue()
	case "COMMENT":
		if len(os.Args) != 6 {
			fmt.Println("Usage: ./githubissuetool COMMENT <owner> <repo> <accessToken> <issueNumber> <body>")
			return
		}
		githubissuetool.CommentIssue()
	case "LIST_COMMENTS":
		if len(os.Args) != 5 {
			fmt.Println("Usage: ./githubissuetool LIST_COMMENTS <owner> <repo> <accessToken> <issueNumber> <comments>")
			return
		}
		githubissuetool.ListComments()
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  ./githubissuetool <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  CREATE         Create a new issue")
	fmt.Println("  GET            List all issues")
	fmt.Println("  GET_ONE        Get a specific issue")
	fmt.Println("  UPDATE         Update an issue")
	fmt.Println("  LOCK           Lock or unlock an issue")
	fmt.Println("  COMMENT        Comment on an issue")
	fmt.Println("  LIST_COMMENTS  List comments on an issue")
	fmt.Println("Run './githubissuetool <command>' without arguments to see the required arguments for each command.")
}
