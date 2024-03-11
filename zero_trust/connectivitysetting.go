// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ConnectivitySettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewConnectivitySettingService]
// method instead.
type ConnectivitySettingService struct {
	Options []option.RequestOption
}

// NewConnectivitySettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectivitySettingService(opts ...option.RequestOption) (r *ConnectivitySettingService) {
	r = &ConnectivitySettingService{}
	r.Options = opts
	return
}

// Updates the Zero Trust Connectivity Settings for the given account.
func (r *ConnectivitySettingService) Edit(ctx context.Context, params ConnectivitySettingEditParams, opts ...option.RequestOption) (res *ConnectivitySettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ConnectivitySettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the Zero Trust Connectivity Settings for the given account.
func (r *ConnectivitySettingService) Get(ctx context.Context, query ConnectivitySettingGetParams, opts ...option.RequestOption) (res *ConnectivitySettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ConnectivitySettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConnectivitySettingEditResponse struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWARPEnabled bool                                `json:"offramp_warp_enabled"`
	JSON               connectivitySettingEditResponseJSON `json:"-"`
}

// connectivitySettingEditResponseJSON contains the JSON metadata for the struct
// [ConnectivitySettingEditResponse]
type connectivitySettingEditResponseJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWARPEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ConnectivitySettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivitySettingEditResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectivitySettingGetResponse struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWARPEnabled bool                               `json:"offramp_warp_enabled"`
	JSON               connectivitySettingGetResponseJSON `json:"-"`
}

// connectivitySettingGetResponseJSON contains the JSON metadata for the struct
// [ConnectivitySettingGetResponse]
type connectivitySettingGetResponseJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWARPEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ConnectivitySettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivitySettingGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectivitySettingEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled param.Field[bool] `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWARPEnabled param.Field[bool] `json:"offramp_warp_enabled"`
}

func (r ConnectivitySettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivitySettingEditResponseEnvelope struct {
	Errors   []ConnectivitySettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConnectivitySettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ConnectivitySettingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ConnectivitySettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    connectivitySettingEditResponseEnvelopeJSON    `json:"-"`
}

// connectivitySettingEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ConnectivitySettingEditResponseEnvelope]
type connectivitySettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivitySettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivitySettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectivitySettingEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    connectivitySettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// connectivitySettingEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ConnectivitySettingEditResponseEnvelopeErrors]
type connectivitySettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivitySettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivitySettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectivitySettingEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    connectivitySettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// connectivitySettingEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ConnectivitySettingEditResponseEnvelopeMessages]
type connectivitySettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivitySettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivitySettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConnectivitySettingEditResponseEnvelopeSuccess bool

const (
	ConnectivitySettingEditResponseEnvelopeSuccessTrue ConnectivitySettingEditResponseEnvelopeSuccess = true
)

type ConnectivitySettingGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConnectivitySettingGetResponseEnvelope struct {
	Errors   []ConnectivitySettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConnectivitySettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ConnectivitySettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ConnectivitySettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    connectivitySettingGetResponseEnvelopeJSON    `json:"-"`
}

// connectivitySettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ConnectivitySettingGetResponseEnvelope]
type connectivitySettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivitySettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivitySettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectivitySettingGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    connectivitySettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// connectivitySettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ConnectivitySettingGetResponseEnvelopeErrors]
type connectivitySettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivitySettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivitySettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectivitySettingGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    connectivitySettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// connectivitySettingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ConnectivitySettingGetResponseEnvelopeMessages]
type connectivitySettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivitySettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivitySettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConnectivitySettingGetResponseEnvelopeSuccess bool

const (
	ConnectivitySettingGetResponseEnvelopeSuccessTrue ConnectivitySettingGetResponseEnvelopeSuccess = true
)
