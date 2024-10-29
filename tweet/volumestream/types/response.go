package types

import "github.com/xxiiaaon/gotwi/resources"

type SampleStreamOutput struct {
	Data     resources.Tweet `json:"data"`
	Includes struct {
		Users  []resources.User  `json:"users,omitempty"`
		Tweets []resources.Tweet `json:"tweets,omitempty"`
		Places []resources.Place `json:"places,omitempty"`
		Media  []resources.Media `json:"media,omitempty"`
		Polls  []resources.Poll  `json:"polls,omitempty"`
	} `json:"includes,omitempty"`
	Errors []resources.PartialError `json:"errors,omitempty"`
}

func (r *SampleStreamOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
