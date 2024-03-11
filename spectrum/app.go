// File generated from our OpenAPI spec by Stainless.

package spectrum

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// AppService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r *AppService) {
	r = &AppService{}
	r.Options = opts
	return
}

// Creates a new Spectrum application from a configuration using a name for the
// origin.
func (r *AppService) New(ctx context.Context, zone string, body AppNewParams, opts ...option.RequestOption) (res *AppNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AppNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps", zone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a previously existing application's configuration that uses a name for
// the origin.
func (r *AppService) Update(ctx context.Context, zone string, appID string, body AppUpdateParams, opts ...option.RequestOption) (res *AppUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AppUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zone, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a list of currently existing Spectrum applications inside a zone.
func (r *AppService) List(ctx context.Context, zone string, query AppListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[AppListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/spectrum/apps", zone)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Retrieves a list of currently existing Spectrum applications inside a zone.
func (r *AppService) ListAutoPaging(ctx context.Context, zone string, query AppListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[AppListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zone, query, opts...))
}

// Deletes a previously existing application.
func (r *AppService) Delete(ctx context.Context, zone string, appID string, opts ...option.RequestOption) (res *AppDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AppDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zone, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the application configuration of a specific application inside a zone.
func (r *AppService) Get(ctx context.Context, zone string, appID string, opts ...option.RequestOption) (res *AppGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AppGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", zone, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AppNewResponse struct {
	// Application identifier.
	ID string `json:"id"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS AppNewResponseDNS `json:"dns"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs AppNewResponseEdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS AppNewResponseOriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort AppNewResponseOriginPort `json:"origin_port"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppNewResponseProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS AppNewResponseTLS `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppNewResponseTrafficType `json:"traffic_type"`
	JSON        appNewResponseJSON        `json:"-"`
}

// appNewResponseJSON contains the JSON metadata for the struct [AppNewResponse]
type appNewResponseJSON struct {
	ID               apijson.Field
	ArgoSmartRouting apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	ModifiedOn       apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	Protocol         apijson.Field
	ProxyProtocol    apijson.Field
	TLS              apijson.Field
	TrafficType      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseJSON) RawJSON() string {
	return r.raw
}

// The name and type of DNS record for the Spectrum application.
type AppNewResponseDNS struct {
	// The name of the DNS record associated with the application.
	Name string `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type AppNewResponseDNSType `json:"type"`
	JSON appNewResponseDNSJSON `json:"-"`
}

// appNewResponseDNSJSON contains the JSON metadata for the struct
// [AppNewResponseDNS]
type appNewResponseDNSJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the application.
type AppNewResponseDNSType string

const (
	AppNewResponseDNSTypeCNAME   AppNewResponseDNSType = "CNAME"
	AppNewResponseDNSTypeAddress AppNewResponseDNSType = "ADDRESS"
)

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by
// [spectrum.AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable] or
// [spectrum.AppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable].
type AppNewResponseEdgeIPs interface {
	implementsSpectrumAppNewResponseEdgeIPs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppNewResponseEdgeIPs)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable{}),
		},
	)
}

type AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType `json:"type"`
	JSON appNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON `json:"-"`
}

// appNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON contains the JSON
// metadata for the struct [AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable]
type appNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON) RawJSON() string {
	return r.raw
}

func (r AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable) implementsSpectrumAppNewResponseEdgeIPs() {
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity string

const (
	AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityAll  AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "all"
	AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV4 AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv4"
	AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV6 AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType string

const (
	AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableTypeDynamic AppNewResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType = "dynamic"
)

type AppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs []string `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type AppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType `json:"type"`
	JSON appNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON `json:"-"`
}

// appNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON contains the
// JSON metadata for the struct
// [AppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable]
type appNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON struct {
	IPs         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON) RawJSON() string {
	return r.raw
}

func (r AppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) implementsSpectrumAppNewResponseEdgeIPs() {
}

// The type of edge IP configuration specified. Statically allocated edge IPs use
// customer IPs in accordance with the ips array you specify. Only valid with
// ADDRESS DNS names.
type AppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType string

const (
	AppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableTypeStatic AppNewResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType = "static"
)

// The name and type of DNS record for the Spectrum application.
type AppNewResponseOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name string `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL int64 `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type AppNewResponseOriginDNSType `json:"type"`
	JSON appNewResponseOriginDNSJSON `json:"-"`
}

// appNewResponseOriginDNSJSON contains the JSON metadata for the struct
// [AppNewResponseOriginDNS]
type appNewResponseOriginDNSJSON struct {
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseOriginDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseOriginDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type AppNewResponseOriginDNSType string

const (
	AppNewResponseOriginDNSTypeEmpty AppNewResponseOriginDNSType = ""
	AppNewResponseOriginDNSTypeA     AppNewResponseOriginDNSType = "A"
	AppNewResponseOriginDNSTypeAAAA  AppNewResponseOriginDNSType = "AAAA"
	AppNewResponseOriginDNSTypeSRV   AppNewResponseOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Union satisfied by [shared.UnionInt] or [shared.UnionString].
type AppNewResponseOriginPort interface {
	ImplementsSpectrumAppNewResponseOriginPort()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppNewResponseOriginPort)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppNewResponseProxyProtocol string

const (
	AppNewResponseProxyProtocolOff    AppNewResponseProxyProtocol = "off"
	AppNewResponseProxyProtocolV1     AppNewResponseProxyProtocol = "v1"
	AppNewResponseProxyProtocolV2     AppNewResponseProxyProtocol = "v2"
	AppNewResponseProxyProtocolSimple AppNewResponseProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type AppNewResponseTLS string

const (
	AppNewResponseTLSOff      AppNewResponseTLS = "off"
	AppNewResponseTLSFlexible AppNewResponseTLS = "flexible"
	AppNewResponseTLSFull     AppNewResponseTLS = "full"
	AppNewResponseTLSStrict   AppNewResponseTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppNewResponseTrafficType string

const (
	AppNewResponseTrafficTypeDirect AppNewResponseTrafficType = "direct"
	AppNewResponseTrafficTypeHTTP   AppNewResponseTrafficType = "http"
	AppNewResponseTrafficTypeHTTPS  AppNewResponseTrafficType = "https"
)

type AppUpdateResponse struct {
	// Application identifier.
	ID string `json:"id"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS AppUpdateResponseDNS `json:"dns"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs AppUpdateResponseEdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS AppUpdateResponseOriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort AppUpdateResponseOriginPort `json:"origin_port"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppUpdateResponseProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS AppUpdateResponseTLS `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppUpdateResponseTrafficType `json:"traffic_type"`
	JSON        appUpdateResponseJSON        `json:"-"`
}

// appUpdateResponseJSON contains the JSON metadata for the struct
// [AppUpdateResponse]
type appUpdateResponseJSON struct {
	ID               apijson.Field
	ArgoSmartRouting apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	ModifiedOn       apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	Protocol         apijson.Field
	ProxyProtocol    apijson.Field
	TLS              apijson.Field
	TrafficType      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The name and type of DNS record for the Spectrum application.
type AppUpdateResponseDNS struct {
	// The name of the DNS record associated with the application.
	Name string `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type AppUpdateResponseDNSType `json:"type"`
	JSON appUpdateResponseDNSJSON `json:"-"`
}

// appUpdateResponseDNSJSON contains the JSON metadata for the struct
// [AppUpdateResponseDNS]
type appUpdateResponseDNSJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the application.
type AppUpdateResponseDNSType string

const (
	AppUpdateResponseDNSTypeCNAME   AppUpdateResponseDNSType = "CNAME"
	AppUpdateResponseDNSTypeAddress AppUpdateResponseDNSType = "ADDRESS"
)

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by
// [spectrum.AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable] or
// [spectrum.AppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable].
type AppUpdateResponseEdgeIPs interface {
	implementsSpectrumAppUpdateResponseEdgeIPs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppUpdateResponseEdgeIPs)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable{}),
		},
	)
}

type AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType `json:"type"`
	JSON appUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON `json:"-"`
}

// appUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON contains the JSON
// metadata for the struct
// [AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable]
type appUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableJSON) RawJSON() string {
	return r.raw
}

func (r AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariable) implementsSpectrumAppUpdateResponseEdgeIPs() {
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity string

const (
	AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityAll  AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "all"
	AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV4 AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv4"
	AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV6 AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType string

const (
	AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableTypeDynamic AppUpdateResponseEdgeIPsSpectrumEdgeIPEyeballIPsVariableType = "dynamic"
)

type AppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs []string `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type AppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType `json:"type"`
	JSON appUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON `json:"-"`
}

// appUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON contains the
// JSON metadata for the struct
// [AppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable]
type appUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON struct {
	IPs         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableJSON) RawJSON() string {
	return r.raw
}

func (r AppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) implementsSpectrumAppUpdateResponseEdgeIPs() {
}

// The type of edge IP configuration specified. Statically allocated edge IPs use
// customer IPs in accordance with the ips array you specify. Only valid with
// ADDRESS DNS names.
type AppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType string

const (
	AppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableTypeStatic AppUpdateResponseEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType = "static"
)

// The name and type of DNS record for the Spectrum application.
type AppUpdateResponseOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name string `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL int64 `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type AppUpdateResponseOriginDNSType `json:"type"`
	JSON appUpdateResponseOriginDNSJSON `json:"-"`
}

// appUpdateResponseOriginDNSJSON contains the JSON metadata for the struct
// [AppUpdateResponseOriginDNS]
type appUpdateResponseOriginDNSJSON struct {
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseOriginDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseOriginDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type AppUpdateResponseOriginDNSType string

const (
	AppUpdateResponseOriginDNSTypeEmpty AppUpdateResponseOriginDNSType = ""
	AppUpdateResponseOriginDNSTypeA     AppUpdateResponseOriginDNSType = "A"
	AppUpdateResponseOriginDNSTypeAAAA  AppUpdateResponseOriginDNSType = "AAAA"
	AppUpdateResponseOriginDNSTypeSRV   AppUpdateResponseOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Union satisfied by [shared.UnionInt] or [shared.UnionString].
type AppUpdateResponseOriginPort interface {
	ImplementsSpectrumAppUpdateResponseOriginPort()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppUpdateResponseOriginPort)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppUpdateResponseProxyProtocol string

const (
	AppUpdateResponseProxyProtocolOff    AppUpdateResponseProxyProtocol = "off"
	AppUpdateResponseProxyProtocolV1     AppUpdateResponseProxyProtocol = "v1"
	AppUpdateResponseProxyProtocolV2     AppUpdateResponseProxyProtocol = "v2"
	AppUpdateResponseProxyProtocolSimple AppUpdateResponseProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type AppUpdateResponseTLS string

const (
	AppUpdateResponseTLSOff      AppUpdateResponseTLS = "off"
	AppUpdateResponseTLSFlexible AppUpdateResponseTLS = "flexible"
	AppUpdateResponseTLSFull     AppUpdateResponseTLS = "full"
	AppUpdateResponseTLSStrict   AppUpdateResponseTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppUpdateResponseTrafficType string

const (
	AppUpdateResponseTrafficTypeDirect AppUpdateResponseTrafficType = "direct"
	AppUpdateResponseTrafficTypeHTTP   AppUpdateResponseTrafficType = "http"
	AppUpdateResponseTrafficTypeHTTPS  AppUpdateResponseTrafficType = "https"
)

type AppListResponse = interface{}

type AppDeleteResponse struct {
	// Application identifier.
	ID   string                `json:"id"`
	JSON appDeleteResponseJSON `json:"-"`
}

// appDeleteResponseJSON contains the JSON metadata for the struct
// [AppDeleteResponse]
type appDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [spectrum.AppGetResponseUnknown] or [shared.UnionString].
type AppGetResponse interface {
	ImplementsSpectrumAppGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AppNewParams struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[AppNewParamsDNS] `json:"dns,required"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[AppNewParamsOriginDNS] `json:"origin_dns,required"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[AppNewParamsOriginPort] `json:"origin_port,required"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[AppNewParamsEdgeIPs] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall param.Field[bool] `json:"ip_firewall"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[AppNewParamsProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[AppNewParamsTLS] `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[AppNewParamsTrafficType] `json:"traffic_type"`
}

func (r AppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name and type of DNS record for the Spectrum application.
type AppNewParamsDNS struct {
	// The name of the DNS record associated with the application.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type param.Field[AppNewParamsDNSType] `json:"type"`
}

func (r AppNewParamsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the application.
type AppNewParamsDNSType string

const (
	AppNewParamsDNSTypeCNAME   AppNewParamsDNSType = "CNAME"
	AppNewParamsDNSTypeAddress AppNewParamsDNSType = "ADDRESS"
)

// The name and type of DNS record for the Spectrum application.
type AppNewParamsOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL param.Field[int64] `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type param.Field[AppNewParamsOriginDNSType] `json:"type"`
}

func (r AppNewParamsOriginDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type AppNewParamsOriginDNSType string

const (
	AppNewParamsOriginDNSTypeEmpty AppNewParamsOriginDNSType = ""
	AppNewParamsOriginDNSTypeA     AppNewParamsOriginDNSType = "A"
	AppNewParamsOriginDNSTypeAAAA  AppNewParamsOriginDNSType = "AAAA"
	AppNewParamsOriginDNSTypeSRV   AppNewParamsOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type AppNewParamsOriginPort interface {
	ImplementsSpectrumAppNewParamsOriginPort()
}

// The anycast edge IP configuration for the hostname of this application.
//
// Satisfied by [spectrum.AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable],
// [spectrum.AppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable].
type AppNewParamsEdgeIPs interface {
	implementsSpectrumAppNewParamsEdgeIPs()
}

type AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType] `json:"type"`
}

func (r AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable) implementsSpectrumAppNewParamsEdgeIPs() {
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity string

const (
	AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityAll  AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "all"
	AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV4 AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv4"
	AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV6 AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType string

const (
	AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableTypeDynamic AppNewParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType = "dynamic"
)

type AppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs param.Field[[]string] `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type param.Field[AppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType] `json:"type"`
}

func (r AppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) implementsSpectrumAppNewParamsEdgeIPs() {
}

// The type of edge IP configuration specified. Statically allocated edge IPs use
// customer IPs in accordance with the ips array you specify. Only valid with
// ADDRESS DNS names.
type AppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType string

const (
	AppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableTypeStatic AppNewParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType = "static"
)

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppNewParamsProxyProtocol string

const (
	AppNewParamsProxyProtocolOff    AppNewParamsProxyProtocol = "off"
	AppNewParamsProxyProtocolV1     AppNewParamsProxyProtocol = "v1"
	AppNewParamsProxyProtocolV2     AppNewParamsProxyProtocol = "v2"
	AppNewParamsProxyProtocolSimple AppNewParamsProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type AppNewParamsTLS string

const (
	AppNewParamsTLSOff      AppNewParamsTLS = "off"
	AppNewParamsTLSFlexible AppNewParamsTLS = "flexible"
	AppNewParamsTLSFull     AppNewParamsTLS = "full"
	AppNewParamsTLSStrict   AppNewParamsTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppNewParamsTrafficType string

const (
	AppNewParamsTrafficTypeDirect AppNewParamsTrafficType = "direct"
	AppNewParamsTrafficTypeHTTP   AppNewParamsTrafficType = "http"
	AppNewParamsTrafficTypeHTTPS  AppNewParamsTrafficType = "https"
)

type AppNewResponseEnvelope struct {
	Errors   []AppNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AppNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AppNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    appNewResponseEnvelopeJSON    `json:"-"`
}

// appNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelope]
type appNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppNewResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    appNewResponseEnvelopeErrorsJSON `json:"-"`
}

// appNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelopeErrors]
type appNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppNewResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    appNewResponseEnvelopeMessagesJSON `json:"-"`
}

// appNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelopeMessages]
type appNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppNewResponseEnvelopeSuccess bool

const (
	AppNewResponseEnvelopeSuccessTrue AppNewResponseEnvelopeSuccess = true
)

type AppUpdateParams struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[AppUpdateParamsDNS] `json:"dns,required"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[AppUpdateParamsOriginDNS] `json:"origin_dns,required"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[AppUpdateParamsOriginPort] `json:"origin_port,required"`
	// The port configuration at Cloudflare’s edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[AppUpdateParamsEdgeIPs] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall param.Field[bool] `json:"ip_firewall"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[AppUpdateParamsProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[AppUpdateParamsTLS] `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[AppUpdateParamsTrafficType] `json:"traffic_type"`
}

func (r AppUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name and type of DNS record for the Spectrum application.
type AppUpdateParamsDNS struct {
	// The name of the DNS record associated with the application.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type param.Field[AppUpdateParamsDNSType] `json:"type"`
}

func (r AppUpdateParamsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the application.
type AppUpdateParamsDNSType string

const (
	AppUpdateParamsDNSTypeCNAME   AppUpdateParamsDNSType = "CNAME"
	AppUpdateParamsDNSTypeAddress AppUpdateParamsDNSType = "ADDRESS"
)

// The name and type of DNS record for the Spectrum application.
type AppUpdateParamsOriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL param.Field[int64] `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type param.Field[AppUpdateParamsOriginDNSType] `json:"type"`
}

func (r AppUpdateParamsOriginDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type AppUpdateParamsOriginDNSType string

const (
	AppUpdateParamsOriginDNSTypeEmpty AppUpdateParamsOriginDNSType = ""
	AppUpdateParamsOriginDNSTypeA     AppUpdateParamsOriginDNSType = "A"
	AppUpdateParamsOriginDNSTypeAAAA  AppUpdateParamsOriginDNSType = "AAAA"
	AppUpdateParamsOriginDNSTypeSRV   AppUpdateParamsOriginDNSType = "SRV"
)

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type AppUpdateParamsOriginPort interface {
	ImplementsSpectrumAppUpdateParamsOriginPort()
}

// The anycast edge IP configuration for the hostname of this application.
//
// Satisfied by [spectrum.AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable],
// [spectrum.AppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable].
type AppUpdateParamsEdgeIPs interface {
	implementsSpectrumAppUpdateParamsEdgeIPs()
}

type AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType] `json:"type"`
}

func (r AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariable) implementsSpectrumAppUpdateParamsEdgeIPs() {
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity string

const (
	AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityAll  AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "all"
	AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV4 AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv4"
	AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivityIPV6 AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableConnectivity = "ipv6"
)

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType string

const (
	AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableTypeDynamic AppUpdateParamsEdgeIPsSpectrumEdgeIPEyeballIPsVariableType = "dynamic"
)

type AppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs param.Field[[]string] `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type param.Field[AppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType] `json:"type"`
}

func (r AppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariable) implementsSpectrumAppUpdateParamsEdgeIPs() {
}

// The type of edge IP configuration specified. Statically allocated edge IPs use
// customer IPs in accordance with the ips array you specify. Only valid with
// ADDRESS DNS names.
type AppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType string

const (
	AppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableTypeStatic AppUpdateParamsEdgeIPsSpectrumEdgeIPCustomerOwnedIPsVariableType = "static"
)

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppUpdateParamsProxyProtocol string

const (
	AppUpdateParamsProxyProtocolOff    AppUpdateParamsProxyProtocol = "off"
	AppUpdateParamsProxyProtocolV1     AppUpdateParamsProxyProtocol = "v1"
	AppUpdateParamsProxyProtocolV2     AppUpdateParamsProxyProtocol = "v2"
	AppUpdateParamsProxyProtocolSimple AppUpdateParamsProxyProtocol = "simple"
)

// The type of TLS termination associated with the application.
type AppUpdateParamsTLS string

const (
	AppUpdateParamsTLSOff      AppUpdateParamsTLS = "off"
	AppUpdateParamsTLSFlexible AppUpdateParamsTLS = "flexible"
	AppUpdateParamsTLSFull     AppUpdateParamsTLS = "full"
	AppUpdateParamsTLSStrict   AppUpdateParamsTLS = "strict"
)

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppUpdateParamsTrafficType string

const (
	AppUpdateParamsTrafficTypeDirect AppUpdateParamsTrafficType = "direct"
	AppUpdateParamsTrafficTypeHTTP   AppUpdateParamsTrafficType = "http"
	AppUpdateParamsTrafficTypeHTTPS  AppUpdateParamsTrafficType = "https"
)

type AppUpdateResponseEnvelope struct {
	Errors   []AppUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AppUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AppUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    appUpdateResponseEnvelopeJSON    `json:"-"`
}

// appUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppUpdateResponseEnvelope]
type appUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppUpdateResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    appUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// appUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppUpdateResponseEnvelopeErrors]
type appUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppUpdateResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    appUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// appUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppUpdateResponseEnvelopeMessages]
type appUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppUpdateResponseEnvelopeSuccess bool

const (
	AppUpdateResponseEnvelopeSuccessTrue AppUpdateResponseEnvelopeSuccess = true
)

type AppListParams struct {
	// Sets the direction by which results are ordered.
	Direction param.Field[AppListParamsDirection] `query:"direction"`
	// Application field by which results are ordered.
	Order param.Field[AppListParamsOrder] `query:"order"`
	// Page number of paginated results. This parameter is required in order to use
	// other pagination parameters. If included in the query, `result_info` will be
	// present in the response.
	Page param.Field[float64] `query:"page"`
	// Sets the maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AppListParams]'s query parameters as `url.Values`.
func (r AppListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sets the direction by which results are ordered.
type AppListParamsDirection string

const (
	AppListParamsDirectionAsc  AppListParamsDirection = "asc"
	AppListParamsDirectionDesc AppListParamsDirection = "desc"
)

// Application field by which results are ordered.
type AppListParamsOrder string

const (
	AppListParamsOrderProtocol   AppListParamsOrder = "protocol"
	AppListParamsOrderAppID      AppListParamsOrder = "app_id"
	AppListParamsOrderCreatedOn  AppListParamsOrder = "created_on"
	AppListParamsOrderModifiedOn AppListParamsOrder = "modified_on"
	AppListParamsOrderDNS        AppListParamsOrder = "dns"
)

type AppDeleteResponseEnvelope struct {
	Errors   []AppDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AppDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AppDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    appDeleteResponseEnvelopeJSON    `json:"-"`
}

// appDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppDeleteResponseEnvelope]
type appDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    appDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// appDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppDeleteResponseEnvelopeErrors]
type appDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    appDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// appDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppDeleteResponseEnvelopeMessages]
type appDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppDeleteResponseEnvelopeSuccess bool

const (
	AppDeleteResponseEnvelopeSuccessTrue AppDeleteResponseEnvelopeSuccess = true
)

type AppGetResponseEnvelope struct {
	Errors   []AppGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AppGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AppGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AppGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    appGetResponseEnvelopeJSON    `json:"-"`
}

// appGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppGetResponseEnvelope]
type appGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppGetResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    appGetResponseEnvelopeErrorsJSON `json:"-"`
}

// appGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppGetResponseEnvelopeErrors]
type appGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppGetResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    appGetResponseEnvelopeMessagesJSON `json:"-"`
}

// appGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppGetResponseEnvelopeMessages]
type appGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AppGetResponseEnvelopeSuccess bool

const (
	AppGetResponseEnvelopeSuccessTrue AppGetResponseEnvelopeSuccess = true
)
