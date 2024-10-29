package types_test

import (
	"testing"
	"time"

	"github.com/xxiiaaon/gotwi/tweet/tweetcount/types"
	"github.com/stretchr/testify/assert"
)

func Test_TweetCountsRecent_SetAccessToken(t *testing.T) {
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

func Test_TweetCountsRecent_ResolveEndpoint(t *testing.T) {
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
			expect: endpointBase + "?granularity=hour&query=from%3Atestuser",
		},
		{
			name: "with end_time",
			params: &types.ListRecentInput{
				Query:   "from:testuser",
				EndTime: &endTime,
			},
			expect: endpointBase + "?end_time=2021-10-24T23%3A59%3A59Z&granularity=hour&query=from%3Atestuser",
		},
		{
			name: "with start_time",
			params: &types.ListRecentInput{
				Query:     "from:testuser",
				StartTime: &startTime,
			},
			expect: endpointBase + "?granularity=hour&query=from%3Atestuser&start_time=2021-10-24T00%3A00%3A00Z",
		},
		{
			name: "with since_id",
			params: &types.ListRecentInput{
				Query:   "from:testuser",
				SinceID: "sid",
			},
			expect: endpointBase + "?granularity=hour&query=from%3Atestuser&since_id=sid",
		},
		{
			name: "with until_id",
			params: &types.ListRecentInput{
				Query:   "from:testuser",
				UntilID: "uid",
			},
			expect: endpointBase + "?granularity=hour&query=from%3Atestuser&until_id=uid",
		},
		{
			name: "with granularity",
			params: &types.ListRecentInput{
				Query:       "from:testuser",
				Granularity: types.TweetCountsGranularityDay,
			},
			expect: endpointBase + "?granularity=day&query=from%3Atestuser",
		},
		{
			name: "all query parameters",
			params: &types.ListRecentInput{
				Query:       "from:testuser",
				EndTime:     &endTime,
				StartTime:   &startTime,
				Granularity: types.TweetCountsGranularityMinute,
				SinceID:     "sid",
				UntilID:     "uid",
			},
			expect: endpointBase + "?end_time=2021-10-24T23%3A59%3A59Z&granularity=minute&query=from%3Atestuser&since_id=sid&start_time=2021-10-24T00%3A00%3A00Z&until_id=uid",
		},
		{
			name:   "has no required parameter",
			params: &types.ListRecentInput{},
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

func Test_TweetCountsRecent_Body(t *testing.T) {
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

func Test_TweetCountsAll_SetAccessToken(t *testing.T) {
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

func Test_TweetCountsAll_ResolveEndpoint(t *testing.T) {
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
			expect: endpointBase + "?granularity=hour&query=from%3Atestuser",
		},
		{
			name: "with end_time",
			params: &types.ListAllInput{
				Query:   "from:testuser",
				EndTime: &endTime,
			},
			expect: endpointBase + "?end_time=2021-10-24T23%3A59%3A59Z&granularity=hour&query=from%3Atestuser",
		},
		{
			name: "with start_time",
			params: &types.ListAllInput{
				Query:     "from:testuser",
				StartTime: &startTime,
			},
			expect: endpointBase + "?granularity=hour&query=from%3Atestuser&start_time=2021-10-24T00%3A00%3A00Z",
		},
		{
			name: "with since_id",
			params: &types.ListAllInput{
				Query:   "from:testuser",
				SinceID: "sid",
			},
			expect: endpointBase + "?granularity=hour&query=from%3Atestuser&since_id=sid",
		},
		{
			name: "with until_id",
			params: &types.ListAllInput{
				Query:   "from:testuser",
				UntilID: "uid",
			},
			expect: endpointBase + "?granularity=hour&query=from%3Atestuser&until_id=uid",
		},
		{
			name: "with granularity",
			params: &types.ListAllInput{
				Query:       "from:testuser",
				Granularity: types.TweetCountsGranularityDay,
			},
			expect: endpointBase + "?granularity=day&query=from%3Atestuser",
		},
		{
			name: "with next_token",
			params: &types.ListAllInput{
				Query:     "from:testuser",
				NextToken: "n_token",
			},
			expect: endpointBase + "?granularity=hour&next_token=n_token&query=from%3Atestuser",
		},
		{
			name: "all query parameters",
			params: &types.ListAllInput{
				Query:       "from:testuser",
				EndTime:     &endTime,
				StartTime:   &startTime,
				Granularity: types.TweetCountsGranularityMinute,
				SinceID:     "sid",
				UntilID:     "uid",
				NextToken:   "n_token",
			},
			expect: endpointBase + "?end_time=2021-10-24T23%3A59%3A59Z&granularity=minute&next_token=n_token&query=from%3Atestuser&since_id=sid&start_time=2021-10-24T00%3A00%3A00Z&until_id=uid",
		},
		{
			name:   "has no required parameter",
			params: &types.ListAllInput{},
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

func Test_TweetCountsAll_Body(t *testing.T) {
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
