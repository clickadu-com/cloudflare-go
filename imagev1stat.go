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

// ImageV1StatService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewImageV1StatService] method
// instead.
type ImageV1StatService struct {
	Options []option.RequestOption
}

// NewImageV1StatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewImageV1StatService(opts ...option.RequestOption) (r *ImageV1StatService) {
	r = &ImageV1StatService{}
	r.Options = opts
	return
}

// Fetch usage statistics details for Cloudflare Images.
func (r *ImageV1StatService) Get(ctx context.Context, query ImageV1StatGetParams, opts ...option.RequestOption) (res *ImagesImagesStats, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1StatGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/stats", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImagesImagesStats struct {
	Count ImagesImagesStatsCount `json:"count"`
	JSON  imagesImagesStatsJSON  `json:"-"`
}

// imagesImagesStatsJSON contains the JSON metadata for the struct
// [ImagesImagesStats]
type imagesImagesStatsJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagesImagesStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImagesImagesStatsCount struct {
	// Cloudflare Images allowed usage.
	Allowed float64 `json:"allowed"`
	// Cloudflare Images current usage.
	Current float64                    `json:"current"`
	JSON    imagesImagesStatsCountJSON `json:"-"`
}

// imagesImagesStatsCountJSON contains the JSON metadata for the struct
// [ImagesImagesStatsCount]
type imagesImagesStatsCountJSON struct {
	Allowed     apijson.Field
	Current     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagesImagesStatsCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1StatGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ImageV1StatGetResponseEnvelope struct {
	Errors   []ImageV1StatGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1StatGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ImagesImagesStats                        `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1StatGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1StatGetResponseEnvelopeJSON    `json:"-"`
}

// imageV1StatGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ImageV1StatGetResponseEnvelope]
type imageV1StatGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1StatGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1StatGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    imageV1StatGetResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1StatGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ImageV1StatGetResponseEnvelopeErrors]
type imageV1StatGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1StatGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1StatGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    imageV1StatGetResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1StatGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ImageV1StatGetResponseEnvelopeMessages]
type imageV1StatGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1StatGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1StatGetResponseEnvelopeSuccess bool

const (
	ImageV1StatGetResponseEnvelopeSuccessTrue ImageV1StatGetResponseEnvelopeSuccess = true
)
