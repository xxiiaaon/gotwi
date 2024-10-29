package types_test

import (
	"testing"
	"time"

	"github.com/xxiiaaon/gotwi/fields"
	"github.com/xxiiaaon/gotwi/tweet/searchtweet/types"
	"github.com/stretchr/testify/assert"
)

func Test_SearchTweetsRecent_SetAccessToken(t *testing.T) {
	cases := []struct {
		name   string
		token  string
		expect string
	}{
		{
			name:   "normal",
			token:  "test-token",
			expect: "test-token",
		},
		{
			name:   "empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.ListRecentInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SearchTweetsRecent_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	startTime := time.Date(2021, 10, 24, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2021, 10, 24, 23, 59, 59, 59, time.UTC)

	cases := []struct {
		name   string
		params *types.ListRecentInput
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListRecentInput{
				Query: "from:testuser",
			},
			expect: endpointBase + "?query=from%3Atestuser",
		},
		{
			name: "with expansions",
			params: &types.ListRecentInput{
				Query:      "from:testuser",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&query=from%3Atestuser",
		},
		{
			name: "with media.fields",
			params: &types.ListRecentInput{
				Query:       "from:testuser",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?media.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with place.fields",
			params: &types.ListRecentInput{
				Query:       "from:testuser",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?place.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with poll.fields",
			params: &types.ListRecentInput{
				Query:      "from:testuser",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?poll.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with tweets.fields",
			params: &types.ListRecentInput{
				Query:       "from:testuser",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?query=from%3Atestuser&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.ListRecentInput{
				Query:      "from:testuser",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?query=from%3Atestuser&user.fields=uf1%2Cuf2",
		},
		{
			name: "with end_time",
			params: &types.ListRecentInput{
				Query:   "from:testuser",
				EndTime: &endTime,
			},
			expect: endpointBase + "?end_time=2021-10-24T23%3A59%3A59Z&query=from%3Atestuser",
		},
		{
			name: "with start_time",
			params: &types.ListRecentInput{
				Query:     "from:testuser",
				StartTime: &startTime,
			},
			expect: endpointBase + "?query=from%3Atestuser&start_time=2021-10-24T00%3A00%3A00Z",
		},
		{
			name: "with since_id",
			params: &types.ListRecentInput{
				Query:   "from:testuser",
				SinceID: "sid",
			},
			expect: endpointBase + "?query=from%3Atestuser&since_id=sid",
		},
		{
			name: "with until_id",
			params: &types.ListRecentInput{
				Query:   "from:testuser",
				UntilID: "uid",
			},
			expect: endpointBase + "?query=from%3Atestuser&until_id=uid",
		},
		{
			name: "with sort_order",
			params: &types.ListRecentInput{
				Query:     "from:testuser",
				SortOrder: types.ListSortOrderRecency,
			},
			expect: endpointBase + "?query=from%3Atestuser&sort_order=recency",
		},
		{
			name: "all query parameters",
			params: &types.ListRecentInput{
				Query:       "from:testuser",
				Expansions:  fields.ExpansionList{"ex"},
				MediaFields: fields.MediaFieldList{"mf"},
				PlaceFields: fields.PlaceFieldList{"plf"},
				PollFields:  fields.PollFieldList{"pof"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
				MaxResults:  10,
				NextToken:   "token",
				SortOrder:   types.ListSortOrderRelevancy,
			},
			expect: endpointBase + "?expansions=ex&max_results=10&media.fields=mf&next_token=token&place.fields=plf&poll.fields=pof&query=from%3Atestuser&sort_order=relevancy&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListRecentInput{
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_SearchTweetsRecent_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListRecentInput
	}{
		{
			name:   "empty params",
			params: &types.ListRecentInput{},
		},
		{
			name:   "some params",
			params: &types.ListRecentInput{Query: "from:testuser"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}

func Test_SearchTweetsAll_SetAccessToken(t *testing.T) {
	cases := []struct {
		name   string
		token  string
		expect string
	}{
		{
			name:   "normal",
			token:  "test-token",
			expect: "test-token",
		},
		{
			name:   "empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.ListAllInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SearchTweetsAll_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"
	startTime := time.Date(2021, 10, 24, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2021, 10, 24, 23, 59, 59, 59, time.UTC)

	cases := []struct {
		name   string
		params *types.ListAllInput
		expect string
	}{
		{
			name: "only required parameter",
			params: &types.ListAllInput{
				Query: "from:testuser",
			},
			expect: endpointBase + "?query=from%3Atestuser",
		},
		{
			name: "with expansions",
			params: &types.ListAllInput{
				Query:      "from:testuser",
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpointBase + "?expansions=ex1%2Cex2&query=from%3Atestuser",
		},
		{
			name: "with media.fields",
			params: &types.ListAllInput{
				Query:       "from:testuser",
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?media.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with place.fields",
			params: &types.ListAllInput{
				Query:       "from:testuser",
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?place.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with poll.fields",
			params: &types.ListAllInput{
				Query:      "from:testuser",
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?poll.fields=tf1%2Ctf2&query=from%3Atestuser",
		},
		{
			name: "with tweets.fields",
			params: &types.ListAllInput{
				Query:       "from:testuser",
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpointBase + "?query=from%3Atestuser&tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.ListAllInput{
				Query:      "from:testuser",
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpointBase + "?query=from%3Atestuser&user.fields=uf1%2Cuf2",
		},
		{
			name: "with end_time",
			params: &types.ListAllInput{
				Query:   "from:testuser",
				EndTime: &endTime,
			},
			expect: endpointBase + "?end_time=2021-10-24T23%3A59%3A59Z&query=from%3Atestuser",
		},
		{
			name: "with start_time",
			params: &types.ListAllInput{
				Query:     "from:testuser",
				StartTime: &startTime,
			},
			expect: endpointBase + "?query=from%3Atestuser&start_time=2021-10-24T00%3A00%3A00Z",
		},
		{
			name: "with since_id",
			params: &types.ListAllInput{
				Query:   "from:testuser",
				SinceID: "sid",
			},
			expect: endpointBase + "?query=from%3Atestuser&since_id=sid",
		},
		{
			name: "with until_id",
			params: &types.ListAllInput{
				Query:   "from:testuser",
				UntilID: "uid",
			},
			expect: endpointBase + "?query=from%3Atestuser&until_id=uid",
		},
		{
			name: "with sort_order",
			params: &types.ListAllInput{
				Query:     "from:testuser",
				SortOrder: types.ListSortOrderRecency,
			},
			expect: endpointBase + "?query=from%3Atestuser&sort_order=recency",
		},
		{
			name: "all query parameters",
			params: &types.ListAllInput{
				Query:       "from:testuser",
				Expansions:  fields.ExpansionList{"ex"},
				MediaFields: fields.MediaFieldList{"mf"},
				PlaceFields: fields.PlaceFieldList{"plf"},
				PollFields:  fields.PollFieldList{"pof"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
				MaxResults:  10,
				NextToken:   "token",
				SortOrder:   types.ListSortOrderRelevancy,
			},
			expect: endpointBase + "?expansions=ex&max_results=10&media.fields=mf&next_token=token&place.fields=plf&poll.fields=pof&query=from%3Atestuser&sort_order=relevancy&tweet.fields=tf&user.fields=uf",
		},
		{
			name: "has no required parameter",
			params: &types.ListAllInput{
				Expansions:  fields.ExpansionList{"ex"},
				UserFields:  fields.UserFieldList{"uf"},
				TweetFields: fields.TweetFieldList{"tf"},
			},
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_SearchTweetsAll_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListAllInput
	}{
		{
			name:   "empty params",
			params: &types.ListAllInput{},
		},
		{
			name:   "some params",
			params: &types.ListAllInput{Query: "from:testuser"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}
