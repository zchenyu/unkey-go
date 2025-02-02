// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/internal/utils"
	"github.com/unkeyed/unkey-go/models/components"
)

type ListIdentitiesRequest struct {
	Environment *string `default:"default" queryParam:"style=form,explode=true,name=environment"`
	Limit       *int64  `default:"100" queryParam:"style=form,explode=true,name=limit"`
	Cursor      *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListIdentitiesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListIdentitiesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListIdentitiesRequest) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *ListIdentitiesRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListIdentitiesRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListIdentitiesRatelimits struct {
	// The name of this limit. You will need to use this again when verifying a key.
	Name string `json:"name"`
	// How many requests may pass within a given window before requests are rejected.
	Limit int64 `json:"limit"`
	// The duration for each ratelimit window in milliseconds.
	Duration int64 `json:"duration"`
}

func (o *ListIdentitiesRatelimits) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListIdentitiesRatelimits) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *ListIdentitiesRatelimits) GetDuration() int64 {
	if o == nil {
		return 0
	}
	return o.Duration
}

type Identities struct {
	// The id of this identity. Used to interact with unkey's API
	ID string `json:"id"`
	// The id in your system
	ExternalID string `json:"externalId"`
	// When verifying keys, you can specify which limits you want to use and all keys attached to this identity, will share the limits.
	Ratelimits []ListIdentitiesRatelimits `json:"ratelimits"`
}

func (o *Identities) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Identities) GetExternalID() string {
	if o == nil {
		return ""
	}
	return o.ExternalID
}

func (o *Identities) GetRatelimits() []ListIdentitiesRatelimits {
	if o == nil {
		return []ListIdentitiesRatelimits{}
	}
	return o.Ratelimits
}

// ListIdentitiesResponseBody - A list of identities.
type ListIdentitiesResponseBody struct {
	// A list of identities.
	Identities []Identities `json:"identities"`
	// The cursor to use for the next page of results, if no cursor is returned, there are no more results
	Cursor *string `json:"cursor,omitempty"`
	// The total number of identities for this environment
	Total int64 `json:"total"`
}

func (o *ListIdentitiesResponseBody) GetIdentities() []Identities {
	if o == nil {
		return []Identities{}
	}
	return o.Identities
}

func (o *ListIdentitiesResponseBody) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListIdentitiesResponseBody) GetTotal() int64 {
	if o == nil {
		return 0
	}
	return o.Total
}

type ListIdentitiesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// A list of identities.
	Object *ListIdentitiesResponseBody

	Next func() (*ListIdentitiesResponse, error)
}

func (o *ListIdentitiesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListIdentitiesResponse) GetObject() *ListIdentitiesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
