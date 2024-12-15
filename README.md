# GitHub Issue Tool

This CLI tool allows you to interact with GitHub issues. You can create, read, update, lock, and comment on issues in a GitHub repository.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/githubissuetool.git
    cd githubissuetool
    ```

2. Build the CLI tool:
    ```sh
    go build -o githubissuetool
    ```

## Usage

The CLI tool supports the following commands:

- `CREATE`: Create a new issue
- `GET`: List all issues
- `GET_ONE`: Get a specific issue
- `UPDATE`: Update an issue
- `LOCK`: Lock or unlock an issue
- `COMMENT`: Comment on an issue
- `LIST_COMMENTS`: List comments on an issue

### Create a New Issue

```sh
./githubissuetool CREATE <owner> <repo> <accessToken> <[assignee1, assignee2]> <milestone> <[label1, label2]> <title> <body>
```

### List All issues

```sh 
./githubissuetool GET <owner> <repo> <accessToken>
```

### Get a Specific Issue

```sh
./githubissuetool GET_ONE <owner> <repo> <accessToken> <issueNumber>
```

### Update an Issue

```sh
./githubissuetool UPDATE <owner> <repo> <accessToken> <issueNumber> <title> <body> <[assignee1, assignee2]> <milestone> <[label1, label2]> <state>
```

### Lock or Unlock an Issue

```sh
./githubissuetool LOCK <owner> <repo> <accessToken> <issueNumber> <LOCK|UNLOCK>
```

### Comment on an Issue

```sh
./githubissuetool COMMENT <owner> <repo> <accessToken> <issueNumber> <body>
```

### List Comments on an Issue

```sh
./githubissuetool LIST_COMMENTS <owner> <repo> <accessToken> <issueNumber> <comments>
```

Each command requires specific arguments. Run the CLI tool without any arguments to see the required arguments for each command.
