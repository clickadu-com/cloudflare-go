// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// AccessGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessGroupService] method instead.
type AccessGroupService struct {
	Options []option.RequestOption
}

// NewAccessGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessGroupService(opts ...option.RequestOption) (r *AccessGroupService) {
	r = &AccessGroupService{}
	r.Options = opts
	return
}

// Creates a new Access group.
func (r *AccessGroupService) New(ctx context.Context, params AccessGroupNewParams, opts ...option.RequestOption) (res *ZeroTrustGroup, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Access groups.
func (r *AccessGroupService) List(ctx context.Context, query AccessGroupListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ZeroTrustGroup], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all Access groups.
func (r *AccessGroupService) ListAutoPaging(ctx context.Context, query AccessGroupListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ZeroTrustGroup] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type ZeroTrustGroup struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessRule `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessRule       `json:"require"`
	UpdatedAt time.Time          `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGroupJSON `json:"-"`
}

// zeroTrustGroupJSON contains the JSON metadata for the struct [ZeroTrustGroup]
type zeroTrustGroupJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	IsDefault   apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupJSON) RawJSON() string {
	return r.raw
}

type AccessGroupNewParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
}

func (r AccessGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessGroupNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccessGroupNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ZeroTrustGroup                        `json:"result"`
	JSON    accessGroupNewResponseEnvelopeJSON    `json:"-"`
}

// accessGroupNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGroupNewResponseEnvelope]
type accessGroupNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGroupNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessGroupNewResponseEnvelopeSuccess bool

const (
	AccessGroupNewResponseEnvelopeSuccessTrue AccessGroupNewResponseEnvelopeSuccess = true
)

func (r AccessGroupNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessGroupListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}
