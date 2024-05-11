// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/unkeyed/unkey/internal/utils"
)

// Permissions - A query for which permissions you require
type Permissions struct {
}

// Authorization - Perform RBAC checks
type Authorization struct {
	// A query for which permissions you require
	Permissions *Permissions `json:"permissions,omitempty"`
}

func (o *Authorization) GetPermissions() *Permissions {
	if o == nil {
		return nil
	}
	return o.Permissions
}

type V1KeysVerifyKeyRequestRatelimit struct {
	// Override how many tokens are deducted during the ratelimit operation.
	Cost *int64 `default:"1" json:"cost"`
}

func (v V1KeysVerifyKeyRequestRatelimit) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V1KeysVerifyKeyRequestRatelimit) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V1KeysVerifyKeyRequestRatelimit) GetCost() *int64 {
	if o == nil {
		return nil
	}
	return o.Cost
}

type V1KeysVerifyKeyRequest struct {
	// The id of the api where the key belongs to. This is optional for now but will be required soon.
	// The key will be verified against the api's configuration. If the key does not belong to the api, the verification will fail.
	APIID *string `json:"apiId,omitempty"`
	// The key to verify
	Key string `json:"key"`
	// Perform RBAC checks
	Authorization *Authorization                   `json:"authorization,omitempty"`
	Ratelimit     *V1KeysVerifyKeyRequestRatelimit `json:"ratelimit,omitempty"`
}

func (o *V1KeysVerifyKeyRequest) GetAPIID() *string {
	if o == nil {
		return nil
	}
	return o.APIID
}

func (o *V1KeysVerifyKeyRequest) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *V1KeysVerifyKeyRequest) GetAuthorization() *Authorization {
	if o == nil {
		return nil
	}
	return o.Authorization
}

func (o *V1KeysVerifyKeyRequest) GetRatelimit() *V1KeysVerifyKeyRequestRatelimit {
	if o == nil {
		return nil
	}
	return o.Ratelimit
}
