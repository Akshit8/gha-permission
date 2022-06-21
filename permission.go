package main

import "github.com/google/go-github/v45/github"

type permission struct {
	client              *github.Client
	repoName            string
	repoOwner           string
	runID               int
	approvers           []string
	minimumApprovals    int
	approvalIssue       *github.Issue
	approvalIssueNumber int
	approvalIssueTitle  string
}
