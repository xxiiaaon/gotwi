package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/xxiiaaon/gotwi/fields"
	"github.com/xxiiaaon/gotwi/internal/util"
)

type ListMembershipsMaxResults int

func (m ListMembershipsMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListMembershipsMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListMembershipsInput struct {
	accessToken string

	// Path parameter
	ID string // User ID

	// Query parameters
	MaxResults      ListMembershipsMaxResults
	PaginationToken string
	Expansions      fields.ExpansionList
	ListFields      fields.ListFieldList
	UserFields      fields.UserFieldList
}

var ListMembersListMembershipsQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"list.fields":      {},
	"user.fields":      {},
}

func (p *ListMembershipsInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListMembershipsInput) AccessToken() string {
	return p.accessToken
}

func (p *ListMembershipsInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, ListMembersListMembershipsQueryParams)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListMembershipsInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListMembershipsInput) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	m = fields.SetFieldsParams(m, p.Expansions, p.ListFields, p.UserFields)

	return m
}

type ListMembersGetMaxResults int

func (m ListMembersGetMaxResults) Valid() bool {
	return m > 1 && m <= 100
}

func (m ListMembersGetMaxResults) String() string {
	return strconv.Itoa(int(m))
}

type ListInput struct {
	accessToken string

	// Path parameter
	ID string // List ID

	// Query parameters
	Expansions      fields.ExpansionList
	ListFields      fields.ListFieldList
	UserFields      fields.UserFieldList
	MaxResults      ListMembersGetMaxResults
	PaginationToken string
}

func (p *ListInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListInput) AccessToken() string {
	return p.accessToken
}

var listQueryParameters = map[string]struct{}{
	"expansions":       {},
	"list.fields":      {},
	"user.fields":      {},
	"max_results":      {},
	"pagination_token": {},
}

func (p *ListInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m = fields.SetFieldsParams(m, p.Expansions, p.ListFields, p.UserFields)

	if p.MaxResults.Valid() {
		m["max_results"] = p.MaxResults.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	return m
}

type CreateInput struct {
	accessToken string

	// Path parameter
	ID string `json:"-"` // List ID

	// JSON body parameter
	UserID string `json:"user_id,"` // required
}

func (p *CreateInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *CreateInput) AccessToken() string {
	return p.accessToken
}

func (p *CreateInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *CreateInput) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *CreateInput) ParameterMap() map[string]string {
	return map[string]string{}
}

type DeleteInput struct {
	accessToken string

	// Path parameter
	ID     string // List ID
	UserID string
}

func (p *DeleteInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *DeleteInput) AccessToken() string {
	return p.accessToken
}

func (p *DeleteInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" || p.UserID == "" {
		return ""
	}

	escapedID := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escapedID, 1)
	escapedUserID := url.QueryEscape(p.UserID)
	endpoint = strings.Replace(endpoint, ":user_id", escapedUserID, 1)

	return endpoint
}

func (p *DeleteInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *DeleteInput) ParameterMap() map[string]string {
	return map[string]string{}
}
