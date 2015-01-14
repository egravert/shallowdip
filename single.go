package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"
)

type RepoAction func(string)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

// START OMIT
func main() {
	token := os.Getenv("TOKEN")
	org := os.Getenv("ORG")

	client := createClient(token)
	pageThroughRepos(client, org, cloneRepo)
}

func createClient(token string) *github.Client {
	trans := &oauth.Transport{Token: &oauth.Token{AccessToken: token}}
	return github.NewClient(trans.Client())
}

func pageThroughRepos(client *github.Client, org string, action RepoAction) {
	opts := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{Page: 1}}
	for opts.ListOptions.Page > 0 {
		repoList, resp, _ := client.Repositories.ListByOrg(org, opts)
		for _, repo := range repoList {
			action(*repo.SSHURL)
		}
		opts.ListOptions.Page = resp.NextPage
	}
}

func cloneRepo(repo string) {
	fmt.Println("starting clone:", repo)
	cmd := exec.Command("git", "clone", repo)
	if err := cmd.Run(); err != nil {
		fmt.Printf("The repo %s could not be cloned. Error: %v\n", repo, err)
	}
	fmt.Println("Completed", repo)
}

// END OMIT
