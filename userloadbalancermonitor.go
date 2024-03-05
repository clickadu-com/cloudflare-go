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

// UserLoadBalancerMonitorService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewUserLoadBalancerMonitorService] method instead.
type UserLoadBalancerMonitorService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerMonitorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserLoadBalancerMonitorService(opts ...option.RequestOption) (r *UserLoadBalancerMonitorService) {
	r = &UserLoadBalancerMonitorService{}
	r.Options = opts
	return
}

// Create a configured monitor.
func (r *UserLoadBalancerMonitorService) New(ctx context.Context, body UserLoadBalancerMonitorNewParams, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerMonitorNewResponseEnvelope
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured monitor.
func (r *UserLoadBalancerMonitorService) Update(ctx context.Context, monitorID string, body UserLoadBalancerMonitorUpdateParams, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerMonitorUpdateResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured monitors for a user.
func (r *UserLoadBalancerMonitorService) List(ctx context.Context, opts ...option.RequestOption) (res *[]LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerMonitorListResponseEnvelope
	path := "user/load_balancers/monitors"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured monitor.
func (r *UserLoadBalancerMonitorService) Delete(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *UserLoadBalancerMonitorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerMonitorDeleteResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing monitor, overwriting the supplied properties.
func (r *UserLoadBalancerMonitorService) Edit(ctx context.Context, monitorID string, body UserLoadBalancerMonitorEditParams, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerMonitorEditResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a single configured monitor for a user.
func (r *UserLoadBalancerMonitorService) Get(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerMonitorGetResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Preview pools using the specified monitor with provided monitor details. The
// returned preview_id can be used in the preview endpoint to retrieve the results.
func (r *UserLoadBalancerMonitorService) Preview(ctx context.Context, monitorID string, body UserLoadBalancerMonitorPreviewParams, opts ...option.RequestOption) (res *UserLoadBalancerMonitorPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerMonitorPreviewResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s/preview", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the list of resources that reference the provided monitor.
func (r *UserLoadBalancerMonitorService) References(ctx context.Context, monitorID string, opts ...option.RequestOption) (res *[]UserLoadBalancerMonitorReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerMonitorReferencesResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/monitors/%s/references", monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancingMonitor struct {
	ID string `json:"id"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure bool `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown int64 `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp int64     `json:"consecutive_up"`
	CreatedOn     time.Time `json:"created_on" format:"date-time"`
	// Object description.
	Description string `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes string `json:"expected_codes"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header interface{} `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval int64 `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method     string    `json:"method"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path string `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port int64 `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone string `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type LoadBalancingMonitorType `json:"type"`
	JSON loadBalancingMonitorJSON `json:"-"`
}

// loadBalancingMonitorJSON contains the JSON metadata for the struct
// [LoadBalancingMonitor]
type loadBalancingMonitorJSON struct {
	ID              apijson.Field
	AllowInsecure   apijson.Field
	ConsecutiveDown apijson.Field
	ConsecutiveUp   apijson.Field
	CreatedOn       apijson.Field
	Description     apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Interval        apijson.Field
	Method          apijson.Field
	ModifiedOn      apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	ProbeZone       apijson.Field
	Retries         apijson.Field
	Timeout         apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LoadBalancingMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancingMonitorType string

const (
	LoadBalancingMonitorTypeHTTP     LoadBalancingMonitorType = "http"
	LoadBalancingMonitorTypeHTTPS    LoadBalancingMonitorType = "https"
	LoadBalancingMonitorTypeTcp      LoadBalancingMonitorType = "tcp"
	LoadBalancingMonitorTypeUdpIcmp  LoadBalancingMonitorType = "udp_icmp"
	LoadBalancingMonitorTypeIcmpPing LoadBalancingMonitorType = "icmp_ping"
	LoadBalancingMonitorTypeSmtp     LoadBalancingMonitorType = "smtp"
)

type UserLoadBalancerMonitorDeleteResponse struct {
	ID   string                                    `json:"id"`
	JSON userLoadBalancerMonitorDeleteResponseJSON `json:"-"`
}

// userLoadBalancerMonitorDeleteResponseJSON contains the JSON metadata for the
// struct [UserLoadBalancerMonitorDeleteResponse]
type userLoadBalancerMonitorDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorPreviewResponse struct {
	// Monitored pool IDs mapped to their respective names.
	Pools     map[string]string                          `json:"pools"`
	PreviewID string                                     `json:"preview_id"`
	JSON      userLoadBalancerMonitorPreviewResponseJSON `json:"-"`
}

// userLoadBalancerMonitorPreviewResponseJSON contains the JSON metadata for the
// struct [UserLoadBalancerMonitorPreviewResponse]
type userLoadBalancerMonitorPreviewResponseJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorReferencesResponse struct {
	ReferenceType UserLoadBalancerMonitorReferencesResponseReferenceType `json:"reference_type"`
	ResourceID    string                                                 `json:"resource_id"`
	ResourceName  string                                                 `json:"resource_name"`
	ResourceType  string                                                 `json:"resource_type"`
	JSON          userLoadBalancerMonitorReferencesResponseJSON          `json:"-"`
}

// userLoadBalancerMonitorReferencesResponseJSON contains the JSON metadata for the
// struct [UserLoadBalancerMonitorReferencesResponse]
type userLoadBalancerMonitorReferencesResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorReferencesResponseReferenceType string

const (
	UserLoadBalancerMonitorReferencesResponseReferenceTypeStar     UserLoadBalancerMonitorReferencesResponseReferenceType = "*"
	UserLoadBalancerMonitorReferencesResponseReferenceTypeReferral UserLoadBalancerMonitorReferencesResponseReferenceType = "referral"
	UserLoadBalancerMonitorReferencesResponseReferenceTypeReferrer UserLoadBalancerMonitorReferencesResponseReferenceType = "referrer"
)

type UserLoadBalancerMonitorNewParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method param.Field[string] `json:"method"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path param.Field[string] `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone param.Field[string] `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[UserLoadBalancerMonitorNewParamsType] `json:"type"`
}

func (r UserLoadBalancerMonitorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerMonitorNewParamsType string

const (
	UserLoadBalancerMonitorNewParamsTypeHTTP     UserLoadBalancerMonitorNewParamsType = "http"
	UserLoadBalancerMonitorNewParamsTypeHTTPS    UserLoadBalancerMonitorNewParamsType = "https"
	UserLoadBalancerMonitorNewParamsTypeTcp      UserLoadBalancerMonitorNewParamsType = "tcp"
	UserLoadBalancerMonitorNewParamsTypeUdpIcmp  UserLoadBalancerMonitorNewParamsType = "udp_icmp"
	UserLoadBalancerMonitorNewParamsTypeIcmpPing UserLoadBalancerMonitorNewParamsType = "icmp_ping"
	UserLoadBalancerMonitorNewParamsTypeSmtp     UserLoadBalancerMonitorNewParamsType = "smtp"
)

type UserLoadBalancerMonitorNewResponseEnvelope struct {
	Errors   []UserLoadBalancerMonitorNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerMonitorNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancingMonitor                                 `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerMonitorNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerMonitorNewResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerMonitorNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserLoadBalancerMonitorNewResponseEnvelope]
type userLoadBalancerMonitorNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    userLoadBalancerMonitorNewResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerMonitorNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserLoadBalancerMonitorNewResponseEnvelopeErrors]
type userLoadBalancerMonitorNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    userLoadBalancerMonitorNewResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerMonitorNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerMonitorNewResponseEnvelopeMessages]
type userLoadBalancerMonitorNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorNewResponseEnvelopeSuccess bool

const (
	UserLoadBalancerMonitorNewResponseEnvelopeSuccessTrue UserLoadBalancerMonitorNewResponseEnvelopeSuccess = true
)

type UserLoadBalancerMonitorUpdateParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method param.Field[string] `json:"method"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path param.Field[string] `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone param.Field[string] `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[UserLoadBalancerMonitorUpdateParamsType] `json:"type"`
}

func (r UserLoadBalancerMonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerMonitorUpdateParamsType string

const (
	UserLoadBalancerMonitorUpdateParamsTypeHTTP     UserLoadBalancerMonitorUpdateParamsType = "http"
	UserLoadBalancerMonitorUpdateParamsTypeHTTPS    UserLoadBalancerMonitorUpdateParamsType = "https"
	UserLoadBalancerMonitorUpdateParamsTypeTcp      UserLoadBalancerMonitorUpdateParamsType = "tcp"
	UserLoadBalancerMonitorUpdateParamsTypeUdpIcmp  UserLoadBalancerMonitorUpdateParamsType = "udp_icmp"
	UserLoadBalancerMonitorUpdateParamsTypeIcmpPing UserLoadBalancerMonitorUpdateParamsType = "icmp_ping"
	UserLoadBalancerMonitorUpdateParamsTypeSmtp     UserLoadBalancerMonitorUpdateParamsType = "smtp"
)

type UserLoadBalancerMonitorUpdateResponseEnvelope struct {
	Errors   []UserLoadBalancerMonitorUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerMonitorUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancingMonitor                                    `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerMonitorUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerMonitorUpdateResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerMonitorUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserLoadBalancerMonitorUpdateResponseEnvelope]
type userLoadBalancerMonitorUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorUpdateResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    userLoadBalancerMonitorUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerMonitorUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserLoadBalancerMonitorUpdateResponseEnvelopeErrors]
type userLoadBalancerMonitorUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorUpdateResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    userLoadBalancerMonitorUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerMonitorUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerMonitorUpdateResponseEnvelopeMessages]
type userLoadBalancerMonitorUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorUpdateResponseEnvelopeSuccess bool

const (
	UserLoadBalancerMonitorUpdateResponseEnvelopeSuccessTrue UserLoadBalancerMonitorUpdateResponseEnvelopeSuccess = true
)

type UserLoadBalancerMonitorListResponseEnvelope struct {
	Errors   []UserLoadBalancerMonitorListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerMonitorListResponseEnvelopeMessages `json:"messages,required"`
	Result   []LoadBalancingMonitor                                `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserLoadBalancerMonitorListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserLoadBalancerMonitorListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userLoadBalancerMonitorListResponseEnvelopeJSON       `json:"-"`
}

// userLoadBalancerMonitorListResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserLoadBalancerMonitorListResponseEnvelope]
type userLoadBalancerMonitorListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    userLoadBalancerMonitorListResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerMonitorListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserLoadBalancerMonitorListResponseEnvelopeErrors]
type userLoadBalancerMonitorListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    userLoadBalancerMonitorListResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerMonitorListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerMonitorListResponseEnvelopeMessages]
type userLoadBalancerMonitorListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorListResponseEnvelopeSuccess bool

const (
	UserLoadBalancerMonitorListResponseEnvelopeSuccessTrue UserLoadBalancerMonitorListResponseEnvelopeSuccess = true
)

type UserLoadBalancerMonitorListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       userLoadBalancerMonitorListResponseEnvelopeResultInfoJSON `json:"-"`
}

// userLoadBalancerMonitorListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [UserLoadBalancerMonitorListResponseEnvelopeResultInfo]
type userLoadBalancerMonitorListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorDeleteResponseEnvelope struct {
	Errors   []UserLoadBalancerMonitorDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerMonitorDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   UserLoadBalancerMonitorDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerMonitorDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerMonitorDeleteResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerMonitorDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserLoadBalancerMonitorDeleteResponseEnvelope]
type userLoadBalancerMonitorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorDeleteResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    userLoadBalancerMonitorDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerMonitorDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserLoadBalancerMonitorDeleteResponseEnvelopeErrors]
type userLoadBalancerMonitorDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorDeleteResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    userLoadBalancerMonitorDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerMonitorDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerMonitorDeleteResponseEnvelopeMessages]
type userLoadBalancerMonitorDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorDeleteResponseEnvelopeSuccess bool

const (
	UserLoadBalancerMonitorDeleteResponseEnvelopeSuccessTrue UserLoadBalancerMonitorDeleteResponseEnvelopeSuccess = true
)

type UserLoadBalancerMonitorEditParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method param.Field[string] `json:"method"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path param.Field[string] `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone param.Field[string] `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[UserLoadBalancerMonitorEditParamsType] `json:"type"`
}

func (r UserLoadBalancerMonitorEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerMonitorEditParamsType string

const (
	UserLoadBalancerMonitorEditParamsTypeHTTP     UserLoadBalancerMonitorEditParamsType = "http"
	UserLoadBalancerMonitorEditParamsTypeHTTPS    UserLoadBalancerMonitorEditParamsType = "https"
	UserLoadBalancerMonitorEditParamsTypeTcp      UserLoadBalancerMonitorEditParamsType = "tcp"
	UserLoadBalancerMonitorEditParamsTypeUdpIcmp  UserLoadBalancerMonitorEditParamsType = "udp_icmp"
	UserLoadBalancerMonitorEditParamsTypeIcmpPing UserLoadBalancerMonitorEditParamsType = "icmp_ping"
	UserLoadBalancerMonitorEditParamsTypeSmtp     UserLoadBalancerMonitorEditParamsType = "smtp"
)

type UserLoadBalancerMonitorEditResponseEnvelope struct {
	Errors   []UserLoadBalancerMonitorEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerMonitorEditResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancingMonitor                                  `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerMonitorEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerMonitorEditResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerMonitorEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserLoadBalancerMonitorEditResponseEnvelope]
type userLoadBalancerMonitorEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorEditResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    userLoadBalancerMonitorEditResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerMonitorEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserLoadBalancerMonitorEditResponseEnvelopeErrors]
type userLoadBalancerMonitorEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorEditResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    userLoadBalancerMonitorEditResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerMonitorEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerMonitorEditResponseEnvelopeMessages]
type userLoadBalancerMonitorEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorEditResponseEnvelopeSuccess bool

const (
	UserLoadBalancerMonitorEditResponseEnvelopeSuccessTrue UserLoadBalancerMonitorEditResponseEnvelopeSuccess = true
)

type UserLoadBalancerMonitorGetResponseEnvelope struct {
	Errors   []UserLoadBalancerMonitorGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerMonitorGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancingMonitor                                 `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerMonitorGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerMonitorGetResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerMonitorGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserLoadBalancerMonitorGetResponseEnvelope]
type userLoadBalancerMonitorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    userLoadBalancerMonitorGetResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerMonitorGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserLoadBalancerMonitorGetResponseEnvelopeErrors]
type userLoadBalancerMonitorGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    userLoadBalancerMonitorGetResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerMonitorGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerMonitorGetResponseEnvelopeMessages]
type userLoadBalancerMonitorGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorGetResponseEnvelopeSuccess bool

const (
	UserLoadBalancerMonitorGetResponseEnvelopeSuccessTrue UserLoadBalancerMonitorGetResponseEnvelopeSuccess = true
)

type UserLoadBalancerMonitorPreviewParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method param.Field[string] `json:"method"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path param.Field[string] `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone param.Field[string] `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[UserLoadBalancerMonitorPreviewParamsType] `json:"type"`
}

func (r UserLoadBalancerMonitorPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerMonitorPreviewParamsType string

const (
	UserLoadBalancerMonitorPreviewParamsTypeHTTP     UserLoadBalancerMonitorPreviewParamsType = "http"
	UserLoadBalancerMonitorPreviewParamsTypeHTTPS    UserLoadBalancerMonitorPreviewParamsType = "https"
	UserLoadBalancerMonitorPreviewParamsTypeTcp      UserLoadBalancerMonitorPreviewParamsType = "tcp"
	UserLoadBalancerMonitorPreviewParamsTypeUdpIcmp  UserLoadBalancerMonitorPreviewParamsType = "udp_icmp"
	UserLoadBalancerMonitorPreviewParamsTypeIcmpPing UserLoadBalancerMonitorPreviewParamsType = "icmp_ping"
	UserLoadBalancerMonitorPreviewParamsTypeSmtp     UserLoadBalancerMonitorPreviewParamsType = "smtp"
)

type UserLoadBalancerMonitorPreviewResponseEnvelope struct {
	Errors   []UserLoadBalancerMonitorPreviewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerMonitorPreviewResponseEnvelopeMessages `json:"messages,required"`
	Result   UserLoadBalancerMonitorPreviewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerMonitorPreviewResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerMonitorPreviewResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerMonitorPreviewResponseEnvelopeJSON contains the JSON metadata
// for the struct [UserLoadBalancerMonitorPreviewResponseEnvelope]
type userLoadBalancerMonitorPreviewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorPreviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorPreviewResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    userLoadBalancerMonitorPreviewResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerMonitorPreviewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserLoadBalancerMonitorPreviewResponseEnvelopeErrors]
type userLoadBalancerMonitorPreviewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorPreviewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorPreviewResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    userLoadBalancerMonitorPreviewResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerMonitorPreviewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerMonitorPreviewResponseEnvelopeMessages]
type userLoadBalancerMonitorPreviewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorPreviewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorPreviewResponseEnvelopeSuccess bool

const (
	UserLoadBalancerMonitorPreviewResponseEnvelopeSuccessTrue UserLoadBalancerMonitorPreviewResponseEnvelopeSuccess = true
)

type UserLoadBalancerMonitorReferencesResponseEnvelope struct {
	Errors   []UserLoadBalancerMonitorReferencesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerMonitorReferencesResponseEnvelopeMessages `json:"messages,required"`
	// List of resources that reference a given monitor.
	Result []UserLoadBalancerMonitorReferencesResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserLoadBalancerMonitorReferencesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserLoadBalancerMonitorReferencesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userLoadBalancerMonitorReferencesResponseEnvelopeJSON       `json:"-"`
}

// userLoadBalancerMonitorReferencesResponseEnvelopeJSON contains the JSON metadata
// for the struct [UserLoadBalancerMonitorReferencesResponseEnvelope]
type userLoadBalancerMonitorReferencesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferencesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorReferencesResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    userLoadBalancerMonitorReferencesResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerMonitorReferencesResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerMonitorReferencesResponseEnvelopeErrors]
type userLoadBalancerMonitorReferencesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferencesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerMonitorReferencesResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    userLoadBalancerMonitorReferencesResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerMonitorReferencesResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerMonitorReferencesResponseEnvelopeMessages]
type userLoadBalancerMonitorReferencesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferencesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerMonitorReferencesResponseEnvelopeSuccess bool

const (
	UserLoadBalancerMonitorReferencesResponseEnvelopeSuccessTrue UserLoadBalancerMonitorReferencesResponseEnvelopeSuccess = true
)

type UserLoadBalancerMonitorReferencesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       userLoadBalancerMonitorReferencesResponseEnvelopeResultInfoJSON `json:"-"`
}

// userLoadBalancerMonitorReferencesResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerMonitorReferencesResponseEnvelopeResultInfo]
type userLoadBalancerMonitorReferencesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerMonitorReferencesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
