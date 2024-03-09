package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Config struct {
	Destination string
}

func main() {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// repository list option
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	repoList, _, _ := client.Repositories.List(ctx, "arfo90", opt)

	for _, repo := range repoList {
		fmt.Printf("Name: %v, URL: %v\n", repo.GetName(), repo.GetCloneURL())
	}

	repo, _, err := client.Repositories.Get(ctx, "arfo90", "chinook")
	if err != nil {
		fmt.Printf("error getting repo: %v\n", err)
		return
	}

	// clone the repo
	_, err = git.PlainClone("/Users/amir/development/arfo90/tmp", false, &git.CloneOptions{
		URL:      repo.GetCloneURL(),
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Printf("error cloning repo: %v\n", err)
		return
	}

	fmt.Print("Repository is clone at /tmp/foo\n")
}
