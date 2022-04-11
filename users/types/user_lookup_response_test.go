package types_test

import (
	"testing"

	"github.com/michimani/gotwi/resources"
	"github.com/michimani/gotwi/users/types"
	"github.com/stretchr/testify/assert"
)

func Test_UserLookup_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListUsersOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListUsersOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListUsersOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListUsersOutput{
				Errors: []resources.PartialError{}},
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

func Test_UserLookupID_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.GetUserOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.GetUserOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.GetUserOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.GetUserOutput{
				Errors: []resources.PartialError{}},
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

func Test_UserLookupBy_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.ListUsersByUsernamesOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.ListUsersByUsernamesOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.ListUsersByUsernamesOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.ListUsersByUsernamesOutput{
				Errors: []resources.PartialError{}},
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

func Test_UserLookupByUsername_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.GetUserByUsernameOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.GetUserByUsernameOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.GetUserByUsernameOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.GetUserByUsernameOutput{
				Errors: []resources.PartialError{}},
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

func Test_UserLookupMe_HasPartialError(t *testing.T) {
	var errorTitle string = "test partical error"
	cases := []struct {
		name   string
		res    *types.GetMeOutput
		expect bool
	}{
		{
			name: "has partical error",
			res: &types.GetMeOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				}},
			expect: true,
		},
		{
			name: "has no partical error",
			res: &types.GetMeOutput{
				Errors: []resources.PartialError{}},
			expect: false,
		},
		{
			name: "partical error is nil",
			res: &types.GetMeOutput{
				Errors: []resources.PartialError{}},
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
