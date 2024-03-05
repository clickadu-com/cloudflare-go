// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// LoadBalancerMonitorService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerMonitorService]
// method instead.
type LoadBalancerMonitorService struct {
	Options    []option.RequestOption
	Previews   *LoadBalancerMonitorPreviewService
	References *LoadBalancerMonitorReferenceService
}

// NewLoadBalancerMonitorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerMonitorService(opts ...option.RequestOption) (r *LoadBalancerMonitorService) {
	r = &LoadBalancerMonitorService{}
	r.Options = opts
	r.Previews = NewLoadBalancerMonitorPreviewService(opts...)
	r.References = NewLoadBalancerMonitorReferenceService(opts...)
	return
}

// Create a configured monitor.
func (r *LoadBalancerMonitorService) New(ctx context.Context, params LoadBalancerMonitorNewParams, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured monitor.
func (r *LoadBalancerMonitorService) Update(ctx context.Context, monitorID string, params LoadBalancerMonitorUpdateParams, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", params.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured monitors for an account.
func (r *LoadBalancerMonitorService) List(ctx context.Context, query LoadBalancerMonitorListParams, opts ...option.RequestOption) (res *[]LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured monitor.
func (r *LoadBalancerMonitorService) Delete(ctx context.Context, monitorID string, body LoadBalancerMonitorDeleteParams, opts ...option.RequestOption) (res *LoadBalancerMonitorDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", body.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing monitor, overwriting the supplied properties.
func (r *LoadBalancerMonitorService) Edit(ctx context.Context, monitorID string, params LoadBalancerMonitorEditParams, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", params.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a single configured monitor for an account.
func (r *LoadBalancerMonitorService) Get(ctx context.Context, monitorID string, query LoadBalancerMonitorGetParams, opts ...option.RequestOption) (res *LoadBalancingMonitor, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", query.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerMonitorDeleteResponse struct {
	ID   string                                `json:"id"`
	JSON loadBalancerMonitorDeleteResponseJSON `json:"-"`
}

// loadBalancerMonitorDeleteResponseJSON contains the JSON metadata for the struct
// [LoadBalancerMonitorDeleteResponse]
type loadBalancerMonitorDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
	Type param.Field[LoadBalancerMonitorNewParamsType] `json:"type"`
}

func (r LoadBalancerMonitorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorNewParamsType string

const (
	LoadBalancerMonitorNewParamsTypeHTTP     LoadBalancerMonitorNewParamsType = "http"
	LoadBalancerMonitorNewParamsTypeHTTPS    LoadBalancerMonitorNewParamsType = "https"
	LoadBalancerMonitorNewParamsTypeTcp      LoadBalancerMonitorNewParamsType = "tcp"
	LoadBalancerMonitorNewParamsTypeUdpIcmp  LoadBalancerMonitorNewParamsType = "udp_icmp"
	LoadBalancerMonitorNewParamsTypeIcmpPing LoadBalancerMonitorNewParamsType = "icmp_ping"
	LoadBalancerMonitorNewParamsTypeSmtp     LoadBalancerMonitorNewParamsType = "smtp"
)

type LoadBalancerMonitorNewResponseEnvelope struct {
	Errors   []LoadBalancerMonitorNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerMonitorNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancingMonitor                             `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerMonitorNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerMonitorNewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorNewResponseEnvelope]
type loadBalancerMonitorNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    loadBalancerMonitorNewResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerMonitorNewResponseEnvelopeErrors]
type loadBalancerMonitorNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    loadBalancerMonitorNewResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorNewResponseEnvelopeMessages]
type loadBalancerMonitorNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorNewResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorNewResponseEnvelopeSuccessTrue LoadBalancerMonitorNewResponseEnvelopeSuccess = true
)

type LoadBalancerMonitorUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
	Type param.Field[LoadBalancerMonitorUpdateParamsType] `json:"type"`
}

func (r LoadBalancerMonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorUpdateParamsType string

const (
	LoadBalancerMonitorUpdateParamsTypeHTTP     LoadBalancerMonitorUpdateParamsType = "http"
	LoadBalancerMonitorUpdateParamsTypeHTTPS    LoadBalancerMonitorUpdateParamsType = "https"
	LoadBalancerMonitorUpdateParamsTypeTcp      LoadBalancerMonitorUpdateParamsType = "tcp"
	LoadBalancerMonitorUpdateParamsTypeUdpIcmp  LoadBalancerMonitorUpdateParamsType = "udp_icmp"
	LoadBalancerMonitorUpdateParamsTypeIcmpPing LoadBalancerMonitorUpdateParamsType = "icmp_ping"
	LoadBalancerMonitorUpdateParamsTypeSmtp     LoadBalancerMonitorUpdateParamsType = "smtp"
)

type LoadBalancerMonitorUpdateResponseEnvelope struct {
	Errors   []LoadBalancerMonitorUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerMonitorUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancingMonitor                                `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerMonitorUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerMonitorUpdateResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorUpdateResponseEnvelope]
type loadBalancerMonitorUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    loadBalancerMonitorUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorUpdateResponseEnvelopeErrors]
type loadBalancerMonitorUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    loadBalancerMonitorUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorUpdateResponseEnvelopeMessages]
type loadBalancerMonitorUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorUpdateResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorUpdateResponseEnvelopeSuccessTrue LoadBalancerMonitorUpdateResponseEnvelopeSuccess = true
)

type LoadBalancerMonitorListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LoadBalancerMonitorListResponseEnvelope struct {
	Errors   []LoadBalancerMonitorListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerMonitorListResponseEnvelopeMessages `json:"messages,required"`
	Result   []LoadBalancingMonitor                            `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerMonitorListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerMonitorListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerMonitorListResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerMonitorListResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorListResponseEnvelope]
type loadBalancerMonitorListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    loadBalancerMonitorListResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerMonitorListResponseEnvelopeErrors]
type loadBalancerMonitorListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    loadBalancerMonitorListResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorListResponseEnvelopeMessages]
type loadBalancerMonitorListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorListResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorListResponseEnvelopeSuccessTrue LoadBalancerMonitorListResponseEnvelopeSuccess = true
)

type LoadBalancerMonitorListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       loadBalancerMonitorListResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerMonitorListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorListResponseEnvelopeResultInfo]
type loadBalancerMonitorListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LoadBalancerMonitorDeleteResponseEnvelope struct {
	Errors   []LoadBalancerMonitorDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerMonitorDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancerMonitorDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerMonitorDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerMonitorDeleteResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorDeleteResponseEnvelope]
type loadBalancerMonitorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    loadBalancerMonitorDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorDeleteResponseEnvelopeErrors]
type loadBalancerMonitorDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    loadBalancerMonitorDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorDeleteResponseEnvelopeMessages]
type loadBalancerMonitorDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorDeleteResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorDeleteResponseEnvelopeSuccessTrue LoadBalancerMonitorDeleteResponseEnvelopeSuccess = true
)

type LoadBalancerMonitorEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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
	Type param.Field[LoadBalancerMonitorEditParamsType] `json:"type"`
}

func (r LoadBalancerMonitorEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerMonitorEditParamsType string

const (
	LoadBalancerMonitorEditParamsTypeHTTP     LoadBalancerMonitorEditParamsType = "http"
	LoadBalancerMonitorEditParamsTypeHTTPS    LoadBalancerMonitorEditParamsType = "https"
	LoadBalancerMonitorEditParamsTypeTcp      LoadBalancerMonitorEditParamsType = "tcp"
	LoadBalancerMonitorEditParamsTypeUdpIcmp  LoadBalancerMonitorEditParamsType = "udp_icmp"
	LoadBalancerMonitorEditParamsTypeIcmpPing LoadBalancerMonitorEditParamsType = "icmp_ping"
	LoadBalancerMonitorEditParamsTypeSmtp     LoadBalancerMonitorEditParamsType = "smtp"
)

type LoadBalancerMonitorEditResponseEnvelope struct {
	Errors   []LoadBalancerMonitorEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerMonitorEditResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancingMonitor                              `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerMonitorEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerMonitorEditResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorEditResponseEnvelope]
type loadBalancerMonitorEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    loadBalancerMonitorEditResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerMonitorEditResponseEnvelopeErrors]
type loadBalancerMonitorEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    loadBalancerMonitorEditResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorEditResponseEnvelopeMessages]
type loadBalancerMonitorEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorEditResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorEditResponseEnvelopeSuccessTrue LoadBalancerMonitorEditResponseEnvelopeSuccess = true
)

type LoadBalancerMonitorGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LoadBalancerMonitorGetResponseEnvelope struct {
	Errors   []LoadBalancerMonitorGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerMonitorGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancingMonitor                             `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerMonitorGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerMonitorGetResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorGetResponseEnvelope]
type loadBalancerMonitorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    loadBalancerMonitorGetResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerMonitorGetResponseEnvelopeErrors]
type loadBalancerMonitorGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    loadBalancerMonitorGetResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorGetResponseEnvelopeMessages]
type loadBalancerMonitorGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorGetResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorGetResponseEnvelopeSuccessTrue LoadBalancerMonitorGetResponseEnvelopeSuccess = true
)
