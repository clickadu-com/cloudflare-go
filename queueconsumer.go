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

// QueueConsumerService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewQueueConsumerService] method
// instead.
type QueueConsumerService struct {
	Options []option.RequestOption
}

// NewQueueConsumerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQueueConsumerService(opts ...option.RequestOption) (r *QueueConsumerService) {
	r = &QueueConsumerService{}
	r.Options = opts
	return
}

// Creates a new consumer for a queue.
func (r *QueueConsumerService) New(ctx context.Context, name string, params QueueConsumerNewParams, opts ...option.RequestOption) (res *WorkersConsumerCreated, err error) {
	opts = append(r.Options[:], opts...)
	var env QueueConsumerNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers", params.AccountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the consumer for a queue, or creates one if it does not exist.
func (r *QueueConsumerService) Update(ctx context.Context, name string, consumerName string, params QueueConsumerUpdateParams, opts ...option.RequestOption) (res *WorkersConsumerUpdated, err error) {
	opts = append(r.Options[:], opts...)
	var env QueueConsumerUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers/%s", params.AccountID, name, consumerName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes the consumer for a queue.
func (r *QueueConsumerService) Delete(ctx context.Context, name string, consumerName string, body QueueConsumerDeleteParams, opts ...option.RequestOption) (res *QueueConsumerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env QueueConsumerDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers/%s", body.AccountID, name, consumerName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the consumers for a queue.
func (r *QueueConsumerService) Get(ctx context.Context, name string, query QueueConsumerGetParams, opts ...option.RequestOption) (res *[]WorkersConsumer, err error) {
	opts = append(r.Options[:], opts...)
	var env QueueConsumerGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/queues/%s/consumers", query.AccountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkersConsumer struct {
	CreatedOn   interface{}             `json:"created_on"`
	Environment interface{}             `json:"environment"`
	QueueName   interface{}             `json:"queue_name"`
	Service     interface{}             `json:"service"`
	Settings    WorkersConsumerSettings `json:"settings"`
	JSON        workersConsumerJSON     `json:"-"`
}

// workersConsumerJSON contains the JSON metadata for the struct [WorkersConsumer]
type workersConsumerJSON struct {
	CreatedOn   apijson.Field
	Environment apijson.Field
	QueueName   apijson.Field
	Service     apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersConsumerSettings struct {
	BatchSize     float64                     `json:"batch_size"`
	MaxRetries    float64                     `json:"max_retries"`
	MaxWaitTimeMs float64                     `json:"max_wait_time_ms"`
	JSON          workersConsumerSettingsJSON `json:"-"`
}

// workersConsumerSettingsJSON contains the JSON metadata for the struct
// [WorkersConsumerSettings]
type workersConsumerSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersConsumerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersConsumerCreated struct {
	CreatedOn       interface{}                    `json:"created_on"`
	DeadLetterQueue string                         `json:"dead_letter_queue"`
	Environment     interface{}                    `json:"environment"`
	QueueName       interface{}                    `json:"queue_name"`
	ScriptName      interface{}                    `json:"script_name"`
	Settings        WorkersConsumerCreatedSettings `json:"settings"`
	JSON            workersConsumerCreatedJSON     `json:"-"`
}

// workersConsumerCreatedJSON contains the JSON metadata for the struct
// [WorkersConsumerCreated]
type workersConsumerCreatedJSON struct {
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	Environment     apijson.Field
	QueueName       apijson.Field
	ScriptName      apijson.Field
	Settings        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkersConsumerCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersConsumerCreatedSettings struct {
	BatchSize     float64                            `json:"batch_size"`
	MaxRetries    float64                            `json:"max_retries"`
	MaxWaitTimeMs float64                            `json:"max_wait_time_ms"`
	JSON          workersConsumerCreatedSettingsJSON `json:"-"`
}

// workersConsumerCreatedSettingsJSON contains the JSON metadata for the struct
// [WorkersConsumerCreatedSettings]
type workersConsumerCreatedSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersConsumerCreatedSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersConsumerUpdated struct {
	CreatedOn       interface{}                    `json:"created_on"`
	DeadLetterQueue interface{}                    `json:"dead_letter_queue"`
	Environment     interface{}                    `json:"environment"`
	QueueName       interface{}                    `json:"queue_name"`
	ScriptName      interface{}                    `json:"script_name"`
	Settings        WorkersConsumerUpdatedSettings `json:"settings"`
	JSON            workersConsumerUpdatedJSON     `json:"-"`
}

// workersConsumerUpdatedJSON contains the JSON metadata for the struct
// [WorkersConsumerUpdated]
type workersConsumerUpdatedJSON struct {
	CreatedOn       apijson.Field
	DeadLetterQueue apijson.Field
	Environment     apijson.Field
	QueueName       apijson.Field
	ScriptName      apijson.Field
	Settings        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkersConsumerUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersConsumerUpdatedSettings struct {
	BatchSize     float64                            `json:"batch_size"`
	MaxRetries    float64                            `json:"max_retries"`
	MaxWaitTimeMs float64                            `json:"max_wait_time_ms"`
	JSON          workersConsumerUpdatedSettingsJSON `json:"-"`
}

// workersConsumerUpdatedSettingsJSON contains the JSON metadata for the struct
// [WorkersConsumerUpdatedSettings]
type workersConsumerUpdatedSettingsJSON struct {
	BatchSize     apijson.Field
	MaxRetries    apijson.Field
	MaxWaitTimeMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersConsumerUpdatedSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [QueueConsumerDeleteResponseUnknown],
// [QueueConsumerDeleteResponseArray] or [shared.UnionString].
type QueueConsumerDeleteResponse interface {
	ImplementsQueueConsumerDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*QueueConsumerDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type QueueConsumerDeleteResponseArray []interface{}

func (r QueueConsumerDeleteResponseArray) ImplementsQueueConsumerDeleteResponse() {}

type QueueConsumerNewParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r QueueConsumerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type QueueConsumerNewResponseEnvelope struct {
	Errors   []QueueConsumerNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []QueueConsumerNewResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersConsumerCreated                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    QueueConsumerNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo QueueConsumerNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       queueConsumerNewResponseEnvelopeJSON       `json:"-"`
}

// queueConsumerNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [QueueConsumerNewResponseEnvelope]
type queueConsumerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type QueueConsumerNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    queueConsumerNewResponseEnvelopeErrorsJSON `json:"-"`
}

// queueConsumerNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [QueueConsumerNewResponseEnvelopeErrors]
type queueConsumerNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type QueueConsumerNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    queueConsumerNewResponseEnvelopeMessagesJSON `json:"-"`
}

// queueConsumerNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [QueueConsumerNewResponseEnvelopeMessages]
type queueConsumerNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type QueueConsumerNewResponseEnvelopeSuccess bool

const (
	QueueConsumerNewResponseEnvelopeSuccessTrue QueueConsumerNewResponseEnvelopeSuccess = true
)

type QueueConsumerNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       queueConsumerNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// queueConsumerNewResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [QueueConsumerNewResponseEnvelopeResultInfo]
type queueConsumerNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type QueueConsumerUpdateParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r QueueConsumerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type QueueConsumerUpdateResponseEnvelope struct {
	Errors   []QueueConsumerUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []QueueConsumerUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersConsumerUpdated                        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    QueueConsumerUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo QueueConsumerUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       queueConsumerUpdateResponseEnvelopeJSON       `json:"-"`
}

// queueConsumerUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [QueueConsumerUpdateResponseEnvelope]
type queueConsumerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type QueueConsumerUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    queueConsumerUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// queueConsumerUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [QueueConsumerUpdateResponseEnvelopeErrors]
type queueConsumerUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type QueueConsumerUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    queueConsumerUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// queueConsumerUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [QueueConsumerUpdateResponseEnvelopeMessages]
type queueConsumerUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type QueueConsumerUpdateResponseEnvelopeSuccess bool

const (
	QueueConsumerUpdateResponseEnvelopeSuccessTrue QueueConsumerUpdateResponseEnvelopeSuccess = true
)

type QueueConsumerUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       queueConsumerUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// queueConsumerUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [QueueConsumerUpdateResponseEnvelopeResultInfo]
type queueConsumerUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type QueueConsumerDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type QueueConsumerDeleteResponseEnvelope struct {
	Errors   []QueueConsumerDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []QueueConsumerDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   QueueConsumerDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    QueueConsumerDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo QueueConsumerDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       queueConsumerDeleteResponseEnvelopeJSON       `json:"-"`
}

// queueConsumerDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [QueueConsumerDeleteResponseEnvelope]
type queueConsumerDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type QueueConsumerDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    queueConsumerDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// queueConsumerDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [QueueConsumerDeleteResponseEnvelopeErrors]
type queueConsumerDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type QueueConsumerDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    queueConsumerDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// queueConsumerDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [QueueConsumerDeleteResponseEnvelopeMessages]
type queueConsumerDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type QueueConsumerDeleteResponseEnvelopeSuccess bool

const (
	QueueConsumerDeleteResponseEnvelopeSuccessTrue QueueConsumerDeleteResponseEnvelopeSuccess = true
)

type QueueConsumerDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       queueConsumerDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// queueConsumerDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [QueueConsumerDeleteResponseEnvelopeResultInfo]
type queueConsumerDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type QueueConsumerGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type QueueConsumerGetResponseEnvelope struct {
	Errors   []interface{}     `json:"errors,required,nullable"`
	Messages []interface{}     `json:"messages,required,nullable"`
	Result   []WorkersConsumer `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    QueueConsumerGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo QueueConsumerGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       queueConsumerGetResponseEnvelopeJSON       `json:"-"`
}

// queueConsumerGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [QueueConsumerGetResponseEnvelope]
type queueConsumerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type QueueConsumerGetResponseEnvelopeSuccess bool

const (
	QueueConsumerGetResponseEnvelopeSuccessTrue QueueConsumerGetResponseEnvelopeSuccess = true
)

type QueueConsumerGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	TotalPages interface{}                                    `json:"total_pages"`
	JSON       queueConsumerGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// queueConsumerGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [QueueConsumerGetResponseEnvelopeResultInfo]
type queueConsumerGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueueConsumerGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
