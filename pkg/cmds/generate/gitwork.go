package generate

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

type user struct {
	languages map[string]struct{}
	repo      map[string]struct{}
	email     string
	name      string
	loginId   string
}

func gitwork(cmd *cobra.Command) error {

	header := "login,\tname, \temail, \trepositories, \tlanguages\n"
	authKey := cmd.Flag("auth-key").Value.String()
	org := cmd.Flag("organization").Value.String()
	if authKey == "" || authKey == " " {
		return fmt.Errorf("invalid auth-key")
	}
	if org == "" || org == " " {
		return fmt.Errorf("invalid organization")
	}
	fmt.Printf(header)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: authKey},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), org, opt)
	if _, ok := err.(*github.RateLimitError); ok {
		fmt.Println("hit rate limit")
	}
	if err != nil {
		fmt.Println("error", err)
	}
	users := make(map[string]user)
	for _, repo := range repos {
		contributors, _, _ := client.Repositories.ListContributors(ctx, org, *repo.Name, &github.ListContributorsOptions{})
		for _, con := range contributors {
			login := con.Login
			userService, _, _ := client.Users.Get(ctx, *login)

			if us, ok := users[*login]; ok {
				if repo.Language != nil {
					us.languages[*repo.Language] = struct{}{}
				}
				if repo.Name != nil {
					us.repo[*repo.Name] = struct{}{}
				}
			} else {
				langMap := make(map[string]struct{})
				repoMap := make(map[string]struct{})
				if repo.Language != nil {
					langMap[*repo.Language] = struct{}{}
				}
				if repo.Name != nil {
					repoMap[*repo.Name] = struct{}{}
				}
				users[*login] = user{
					languages: langMap,
					repo:      repoMap,
					email:     userService.GetEmail(),
					name:      userService.GetName(),
					loginId:   userService.GetLogin(),
				}
			}
		}
	}
	for _, val := range users {

		var lan string
		if val.languages != nil {
			for k, _ := range val.languages {
				lan = lan + ", " + k
			}
		}
		var rep string
		if val.repo != nil {
			for k, _ := range val.repo {
				rep = rep + ", " + k
			}
		}
		fmt.Printf("%s\t;%s\t;%s\t; %s; %s\n", val.name, val.loginId, val.email, rep, lan)
	}
	return nil
}
