package main

import (
	"fmt"
	"os"
	"os/exec"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

func main() {
	token := os.Getenv("TOKEN")
	org := os.Getenv("ORG")

	client := createClient(token)

	repos := make(chan string)
	done := cloner(repos)
	pageThroughRepos(client, org, repos)
	<-done
}

// createClient create an http client that uses oauth for authorization
// and returns a github client configured to use that client
func createClient(token string) *github.Client {
	trans := &oauth.Transport{Token: &oauth.Token{AccessToken: token}}
	return github.NewClient(trans.Client())
}

// pageThroughRepos queries the github API for all of the repos and sends the repo's ssh url
// to the repos channel. The API uses paging, and requires multiple queries to pull back a full list
func pageThroughRepos(client *github.Client, org string, repos chan<- string) {
	opts := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{Page: 1}}
	for opts.ListOptions.Page > 0 {
		repoList, resp, _ := client.Repositories.ListByOrg(org, opts)
		for _, repo := range repoList {
			repos <- *repo.SSHURL
		}
		opts.ListOptions.Page = resp.NextPage
	}
	close(repos)
}

// Cloner setups up a goroutine that receives repo urls
// and clones them. When the repo channel is closed, it will
// close the done channel.
func cloner(repos chan string) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		for repo := range repos {
			cloneRepo(repo)
		}
		close(done)
	}()
	return done
}

// cloneRepo does the actual work of cloning a repo by
// shelling out to the os to run the git clone command.
func cloneRepo(repo string) {
	fmt.Println("starting clone:", repo)
	cmd := exec.Command("git", "clone", repo)
	if err := cmd.Run(); err != nil {
		fmt.Printf("The repo %s could not be cloned. Error: %v\n", repo, err)
	}
	fmt.Println("Completed", repo)
}
