package main

import (
	"context"
	"fmt"

	"github.com/xxiiaaon/gotwi"
	"github.com/xxiiaaon/gotwi/fields"
	"github.com/xxiiaaon/gotwi/tweet/searchtweet"
	sttypes "github.com/xxiiaaon/gotwi/tweet/searchtweet/types"
	"github.com/xxiiaaon/gotwi/user/follow"
	ftypes "github.com/xxiiaaon/gotwi/user/follow/types"
)

type twitterUser struct {
	ID       string
	Name     string
	Username string
}

func (f twitterUser) displayName() string {
	return fmt.Sprintf("%s@%s", f.Name, f.Username)
}

// onlyFollowsRecentActivity will output the accounts that are unilaterally following
// the specified user ID, along with up to three most recent tweets.
func onlyFollowsRecentActivity(c *gotwi.Client, userID string) {
	// list follows
	followings := map[string]twitterUser{}

	paginationToken := "init"
	for paginationToken != "" {
		p := &ftypes.ListFollowingsInput{
			ID:         userID,
			MaxResults: 1000,
		}

		if paginationToken != "init" && paginationToken != "" {
			p.PaginationToken = paginationToken
		}

		res, err := follow.ListFollowings(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		for _, u := range res.Data {
			followings[gotwi.StringValue(u.ID)] = twitterUser{
				ID:       gotwi.StringValue(u.ID),
				Name:     gotwi.StringValue(u.Name),
				Username: gotwi.StringValue(u.Username),
			}
		}

		if res.Meta.NextToken != nil {
			paginationToken = gotwi.StringValue(res.Meta.NextToken)
		} else {
			paginationToken = ""
		}
	}

	// list followers
	followers := map[string]twitterUser{}

	paginationToken = "init"
	for paginationToken != "" {
		p := &ftypes.ListFollowersInput{
			ID:         userID,
			MaxResults: 1000,
		}

		if paginationToken != "init" && paginationToken != "" {
			p.PaginationToken = paginationToken
		}

		res, err := follow.ListFollowers(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		for _, u := range res.Data {
			followers[gotwi.StringValue(u.ID)] = twitterUser{
				ID:       gotwi.StringValue(u.ID),
				Name:     gotwi.StringValue(u.Name),
				Username: gotwi.StringValue(u.Username),
			}
		}

		if res.Meta.NextToken != nil {
			paginationToken = gotwi.StringValue(res.Meta.NextToken)
		} else {
			paginationToken = ""
		}
	}

	// only following
	onlyFollowings := map[string]twitterUser{}
	for fid, u := range followings {
		if _, ok := followers[fid]; ok {
			continue
		}

		onlyFollowings[fid] = u
	}

	// get recent tweets
	for _, onlyFollow := range onlyFollowings {
		p := &sttypes.ListRecentInput{
			MaxResults:  10,
			Query:       "from:" + onlyFollow.Username + " -is:retweet -is:reply",
			TweetFields: fields.TweetFieldList{fields.TweetFieldCreatedAt},
		}
		res, err := searchtweet.ListRecent(context.Background(), c, p)
		if err != nil {
			panic(err)
		}

		fmt.Printf("----- %s's recent Tweets -----\n", onlyFollow.displayName())
		c := 0
		for _, t := range res.Data {
			if c > 3 {
				break
			}
			fmt.Printf("[%s] %s\n", t.CreatedAt, gotwi.StringValue(t.Text))
			c++
		}

		fmt.Println()
	}
}
