package main

import "time"

const (
	pollingInterval time.Duration = time.Second * 10

	repoNameEnv          string = "GITHUB_REPOSITORY"
	runIDEnv             string = "GITHUB_RUN_ID"
	repoOwnerEnv         string = "GITHUB_REPOSITORY_OWNER"
	inputTokenEnv        string = "INPUT_SECRET"
	inputApproversEnv    string = "INPUT_APPROVERS"
	inputMinApprovalsEnv string = "INPUT_MIN_APPROVALS"
	inputIssueTitleEnv   string = "INPUT_ISSUE_TITLE"
)

var (
	approvedKeywords = []string{"approved", "approve", "lgtm", "yes"}
	deniedKeywords   = []string{"denied", "deny", "no"}
)
