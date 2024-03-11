// File generated from our OpenAPI spec by Stainless.

package logpush

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ValidateService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewValidateService] method instead.
type ValidateService struct {
	Options []option.RequestOption
}

// NewValidateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewValidateService(opts ...option.RequestOption) (r *ValidateService) {
	r = &ValidateService{}
	r.Options = opts
	return
}

// Checks if there is an existing job with a destination.
func (r *ValidateService) Destination(ctx context.Context, params ValidateDestinationParams, opts ...option.RequestOption) (res *ValidateDestinationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ValidateDestinationResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/validate/destination/exists", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Validates logpull origin with logpull_options.
func (r *ValidateService) Origin(ctx context.Context, params ValidateOriginParams, opts ...option.RequestOption) (res *ValidateOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ValidateOriginResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/validate/origin", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ValidateDestinationResponse struct {
	Exists bool                            `json:"exists"`
	JSON   validateDestinationResponseJSON `json:"-"`
}

// validateDestinationResponseJSON contains the JSON metadata for the struct
// [ValidateDestinationResponse]
type validateDestinationResponseJSON struct {
	Exists      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateDestinationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateDestinationResponseJSON) RawJSON() string {
	return r.raw
}

type ValidateOriginResponse struct {
	Message string                     `json:"message"`
	Valid   bool                       `json:"valid"`
	JSON    validateOriginResponseJSON `json:"-"`
}

// validateOriginResponseJSON contains the JSON metadata for the struct
// [ValidateOriginResponse]
type validateOriginResponseJSON struct {
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateOriginResponseJSON) RawJSON() string {
	return r.raw
}

type ValidateDestinationParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r ValidateDestinationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ValidateDestinationResponseEnvelope struct {
	Errors   []ValidateDestinationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ValidateDestinationResponseEnvelopeMessages `json:"messages,required"`
	Result   ValidateDestinationResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ValidateDestinationResponseEnvelopeSuccess `json:"success,required"`
	JSON    validateDestinationResponseEnvelopeJSON    `json:"-"`
}

// validateDestinationResponseEnvelopeJSON contains the JSON metadata for the
// struct [ValidateDestinationResponseEnvelope]
type validateDestinationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateDestinationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateDestinationResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ValidateDestinationResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    validateDestinationResponseEnvelopeErrorsJSON `json:"-"`
}

// validateDestinationResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ValidateDestinationResponseEnvelopeErrors]
type validateDestinationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateDestinationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateDestinationResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ValidateDestinationResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    validateDestinationResponseEnvelopeMessagesJSON `json:"-"`
}

// validateDestinationResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ValidateDestinationResponseEnvelopeMessages]
type validateDestinationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateDestinationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateDestinationResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ValidateDestinationResponseEnvelopeSuccess bool

const (
	ValidateDestinationResponseEnvelopeSuccessTrue ValidateDestinationResponseEnvelopeSuccess = true
)

type ValidateOriginParams struct {
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options,required" format:"uri-reference"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r ValidateOriginParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ValidateOriginResponseEnvelope struct {
	Errors   []ValidateOriginResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ValidateOriginResponseEnvelopeMessages `json:"messages,required"`
	Result   ValidateOriginResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ValidateOriginResponseEnvelopeSuccess `json:"success,required"`
	JSON    validateOriginResponseEnvelopeJSON    `json:"-"`
}

// validateOriginResponseEnvelopeJSON contains the JSON metadata for the struct
// [ValidateOriginResponseEnvelope]
type validateOriginResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateOriginResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ValidateOriginResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    validateOriginResponseEnvelopeErrorsJSON `json:"-"`
}

// validateOriginResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ValidateOriginResponseEnvelopeErrors]
type validateOriginResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateOriginResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateOriginResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ValidateOriginResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    validateOriginResponseEnvelopeMessagesJSON `json:"-"`
}

// validateOriginResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ValidateOriginResponseEnvelopeMessages]
type validateOriginResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValidateOriginResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r validateOriginResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ValidateOriginResponseEnvelopeSuccess bool

const (
	ValidateOriginResponseEnvelopeSuccessTrue ValidateOriginResponseEnvelopeSuccess = true
)
