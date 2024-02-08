// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// FirewallWAFPackageRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallWAFPackageRuleService]
// method instead.
type FirewallWAFPackageRuleService struct {
	Options []option.RequestOption
}

// NewFirewallWAFPackageRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallWAFPackageRuleService(opts ...option.RequestOption) (r *FirewallWAFPackageRuleService) {
	r = &FirewallWAFPackageRuleService{}
	r.Options = opts
	return
}

// Fetches the details of a WAF rule in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageRuleService) Get(ctx context.Context, zoneID string, packageID string, ruleID string, opts ...option.RequestOption) (res *FirewallWAFPackageRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageRuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", zoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a WAF rule. You can only update the mode/action of the rule.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageRuleService) Update(ctx context.Context, zoneID string, packageID string, ruleID string, body FirewallWAFPackageRuleUpdateParams, opts ...option.RequestOption) (res *FirewallWAFPackageRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageRuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", zoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [FirewallWAFPackageRuleGetResponseUnknown],
// [FirewallWAFPackageRuleGetResponseArray] or [shared.UnionString].
type FirewallWAFPackageRuleGetResponse interface {
	ImplementsFirewallWAFPackageRuleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallWAFPackageRuleGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallWAFPackageRuleGetResponseArray []interface{}

func (r FirewallWAFPackageRuleGetResponseArray) ImplementsFirewallWAFPackageRuleGetResponse() {}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by
// [FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRule],
// [FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRule] or
// [FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRule].
type FirewallWAFPackageRuleUpdateResponse interface {
	implementsFirewallWAFPackageRuleUpdateResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallWAFPackageRuleUpdateResponse)(nil)).Elem(), "")
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleGroup `json:"group,required"`
	// When set to `on`, the current WAF rule will be used when evaluating the request.
	// Applies to anomaly detection WAF rules.
	Mode FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                             `json:"priority,required"`
	JSON     firewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleJSON contains the
// JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRule]
type firewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRule) implementsFirewallWAFPackageRuleUpdateResponse() {
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleAllowedMode string

const (
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleAllowedModeOn  FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleAllowedMode = "on"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleAllowedModeOff FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                  `json:"name"`
	JSON firewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleGroupJSON contains
// the JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleGroup]
type firewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current WAF rule will be used when evaluating the request.
// Applies to anomaly detection WAF rules.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleMode string

const (
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleModeOn  FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleMode = "on"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleModeOff FirewallWAFPackageRuleUpdateResponseWAFManagedRulesAnomalyRuleMode = "off"
)

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// The list of possible actions of the WAF rule when it is triggered.
	AllowedModes []FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// The default action/mode of a rule.
	DefaultMode FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleGroup `json:"group,required"`
	// The action that the current WAF rule will perform when triggered. Applies to
	// traditional (deny) WAF rules.
	Mode FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                     `json:"priority,required"`
	JSON     firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRule]
type firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	DefaultMode  apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRule) implementsFirewallWAFPackageRuleUpdateResponse() {
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedMode string

const (
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDefault   FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "default"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDisable   FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "disable"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate  FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "simulate"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedModeBlock     FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "block"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "challenge"
)

// The default action/mode of a rule.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                          `json:"name"`
	JSON firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleGroup]
type firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action that the current WAF rule will perform when triggered. Applies to
// traditional (deny) WAF rules.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleMode string

const (
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleModeDefault   FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleMode = "default"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleModeDisable   FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleMode = "disable"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleModeSimulate  FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleMode = "simulate"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleModeBlock     FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleMode = "block"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleModeChallenge FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalDenyRuleMode = "challenge"
)

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRule struct {
	// The unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// The public description of the WAF rule.
	Description string `json:"description,required"`
	// The rule group to which the current WAF rule belongs.
	Group FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleMode `json:"mode,required"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// The order in which the individual WAF rule is executed within its rule group.
	Priority string                                                                      `json:"priority,required"`
	JSON     firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRule]
type firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRule) implementsFirewallWAFPackageRuleUpdateResponse() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleAllowedMode string

const (
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOn  FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "on"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOff FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "off"
)

// The rule group to which the current WAF rule belongs.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleGroup struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The name of the rule group.
	Name string                                                                           `json:"name"`
	JSON firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleGroupJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleGroupJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleGroup]
type firewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleMode string

const (
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleModeOn  FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleMode = "on"
	FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleModeOff FirewallWAFPackageRuleUpdateResponseWAFManagedRulesTraditionalAllowRuleMode = "off"
)

type FirewallWAFPackageRuleGetResponseEnvelope struct {
	Errors   []FirewallWAFPackageRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFPackageRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallWAFPackageRuleGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FirewallWAFPackageRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFPackageRuleGetResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFPackageRuleGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallWAFPackageRuleGetResponseEnvelope]
type firewallWAFPackageRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    firewallWAFPackageRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFPackageRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FirewallWAFPackageRuleGetResponseEnvelopeErrors]
type firewallWAFPackageRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    firewallWAFPackageRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFPackageRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallWAFPackageRuleGetResponseEnvelopeMessages]
type firewallWAFPackageRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageRuleGetResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageRuleGetResponseEnvelopeSuccessTrue FirewallWAFPackageRuleGetResponseEnvelopeSuccess = true
)

type FirewallWAFPackageRuleUpdateParams struct {
	// The mode/action of the rule when triggered. You must use a value from the
	// `allowed_modes` array of the current rule.
	Mode param.Field[FirewallWAFPackageRuleUpdateParamsMode] `json:"mode"`
}

func (r FirewallWAFPackageRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The mode/action of the rule when triggered. You must use a value from the
// `allowed_modes` array of the current rule.
type FirewallWAFPackageRuleUpdateParamsMode string

const (
	FirewallWAFPackageRuleUpdateParamsModeDefault   FirewallWAFPackageRuleUpdateParamsMode = "default"
	FirewallWAFPackageRuleUpdateParamsModeDisable   FirewallWAFPackageRuleUpdateParamsMode = "disable"
	FirewallWAFPackageRuleUpdateParamsModeSimulate  FirewallWAFPackageRuleUpdateParamsMode = "simulate"
	FirewallWAFPackageRuleUpdateParamsModeBlock     FirewallWAFPackageRuleUpdateParamsMode = "block"
	FirewallWAFPackageRuleUpdateParamsModeChallenge FirewallWAFPackageRuleUpdateParamsMode = "challenge"
	FirewallWAFPackageRuleUpdateParamsModeOn        FirewallWAFPackageRuleUpdateParamsMode = "on"
	FirewallWAFPackageRuleUpdateParamsModeOff       FirewallWAFPackageRuleUpdateParamsMode = "off"
)

type FirewallWAFPackageRuleUpdateResponseEnvelope struct {
	Errors   []FirewallWAFPackageRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFPackageRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	// When triggered, anomaly detection WAF rules contribute to an overall threat
	// score that will determine if a request is considered malicious. You can
	// configure the total scoring threshold through the 'sensitivity' property of the
	// WAF package.
	Result FirewallWAFPackageRuleUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success FirewallWAFPackageRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFPackageRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [FirewallWAFPackageRuleUpdateResponseEnvelope]
type firewallWAFPackageRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    firewallWAFPackageRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [FirewallWAFPackageRuleUpdateResponseEnvelopeErrors]
type firewallWAFPackageRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    firewallWAFPackageRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFPackageRuleUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [FirewallWAFPackageRuleUpdateResponseEnvelopeMessages]
type firewallWAFPackageRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageRuleUpdateResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageRuleUpdateResponseEnvelopeSuccessTrue FirewallWAFPackageRuleUpdateResponseEnvelopeSuccess = true
)
