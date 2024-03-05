// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSettingOpportunisticEncryptionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingOpportunisticEncryptionService] method instead.
type ZoneSettingOpportunisticEncryptionService struct {
	Options []option.RequestOption
}

// NewZoneSettingOpportunisticEncryptionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneSettingOpportunisticEncryptionService(opts ...option.RequestOption) (r *ZoneSettingOpportunisticEncryptionService) {
	r = &ZoneSettingOpportunisticEncryptionService{}
	r.Options = opts
	return
}

// Changes Opportunistic Encryption setting.
func (r *ZoneSettingOpportunisticEncryptionService) Edit(ctx context.Context, params ZoneSettingOpportunisticEncryptionEditParams, opts ...option.RequestOption) (res *ZonesOpportunisticEncryption, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingOpportunisticEncryptionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/opportunistic_encryption", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets Opportunistic Encryption setting.
func (r *ZoneSettingOpportunisticEncryptionService) Get(ctx context.Context, query ZoneSettingOpportunisticEncryptionGetParams, opts ...option.RequestOption) (res *ZonesOpportunisticEncryption, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingOpportunisticEncryptionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/opportunistic_encryption", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enables the Opportunistic Encryption feature for a zone.
type ZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID ZonesOpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesOpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesOpportunisticEncryptionJSON `json:"-"`
}

// zonesOpportunisticEncryptionJSON contains the JSON metadata for the struct
// [ZonesOpportunisticEncryption]
type zonesOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesOpportunisticEncryption) implementsZoneSettingEditResponse() {}

func (r ZonesOpportunisticEncryption) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesOpportunisticEncryptionID string

const (
	ZonesOpportunisticEncryptionIDOpportunisticEncryption ZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

// Current value of the zone setting.
type ZonesOpportunisticEncryptionValue string

const (
	ZonesOpportunisticEncryptionValueOn  ZonesOpportunisticEncryptionValue = "on"
	ZonesOpportunisticEncryptionValueOff ZonesOpportunisticEncryptionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesOpportunisticEncryptionEditable bool

const (
	ZonesOpportunisticEncryptionEditableTrue  ZonesOpportunisticEncryptionEditable = true
	ZonesOpportunisticEncryptionEditableFalse ZonesOpportunisticEncryptionEditable = false
)

// Enables the Opportunistic Encryption feature for a zone.
type ZonesOpportunisticEncryptionParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesOpportunisticEncryptionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesOpportunisticEncryptionValue] `json:"value,required"`
}

func (r ZonesOpportunisticEncryptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesOpportunisticEncryptionParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingOpportunisticEncryptionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingOpportunisticEncryptionEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingOpportunisticEncryptionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingOpportunisticEncryptionEditParamsValue string

const (
	ZoneSettingOpportunisticEncryptionEditParamsValueOn  ZoneSettingOpportunisticEncryptionEditParamsValue = "on"
	ZoneSettingOpportunisticEncryptionEditParamsValueOff ZoneSettingOpportunisticEncryptionEditParamsValue = "off"
)

type ZoneSettingOpportunisticEncryptionEditResponseEnvelope struct {
	Errors   []ZoneSettingOpportunisticEncryptionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingOpportunisticEncryptionEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result ZonesOpportunisticEncryption                               `json:"result"`
	JSON   zoneSettingOpportunisticEncryptionEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingOpportunisticEncryptionEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZoneSettingOpportunisticEncryptionEditResponseEnvelope]
type zoneSettingOpportunisticEncryptionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticEncryptionEditResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zoneSettingOpportunisticEncryptionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingOpportunisticEncryptionEditResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZoneSettingOpportunisticEncryptionEditResponseEnvelopeErrors]
type zoneSettingOpportunisticEncryptionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticEncryptionEditResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    zoneSettingOpportunisticEncryptionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingOpportunisticEncryptionEditResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZoneSettingOpportunisticEncryptionEditResponseEnvelopeMessages]
type zoneSettingOpportunisticEncryptionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticEncryptionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingOpportunisticEncryptionGetResponseEnvelope struct {
	Errors   []ZoneSettingOpportunisticEncryptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingOpportunisticEncryptionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enables the Opportunistic Encryption feature for a zone.
	Result ZonesOpportunisticEncryption                              `json:"result"`
	JSON   zoneSettingOpportunisticEncryptionGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingOpportunisticEncryptionGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZoneSettingOpportunisticEncryptionGetResponseEnvelope]
type zoneSettingOpportunisticEncryptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticEncryptionGetResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zoneSettingOpportunisticEncryptionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingOpportunisticEncryptionGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZoneSettingOpportunisticEncryptionGetResponseEnvelopeErrors]
type zoneSettingOpportunisticEncryptionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingOpportunisticEncryptionGetResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zoneSettingOpportunisticEncryptionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingOpportunisticEncryptionGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZoneSettingOpportunisticEncryptionGetResponseEnvelopeMessages]
type zoneSettingOpportunisticEncryptionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOpportunisticEncryptionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
