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

// ZoneSettingHotlinkProtectionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSettingHotlinkProtectionService] method instead.
type ZoneSettingHotlinkProtectionService struct {
	Options []option.RequestOption
}

// NewZoneSettingHotlinkProtectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingHotlinkProtectionService(opts ...option.RequestOption) (r *ZoneSettingHotlinkProtectionService) {
	r = &ZoneSettingHotlinkProtectionService{}
	r.Options = opts
	return
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
func (r *ZoneSettingHotlinkProtectionService) Edit(ctx context.Context, params ZoneSettingHotlinkProtectionEditParams, opts ...option.RequestOption) (res *ZonesHotlinkProtection, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingHotlinkProtectionEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/hotlink_protection", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
func (r *ZoneSettingHotlinkProtectionService) Get(ctx context.Context, query ZoneSettingHotlinkProtectionGetParams, opts ...option.RequestOption) (res *ZonesHotlinkProtection, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingHotlinkProtectionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/hotlink_protection", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID ZonesHotlinkProtectionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesHotlinkProtectionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesHotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesHotlinkProtectionJSON `json:"-"`
}

// zonesHotlinkProtectionJSON contains the JSON metadata for the struct
// [ZonesHotlinkProtection]
type zonesHotlinkProtectionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesHotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesHotlinkProtection) implementsZoneSettingEditResponse() {}

func (r ZonesHotlinkProtection) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesHotlinkProtectionID string

const (
	ZonesHotlinkProtectionIDHotlinkProtection ZonesHotlinkProtectionID = "hotlink_protection"
)

// Current value of the zone setting.
type ZonesHotlinkProtectionValue string

const (
	ZonesHotlinkProtectionValueOn  ZonesHotlinkProtectionValue = "on"
	ZonesHotlinkProtectionValueOff ZonesHotlinkProtectionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesHotlinkProtectionEditable bool

const (
	ZonesHotlinkProtectionEditableTrue  ZonesHotlinkProtectionEditable = true
	ZonesHotlinkProtectionEditableFalse ZonesHotlinkProtectionEditable = false
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZonesHotlinkProtectionParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesHotlinkProtectionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesHotlinkProtectionValue] `json:"value,required"`
}

func (r ZonesHotlinkProtectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesHotlinkProtectionParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingHotlinkProtectionEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingHotlinkProtectionEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingHotlinkProtectionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingHotlinkProtectionEditParamsValue string

const (
	ZoneSettingHotlinkProtectionEditParamsValueOn  ZoneSettingHotlinkProtectionEditParamsValue = "on"
	ZoneSettingHotlinkProtectionEditParamsValueOff ZoneSettingHotlinkProtectionEditParamsValue = "off"
)

type ZoneSettingHotlinkProtectionEditResponseEnvelope struct {
	Errors   []ZoneSettingHotlinkProtectionEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingHotlinkProtectionEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, the Hotlink Protection option ensures that other sites cannot suck
	// up your bandwidth by building pages that use images hosted on your site. Anytime
	// a request for an image on your site hits Cloudflare, we check to ensure that
	// it's not another site requesting them. People will still be able to download and
	// view images from your page, but other sites won't be able to steal them for use
	// on their own pages.
	// (https://support.cloudflare.com/hc/en-us/articles/200170026).
	Result ZonesHotlinkProtection                               `json:"result"`
	JSON   zoneSettingHotlinkProtectionEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingHotlinkProtectionEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingHotlinkProtectionEditResponseEnvelope]
type zoneSettingHotlinkProtectionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHotlinkProtectionEditResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingHotlinkProtectionEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingHotlinkProtectionEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingHotlinkProtectionEditResponseEnvelopeErrors]
type zoneSettingHotlinkProtectionEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHotlinkProtectionEditResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zoneSettingHotlinkProtectionEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingHotlinkProtectionEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingHotlinkProtectionEditResponseEnvelopeMessages]
type zoneSettingHotlinkProtectionEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHotlinkProtectionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingHotlinkProtectionGetResponseEnvelope struct {
	Errors   []ZoneSettingHotlinkProtectionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingHotlinkProtectionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// When enabled, the Hotlink Protection option ensures that other sites cannot suck
	// up your bandwidth by building pages that use images hosted on your site. Anytime
	// a request for an image on your site hits Cloudflare, we check to ensure that
	// it's not another site requesting them. People will still be able to download and
	// view images from your page, but other sites won't be able to steal them for use
	// on their own pages.
	// (https://support.cloudflare.com/hc/en-us/articles/200170026).
	Result ZonesHotlinkProtection                              `json:"result"`
	JSON   zoneSettingHotlinkProtectionGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingHotlinkProtectionGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingHotlinkProtectionGetResponseEnvelope]
type zoneSettingHotlinkProtectionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHotlinkProtectionGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingHotlinkProtectionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingHotlinkProtectionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingHotlinkProtectionGetResponseEnvelopeErrors]
type zoneSettingHotlinkProtectionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingHotlinkProtectionGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zoneSettingHotlinkProtectionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingHotlinkProtectionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneSettingHotlinkProtectionGetResponseEnvelopeMessages]
type zoneSettingHotlinkProtectionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingHotlinkProtectionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
