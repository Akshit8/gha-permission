package main

import (
	"context"
	"log"
	"os"
	"strconv"
)

func main() {
	repoName := os.Getenv(repoNameEnv)

	runID, err := strconv.Atoi(os.Getenv(runIDEnv))
	if err != nil {
		log.Fatalf("error getting runID: %v", err)
	}

	repoOwner := os.Getenv(repoOwnerEnv)

	ctx := context.Background()

}
