// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// TunnelService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewTunnelService] method instead.
type TunnelService struct {
	Options     []option.RequestOption
	Connections *TunnelConnectionService
}

// NewTunnelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTunnelService(opts ...option.RequestOption) (r *TunnelService) {
	r = &TunnelService{}
	r.Options = opts
	r.Connections = NewTunnelConnectionService(opts...)
	return
}

// Fetches a single Argo Tunnel.
func (r *TunnelService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *TunnelGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an Argo Tunnel from an account.
func (r *TunnelService) Delete(ctx context.Context, accountID string, tunnelID string, body TunnelDeleteParams, opts ...option.RequestOption) (res *TunnelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a new Argo Tunnel in an account.
func (r *TunnelService) ArgoTunnelNewAnArgoTunnel(ctx context.Context, accountID string, body TunnelArgoTunnelNewAnArgoTunnelParams, opts ...option.RequestOption) (res *TunnelArgoTunnelNewAnArgoTunnelResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelArgoTunnelNewAnArgoTunnelResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters all types of Tunnels in an account.
func (r *TunnelService) ArgoTunnelListArgoTunnels(ctx context.Context, accountID string, query TunnelArgoTunnelListArgoTunnelsParams, opts ...option.RequestOption) (res *[]TunnelArgoTunnelListArgoTunnelsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TunnelArgoTunnelListArgoTunnelsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/tunnels", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelGetResponse struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelGetResponseConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time             `json:"deleted_at,nullable" format:"date-time"`
	JSON      tunnelGetResponseJSON `json:"-"`
}

// tunnelGetResponseJSON contains the JSON metadata for the struct
// [TunnelGetResponse]
type tunnelGetResponseJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelGetResponseConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                          `json:"uuid"`
	JSON tunnelGetResponseConnectionJSON `json:"-"`
}

// tunnelGetResponseConnectionJSON contains the JSON metadata for the struct
// [TunnelGetResponseConnection]
type tunnelGetResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelGetResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeleteResponse struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelDeleteResponseConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                `json:"deleted_at,nullable" format:"date-time"`
	JSON      tunnelDeleteResponseJSON `json:"-"`
}

// tunnelDeleteResponseJSON contains the JSON metadata for the struct
// [TunnelDeleteResponse]
type tunnelDeleteResponseJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeleteResponseConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                             `json:"uuid"`
	JSON tunnelDeleteResponseConnectionJSON `json:"-"`
}

// tunnelDeleteResponseConnectionJSON contains the JSON metadata for the struct
// [TunnelDeleteResponseConnection]
type tunnelDeleteResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelDeleteResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelNewAnArgoTunnelResponse struct {
	// UUID of the tunnel.
	ID string `json:"id,required"`
	// The tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelArgoTunnelNewAnArgoTunnelResponseConnection `json:"connections,required"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// A user-friendly name for the tunnel.
	Name string `json:"name,required"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time                                   `json:"deleted_at,nullable" format:"date-time"`
	JSON      tunnelArgoTunnelNewAnArgoTunnelResponseJSON `json:"-"`
}

// tunnelArgoTunnelNewAnArgoTunnelResponseJSON contains the JSON metadata for the
// struct [TunnelArgoTunnelNewAnArgoTunnelResponse]
type tunnelArgoTunnelNewAnArgoTunnelResponseJSON struct {
	ID          apijson.Field
	Connections apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	DeletedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelNewAnArgoTunnelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelNewAnArgoTunnelResponseConnection struct {
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                `json:"uuid"`
	JSON tunnelArgoTunnelNewAnArgoTunnelResponseConnectionJSON `json:"-"`
}

// tunnelArgoTunnelNewAnArgoTunnelResponseConnectionJSON contains the JSON metadata
// for the struct [TunnelArgoTunnelNewAnArgoTunnelResponseConnection]
type tunnelArgoTunnelNewAnArgoTunnelResponseConnectionJSON struct {
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelArgoTunnelNewAnArgoTunnelResponseConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnel] or
// [TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnel].
type TunnelArgoTunnelListArgoTunnelsResponse interface {
	implementsTunnelArgoTunnelListArgoTunnelsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*TunnelArgoTunnelListArgoTunnelsResponse)(nil)).Elem(), "")
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunType `json:"tun_type"`
	JSON    tunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelJSON    `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelJSON contains the JSON
// metadata for the struct [TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnel]
type tunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnel) implementsTunnelArgoTunnelListArgoTunnelsResponse() {
}

type TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                               `json:"uuid"`
	JSON tunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelConnectionJSON `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelConnectionJSON contains
// the JSON metadata for the struct
// [TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelConnection]
type tunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunType string

const (
	TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunTypeCfdTunnel     TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunType = "cfd_tunnel"
	TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunTypeWarpConnector TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunType = "warp_connector"
	TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunTypeIPSec         TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunType = "ip_sec"
	TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunTypeGre           TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunType = "gre"
	TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunTypeCni           TunnelArgoTunnelListArgoTunnelsResponseTunnelCfdTunnelTunType = "cni"
)

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at,nullable" format:"date-time"`
	// Timestamp of when the tunnel was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the tunnel was deleted. If `null`, the tunnel has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for the tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status string `json:"status"`
	// The type of tunnel.
	TunType TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelJSON    `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelJSON contains
// the JSON metadata for the struct
// [TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnel]
type tunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnel) implementsTunnelArgoTunnelListArgoTunnelsResponse() {
}

type TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id"`
	// UUID of the cloudflared instance.
	ClientID interface{} `json:"client_id"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                                                                         `json:"uuid"`
	JSON tunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelConnectionJSON `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelConnectionJSON
// contains the JSON metadata for the struct
// [TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelConnection]
type tunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tunnel.
type TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunType string

const (
	TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunTypeCfdTunnel     TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunType = "cfd_tunnel"
	TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunTypeWarpConnector TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunType = "warp_connector"
	TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunTypeIPSec         TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunType = "ip_sec"
	TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunTypeGre           TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunType = "gre"
	TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunTypeCni           TunnelArgoTunnelListArgoTunnelsResponseTunnelWarpConnectorTunnelTunType = "cni"
)

type TunnelGetResponseEnvelope struct {
	Errors   []TunnelGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelGetResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TunnelGetResponseEnvelope]
type tunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    tunnelGetResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TunnelGetResponseEnvelopeErrors]
type tunnelGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    tunnelGetResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [TunnelGetResponseEnvelopeMessages]
type tunnelGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelGetResponseEnvelopeSuccess bool

const (
	TunnelGetResponseEnvelopeSuccessTrue TunnelGetResponseEnvelopeSuccess = true
)

type TunnelDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r TunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type TunnelDeleteResponseEnvelope struct {
	Errors   []TunnelDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// tunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TunnelDeleteResponseEnvelope]
type tunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeleteResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    tunnelDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [TunnelDeleteResponseEnvelopeErrors]
type tunnelDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelDeleteResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    tunnelDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TunnelDeleteResponseEnvelopeMessages]
type tunnelDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelDeleteResponseEnvelopeSuccess bool

const (
	TunnelDeleteResponseEnvelopeSuccessTrue TunnelDeleteResponseEnvelopeSuccess = true
)

type TunnelArgoTunnelNewAnArgoTunnelParams struct {
	// A user-friendly name for the tunnel.
	Name param.Field[string] `json:"name,required"`
	// Sets the password required to run the tunnel. Must be at least 32 bytes and
	// encoded as a base64 string.
	TunnelSecret param.Field[interface{}] `json:"tunnel_secret,required"`
}

func (r TunnelArgoTunnelNewAnArgoTunnelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelArgoTunnelNewAnArgoTunnelResponseEnvelope struct {
	Errors   []TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeMessages `json:"messages,required"`
	Result   TunnelArgoTunnelNewAnArgoTunnelResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeJSON    `json:"-"`
}

// tunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeJSON contains the JSON metadata
// for the struct [TunnelArgoTunnelNewAnArgoTunnelResponseEnvelope]
type tunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelNewAnArgoTunnelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    tunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeErrors]
type tunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    tunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeMessages]
type tunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeSuccess bool

const (
	TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeSuccessTrue TunnelArgoTunnelNewAnArgoTunnelResponseEnvelopeSuccess = true
)

type TunnelArgoTunnelListArgoTunnelsParams struct {
	ExcludePrefix param.Field[string] `query:"exclude_prefix"`
	// If provided, include only tunnels that were created (and not deleted) before
	// this time.
	ExistedAt     param.Field[time.Time] `query:"existed_at" format:"date-time"`
	IncludePrefix param.Field[string]    `query:"include_prefix"`
	// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If
	// empty, all tunnels will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// A user-friendly name for the tunnel.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// The types of tunnels to filter separated by a comma.
	TunTypes      param.Field[string]    `query:"tun_types"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [TunnelArgoTunnelListArgoTunnelsParams]'s query parameters
// as `url.Values`.
func (r TunnelArgoTunnelListArgoTunnelsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TunnelArgoTunnelListArgoTunnelsResponseEnvelope struct {
	Errors   []TunnelArgoTunnelListArgoTunnelsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelArgoTunnelListArgoTunnelsResponseEnvelopeMessages `json:"messages,required"`
	Result   []TunnelArgoTunnelListArgoTunnelsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    TunnelArgoTunnelListArgoTunnelsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo TunnelArgoTunnelListArgoTunnelsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       tunnelArgoTunnelListArgoTunnelsResponseEnvelopeJSON       `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseEnvelopeJSON contains the JSON metadata
// for the struct [TunnelArgoTunnelListArgoTunnelsResponseEnvelope]
type tunnelArgoTunnelListArgoTunnelsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelListArgoTunnelsResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    tunnelArgoTunnelListArgoTunnelsResponseEnvelopeErrorsJSON `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [TunnelArgoTunnelListArgoTunnelsResponseEnvelopeErrors]
type tunnelArgoTunnelListArgoTunnelsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelArgoTunnelListArgoTunnelsResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    tunnelArgoTunnelListArgoTunnelsResponseEnvelopeMessagesJSON `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [TunnelArgoTunnelListArgoTunnelsResponseEnvelopeMessages]
type tunnelArgoTunnelListArgoTunnelsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelArgoTunnelListArgoTunnelsResponseEnvelopeSuccess bool

const (
	TunnelArgoTunnelListArgoTunnelsResponseEnvelopeSuccessTrue TunnelArgoTunnelListArgoTunnelsResponseEnvelopeSuccess = true
)

type TunnelArgoTunnelListArgoTunnelsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       tunnelArgoTunnelListArgoTunnelsResponseEnvelopeResultInfoJSON `json:"-"`
}

// tunnelArgoTunnelListArgoTunnelsResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [TunnelArgoTunnelListArgoTunnelsResponseEnvelopeResultInfo]
type tunnelArgoTunnelListArgoTunnelsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelArgoTunnelListArgoTunnelsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
