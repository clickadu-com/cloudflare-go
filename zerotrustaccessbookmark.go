// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustAccessBookmarkService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustAccessBookmarkService] method instead.
type ZeroTrustAccessBookmarkService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessBookmarkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustAccessBookmarkService(opts ...option.RequestOption) (r *ZeroTrustAccessBookmarkService) {
	r = &ZeroTrustAccessBookmarkService{}
	r.Options = opts
	return
}

// Create a new Bookmark application.
func (r *ZeroTrustAccessBookmarkService) New(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarks, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessBookmarkNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Bookmark application.
func (r *ZeroTrustAccessBookmarkService) Update(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarks, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessBookmarkUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Bookmark applications.
func (r *ZeroTrustAccessBookmarkService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *[]AccessBookmarks, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessBookmarkListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Bookmark application.
func (r *ZeroTrustAccessBookmarkService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *ZeroTrustAccessBookmarkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessBookmarkDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Bookmark application.
func (r *ZeroTrustAccessBookmarkService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccessBookmarks, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessBookmarkGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/access/bookmarks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessBookmarks struct {
	// The unique identifier for the Bookmark application.
	ID interface{} `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool      `json:"app_launcher_visible"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// The domain of the Bookmark application.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the Bookmark application.
	Name      string              `json:"name"`
	UpdatedAt time.Time           `json:"updated_at" format:"date-time"`
	JSON      accessBookmarksJSON `json:"-"`
}

// accessBookmarksJSON contains the JSON metadata for the struct [AccessBookmarks]
type accessBookmarksJSON struct {
	ID                 apijson.Field
	AppLauncherVisible apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessBookmarks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkDeleteResponse struct {
	// UUID
	ID   string                                    `json:"id"`
	JSON zeroTrustAccessBookmarkDeleteResponseJSON `json:"-"`
}

// zeroTrustAccessBookmarkDeleteResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessBookmarkDeleteResponse]
type zeroTrustAccessBookmarkDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkNewResponseEnvelope struct {
	Errors   []ZeroTrustAccessBookmarkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessBookmarkNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessBookmarks                                      `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessBookmarkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessBookmarkNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessBookmarkNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessBookmarkNewResponseEnvelope]
type zeroTrustAccessBookmarkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustAccessBookmarkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessBookmarkNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustAccessBookmarkNewResponseEnvelopeErrors]
type zeroTrustAccessBookmarkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustAccessBookmarkNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessBookmarkNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessBookmarkNewResponseEnvelopeMessages]
type zeroTrustAccessBookmarkNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessBookmarkNewResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessBookmarkNewResponseEnvelopeSuccessTrue ZeroTrustAccessBookmarkNewResponseEnvelopeSuccess = true
)

type ZeroTrustAccessBookmarkUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessBookmarkUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessBookmarkUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessBookmarks                                         `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessBookmarkUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessBookmarkUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessBookmarkUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessBookmarkUpdateResponseEnvelope]
type zeroTrustAccessBookmarkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkUpdateResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustAccessBookmarkUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessBookmarkUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessBookmarkUpdateResponseEnvelopeErrors]
type zeroTrustAccessBookmarkUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkUpdateResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessBookmarkUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessBookmarkUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessBookmarkUpdateResponseEnvelopeMessages]
type zeroTrustAccessBookmarkUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessBookmarkUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessBookmarkUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessBookmarkUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessBookmarkListResponseEnvelope struct {
	Errors   []ZeroTrustAccessBookmarkListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessBookmarkListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessBookmarks                                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessBookmarkListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessBookmarkListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessBookmarkListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessBookmarkListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessBookmarkListResponseEnvelope]
type zeroTrustAccessBookmarkListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zeroTrustAccessBookmarkListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessBookmarkListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustAccessBookmarkListResponseEnvelopeErrors]
type zeroTrustAccessBookmarkListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustAccessBookmarkListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessBookmarkListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessBookmarkListResponseEnvelopeMessages]
type zeroTrustAccessBookmarkListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessBookmarkListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessBookmarkListResponseEnvelopeSuccessTrue ZeroTrustAccessBookmarkListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessBookmarkListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       zeroTrustAccessBookmarkListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessBookmarkListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustAccessBookmarkListResponseEnvelopeResultInfo]
type zeroTrustAccessBookmarkListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkDeleteResponseEnvelope struct {
	Errors   []ZeroTrustAccessBookmarkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessBookmarkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessBookmarkDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessBookmarkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessBookmarkDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessBookmarkDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessBookmarkDeleteResponseEnvelope]
type zeroTrustAccessBookmarkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkDeleteResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustAccessBookmarkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessBookmarkDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessBookmarkDeleteResponseEnvelopeErrors]
type zeroTrustAccessBookmarkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkDeleteResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessBookmarkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessBookmarkDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessBookmarkDeleteResponseEnvelopeMessages]
type zeroTrustAccessBookmarkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessBookmarkDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessBookmarkDeleteResponseEnvelopeSuccessTrue ZeroTrustAccessBookmarkDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustAccessBookmarkGetResponseEnvelope struct {
	Errors   []ZeroTrustAccessBookmarkGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessBookmarkGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessBookmarks                                      `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessBookmarkGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessBookmarkGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessBookmarkGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessBookmarkGetResponseEnvelope]
type zeroTrustAccessBookmarkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustAccessBookmarkGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessBookmarkGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustAccessBookmarkGetResponseEnvelopeErrors]
type zeroTrustAccessBookmarkGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessBookmarkGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustAccessBookmarkGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessBookmarkGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessBookmarkGetResponseEnvelopeMessages]
type zeroTrustAccessBookmarkGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessBookmarkGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessBookmarkGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessBookmarkGetResponseEnvelopeSuccessTrue ZeroTrustAccessBookmarkGetResponseEnvelopeSuccess = true
)
