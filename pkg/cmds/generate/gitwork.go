package generate

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

func gitwork(cmd *cobra.Command) error {
	fmt.Println("Hello World")
	fmt.Println("organization", cmd.Flag("organization").Value)
	fmt.Println("auth-keu", cmd.Flag("auth-key").Value)

/*	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cmd.Flag("auth-key").Value.String()},
	)
	tc := oauth2.NewClient(ctx, ts)*/
	client := github.NewClient(nil)
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), cmd.Flag("organization").Value.String(), opt)
	if _, ok := err.(*github.RateLimitError); ok {
		fmt.Println("hit rate limit")
	}
	if err!=nil{
		fmt.Println("error", err)
	}
	//fmt.Println("got the repos",repos)
	for i,val:= range repos{
		fmt.Printf("name-%d ---- %s\n",i,*val.Name)
	}
	return nil
}
