package types_test

import (
	"testing"

	"github.com/xxiiaaon/gotwi"
	"github.com/xxiiaaon/gotwi/tweet/managetweet/types"
	"github.com/stretchr/testify/assert"
)

func Test_CreateOutput_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.CreateOutput
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.CreateOutput{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.CreateOutput{
				Data: struct {
					ID   *string "json:\"id\""
					Text *string "json:\"text\""
				}{
					ID:   gotwi.String("id"),
					Text: gotwi.String("text"),
				},
			},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			hpe := c.res.HasPartialError()
			assert.Equal(tt, c.expect, hpe)
		})
	}
}

func Test_DeleteOutput_HasPartialError(t *testing.T) {
	cases := []struct {
		name   string
		res    *types.DeleteOutput
		expect bool
	}{
		{
			name:   "initial struct",
			res:    &types.DeleteOutput{},
			expect: false,
		},
		{
			name: "has data",
			res: &types.DeleteOutput{
				Data: struct {
					Deleted *bool "json:\"deleted\""
				}{
					Deleted: gotwi.Bool(true),
				},
			},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			hpe := c.res.HasPartialError()
			assert.Equal(tt, c.expect, hpe)
		})
	}
}
