// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package roarkanalytics

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/roarkhq/sdk-roark-analytics-go/internal/apijson"
	"github.com/roarkhq/sdk-roark-analytics-go/internal/param"
	"github.com/roarkhq/sdk-roark-analytics-go/internal/requestconfig"
	"github.com/roarkhq/sdk-roark-analytics-go/option"
)

// CallAnalysisService contains methods and other services that help with
// interacting with the roark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallAnalysisService] method instead.
type CallAnalysisService struct {
	Options []option.RequestOption
}

// NewCallAnalysisService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCallAnalysisService(opts ...option.RequestOption) (r *CallAnalysisService) {
	r = &CallAnalysisService{}
	r.Options = opts
	return
}

// Upload a new call recording
func (r *CallAnalysisService) New(ctx context.Context, body CallAnalysisNewParams, opts ...option.RequestOption) (res *CallAnalysisNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/call-analysis"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a call analysis job by ID
func (r *CallAnalysisService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *CallAnalysisGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("v1/call-analysis/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CallAnalysisNewResponse struct {
	// Analysis job with associated call context
	Data CallAnalysisNewResponseData `json:"data,required"`
	JSON callAnalysisNewResponseJSON `json:"-"`
}

// callAnalysisNewResponseJSON contains the JSON metadata for the struct
// [CallAnalysisNewResponse]
type callAnalysisNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallAnalysisNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callAnalysisNewResponseJSON) RawJSON() string {
	return r.raw
}

// Analysis job with associated call context
type CallAnalysisNewResponseData struct {
	Call CallAnalysisNewResponseDataCall `json:"call,required"`
	// Analysis job ID for tracking progress
	JobID  string                            `json:"jobId,required" format:"uuid"`
	Status CallAnalysisNewResponseDataStatus `json:"status,required"`
	JSON   callAnalysisNewResponseDataJSON   `json:"-"`
}

// callAnalysisNewResponseDataJSON contains the JSON metadata for the struct
// [CallAnalysisNewResponseData]
type callAnalysisNewResponseDataJSON struct {
	Call        apijson.Field
	JobID       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallAnalysisNewResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callAnalysisNewResponseDataJSON) RawJSON() string {
	return r.raw
}

type CallAnalysisNewResponseDataCall struct {
	ID            string                                       `json:"id,required" format:"uuid"`
	CallDirection CallAnalysisNewResponseDataCallCallDirection `json:"callDirection,required"`
	IsTest        bool                                         `json:"isTest,required"`
	Participants  []CallAnalysisNewResponseDataCallParticipant `json:"participants,required"`
	StartedAt     string                                       `json:"startedAt,required"`
	Status        CallAnalysisNewResponseDataCallStatus        `json:"status,required,nullable"`
	DurationMs    float64                                      `json:"durationMs,nullable"`
	EndedAt       string                                       `json:"endedAt,nullable"`
	EndedReason   string                                       `json:"endedReason,nullable"`
	Summary       string                                       `json:"summary,nullable"`
	JSON          callAnalysisNewResponseDataCallJSON          `json:"-"`
}

// callAnalysisNewResponseDataCallJSON contains the JSON metadata for the struct
// [CallAnalysisNewResponseDataCall]
type callAnalysisNewResponseDataCallJSON struct {
	ID            apijson.Field
	CallDirection apijson.Field
	IsTest        apijson.Field
	Participants  apijson.Field
	StartedAt     apijson.Field
	Status        apijson.Field
	DurationMs    apijson.Field
	EndedAt       apijson.Field
	EndedReason   apijson.Field
	Summary       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CallAnalysisNewResponseDataCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callAnalysisNewResponseDataCallJSON) RawJSON() string {
	return r.raw
}

type CallAnalysisNewResponseDataCallCallDirection string

const (
	CallAnalysisNewResponseDataCallCallDirectionInbound  CallAnalysisNewResponseDataCallCallDirection = "INBOUND"
	CallAnalysisNewResponseDataCallCallDirectionOutbound CallAnalysisNewResponseDataCallCallDirection = "OUTBOUND"
)

func (r CallAnalysisNewResponseDataCallCallDirection) IsKnown() bool {
	switch r {
	case CallAnalysisNewResponseDataCallCallDirectionInbound, CallAnalysisNewResponseDataCallCallDirectionOutbound:
		return true
	}
	return false
}

type CallAnalysisNewResponseDataCallParticipant struct {
	Role        CallAnalysisNewResponseDataCallParticipantsRole `json:"role,required"`
	IsSimulated bool                                            `json:"isSimulated"`
	Name        string                                          `json:"name,nullable"`
	PhoneNumber string                                          `json:"phoneNumber,nullable"`
	SpokeFirst  bool                                            `json:"spokeFirst"`
	JSON        callAnalysisNewResponseDataCallParticipantJSON  `json:"-"`
}

// callAnalysisNewResponseDataCallParticipantJSON contains the JSON metadata for
// the struct [CallAnalysisNewResponseDataCallParticipant]
type callAnalysisNewResponseDataCallParticipantJSON struct {
	Role        apijson.Field
	IsSimulated apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	SpokeFirst  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallAnalysisNewResponseDataCallParticipant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callAnalysisNewResponseDataCallParticipantJSON) RawJSON() string {
	return r.raw
}

type CallAnalysisNewResponseDataCallParticipantsRole string

const (
	CallAnalysisNewResponseDataCallParticipantsRoleAgent    CallAnalysisNewResponseDataCallParticipantsRole = "AGENT"
	CallAnalysisNewResponseDataCallParticipantsRoleCustomer CallAnalysisNewResponseDataCallParticipantsRole = "CUSTOMER"
)

func (r CallAnalysisNewResponseDataCallParticipantsRole) IsKnown() bool {
	switch r {
	case CallAnalysisNewResponseDataCallParticipantsRoleAgent, CallAnalysisNewResponseDataCallParticipantsRoleCustomer:
		return true
	}
	return false
}

type CallAnalysisNewResponseDataCallStatus string

const (
	CallAnalysisNewResponseDataCallStatusRinging    CallAnalysisNewResponseDataCallStatus = "RINGING"
	CallAnalysisNewResponseDataCallStatusInProgress CallAnalysisNewResponseDataCallStatus = "IN_PROGRESS"
	CallAnalysisNewResponseDataCallStatusEnded      CallAnalysisNewResponseDataCallStatus = "ENDED"
)

func (r CallAnalysisNewResponseDataCallStatus) IsKnown() bool {
	switch r {
	case CallAnalysisNewResponseDataCallStatusRinging, CallAnalysisNewResponseDataCallStatusInProgress, CallAnalysisNewResponseDataCallStatusEnded:
		return true
	}
	return false
}

type CallAnalysisNewResponseDataStatus string

const (
	CallAnalysisNewResponseDataStatusPending    CallAnalysisNewResponseDataStatus = "PENDING"
	CallAnalysisNewResponseDataStatusInProgress CallAnalysisNewResponseDataStatus = "IN_PROGRESS"
	CallAnalysisNewResponseDataStatusCompleted  CallAnalysisNewResponseDataStatus = "COMPLETED"
	CallAnalysisNewResponseDataStatusFailed     CallAnalysisNewResponseDataStatus = "FAILED"
)

func (r CallAnalysisNewResponseDataStatus) IsKnown() bool {
	switch r {
	case CallAnalysisNewResponseDataStatusPending, CallAnalysisNewResponseDataStatusInProgress, CallAnalysisNewResponseDataStatusCompleted, CallAnalysisNewResponseDataStatusFailed:
		return true
	}
	return false
}

type CallAnalysisGetResponse struct {
	// Analysis job with associated call context
	Data CallAnalysisGetResponseData `json:"data,required"`
	JSON callAnalysisGetResponseJSON `json:"-"`
}

// callAnalysisGetResponseJSON contains the JSON metadata for the struct
// [CallAnalysisGetResponse]
type callAnalysisGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallAnalysisGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callAnalysisGetResponseJSON) RawJSON() string {
	return r.raw
}

// Analysis job with associated call context
type CallAnalysisGetResponseData struct {
	Call CallAnalysisGetResponseDataCall `json:"call,required"`
	// Analysis job ID for tracking progress
	JobID  string                            `json:"jobId,required" format:"uuid"`
	Status CallAnalysisGetResponseDataStatus `json:"status,required"`
	JSON   callAnalysisGetResponseDataJSON   `json:"-"`
}

// callAnalysisGetResponseDataJSON contains the JSON metadata for the struct
// [CallAnalysisGetResponseData]
type callAnalysisGetResponseDataJSON struct {
	Call        apijson.Field
	JobID       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallAnalysisGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callAnalysisGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type CallAnalysisGetResponseDataCall struct {
	ID            string                                       `json:"id,required" format:"uuid"`
	CallDirection CallAnalysisGetResponseDataCallCallDirection `json:"callDirection,required"`
	IsTest        bool                                         `json:"isTest,required"`
	Participants  []CallAnalysisGetResponseDataCallParticipant `json:"participants,required"`
	StartedAt     string                                       `json:"startedAt,required"`
	Status        CallAnalysisGetResponseDataCallStatus        `json:"status,required,nullable"`
	DurationMs    float64                                      `json:"durationMs,nullable"`
	EndedAt       string                                       `json:"endedAt,nullable"`
	EndedReason   string                                       `json:"endedReason,nullable"`
	Summary       string                                       `json:"summary,nullable"`
	JSON          callAnalysisGetResponseDataCallJSON          `json:"-"`
}

// callAnalysisGetResponseDataCallJSON contains the JSON metadata for the struct
// [CallAnalysisGetResponseDataCall]
type callAnalysisGetResponseDataCallJSON struct {
	ID            apijson.Field
	CallDirection apijson.Field
	IsTest        apijson.Field
	Participants  apijson.Field
	StartedAt     apijson.Field
	Status        apijson.Field
	DurationMs    apijson.Field
	EndedAt       apijson.Field
	EndedReason   apijson.Field
	Summary       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CallAnalysisGetResponseDataCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callAnalysisGetResponseDataCallJSON) RawJSON() string {
	return r.raw
}

type CallAnalysisGetResponseDataCallCallDirection string

const (
	CallAnalysisGetResponseDataCallCallDirectionInbound  CallAnalysisGetResponseDataCallCallDirection = "INBOUND"
	CallAnalysisGetResponseDataCallCallDirectionOutbound CallAnalysisGetResponseDataCallCallDirection = "OUTBOUND"
)

func (r CallAnalysisGetResponseDataCallCallDirection) IsKnown() bool {
	switch r {
	case CallAnalysisGetResponseDataCallCallDirectionInbound, CallAnalysisGetResponseDataCallCallDirectionOutbound:
		return true
	}
	return false
}

type CallAnalysisGetResponseDataCallParticipant struct {
	Role        CallAnalysisGetResponseDataCallParticipantsRole `json:"role,required"`
	IsSimulated bool                                            `json:"isSimulated"`
	Name        string                                          `json:"name,nullable"`
	PhoneNumber string                                          `json:"phoneNumber,nullable"`
	SpokeFirst  bool                                            `json:"spokeFirst"`
	JSON        callAnalysisGetResponseDataCallParticipantJSON  `json:"-"`
}

// callAnalysisGetResponseDataCallParticipantJSON contains the JSON metadata for
// the struct [CallAnalysisGetResponseDataCallParticipant]
type callAnalysisGetResponseDataCallParticipantJSON struct {
	Role        apijson.Field
	IsSimulated apijson.Field
	Name        apijson.Field
	PhoneNumber apijson.Field
	SpokeFirst  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallAnalysisGetResponseDataCallParticipant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callAnalysisGetResponseDataCallParticipantJSON) RawJSON() string {
	return r.raw
}

type CallAnalysisGetResponseDataCallParticipantsRole string

const (
	CallAnalysisGetResponseDataCallParticipantsRoleAgent    CallAnalysisGetResponseDataCallParticipantsRole = "AGENT"
	CallAnalysisGetResponseDataCallParticipantsRoleCustomer CallAnalysisGetResponseDataCallParticipantsRole = "CUSTOMER"
)

func (r CallAnalysisGetResponseDataCallParticipantsRole) IsKnown() bool {
	switch r {
	case CallAnalysisGetResponseDataCallParticipantsRoleAgent, CallAnalysisGetResponseDataCallParticipantsRoleCustomer:
		return true
	}
	return false
}

type CallAnalysisGetResponseDataCallStatus string

const (
	CallAnalysisGetResponseDataCallStatusRinging    CallAnalysisGetResponseDataCallStatus = "RINGING"
	CallAnalysisGetResponseDataCallStatusInProgress CallAnalysisGetResponseDataCallStatus = "IN_PROGRESS"
	CallAnalysisGetResponseDataCallStatusEnded      CallAnalysisGetResponseDataCallStatus = "ENDED"
)

func (r CallAnalysisGetResponseDataCallStatus) IsKnown() bool {
	switch r {
	case CallAnalysisGetResponseDataCallStatusRinging, CallAnalysisGetResponseDataCallStatusInProgress, CallAnalysisGetResponseDataCallStatusEnded:
		return true
	}
	return false
}

type CallAnalysisGetResponseDataStatus string

const (
	CallAnalysisGetResponseDataStatusPending    CallAnalysisGetResponseDataStatus = "PENDING"
	CallAnalysisGetResponseDataStatusInProgress CallAnalysisGetResponseDataStatus = "IN_PROGRESS"
	CallAnalysisGetResponseDataStatusCompleted  CallAnalysisGetResponseDataStatus = "COMPLETED"
	CallAnalysisGetResponseDataStatusFailed     CallAnalysisGetResponseDataStatus = "FAILED"
)

func (r CallAnalysisGetResponseDataStatus) IsKnown() bool {
	switch r {
	case CallAnalysisGetResponseDataStatusPending, CallAnalysisGetResponseDataStatusInProgress, CallAnalysisGetResponseDataStatusCompleted, CallAnalysisGetResponseDataStatusFailed:
		return true
	}
	return false
}

type CallAnalysisNewParams struct {
	// Direction of the call (INBOUND or OUTBOUND)
	CallDirection param.Field[CallAnalysisNewParamsCallDirection] `json:"callDirection,required"`
	// Interface type of the call (PHONE or WEB)
	InterfaceType param.Field[CallAnalysisNewParamsInterfaceType] `json:"interfaceType,required"`
	// Exactly two participants in the call
	Participants param.Field[[]CallAnalysisNewParamsParticipant] `json:"participants,required"`
	// URL of source recording (must be an accessible WAV file). Can be a signed URL.
	RecordingURL param.Field[string] `json:"recordingUrl,required" format:"uri"`
	// When the call started (ISO 8601 format)
	StartedAt param.Field[string] `json:"startedAt,required"`
	// Reason why the call ended (optional)
	EndedReason param.Field[string] `json:"endedReason"`
	// Whether this is a test call
	IsTest param.Field[bool] `json:"isTest"`
	// URL of source stereo recording in WAV format. Must be accessible. Can be a
	// signed URL. While optional it allows for a richer audio player
	StereoRecordingURL param.Field[string] `json:"stereoRecordingUrl" format:"uri"`
}

func (r CallAnalysisNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Direction of the call (INBOUND or OUTBOUND)
type CallAnalysisNewParamsCallDirection string

const (
	CallAnalysisNewParamsCallDirectionInbound  CallAnalysisNewParamsCallDirection = "INBOUND"
	CallAnalysisNewParamsCallDirectionOutbound CallAnalysisNewParamsCallDirection = "OUTBOUND"
)

func (r CallAnalysisNewParamsCallDirection) IsKnown() bool {
	switch r {
	case CallAnalysisNewParamsCallDirectionInbound, CallAnalysisNewParamsCallDirectionOutbound:
		return true
	}
	return false
}

// Interface type of the call (PHONE or WEB)
type CallAnalysisNewParamsInterfaceType string

const (
	CallAnalysisNewParamsInterfaceTypePhone CallAnalysisNewParamsInterfaceType = "PHONE"
	CallAnalysisNewParamsInterfaceTypeWeb   CallAnalysisNewParamsInterfaceType = "WEB"
)

func (r CallAnalysisNewParamsInterfaceType) IsKnown() bool {
	switch r {
	case CallAnalysisNewParamsInterfaceTypePhone, CallAnalysisNewParamsInterfaceTypeWeb:
		return true
	}
	return false
}

type CallAnalysisNewParamsParticipant struct {
	Role        param.Field[CallAnalysisNewParamsParticipantsRole] `json:"role,required"`
	IsSimulated param.Field[bool]                                  `json:"isSimulated"`
	Name        param.Field[string]                                `json:"name"`
	PhoneNumber param.Field[string]                                `json:"phoneNumber"`
	SpokeFirst  param.Field[bool]                                  `json:"spokeFirst"`
}

func (r CallAnalysisNewParamsParticipant) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CallAnalysisNewParamsParticipantsRole string

const (
	CallAnalysisNewParamsParticipantsRoleAgent    CallAnalysisNewParamsParticipantsRole = "AGENT"
	CallAnalysisNewParamsParticipantsRoleCustomer CallAnalysisNewParamsParticipantsRole = "CUSTOMER"
)

func (r CallAnalysisNewParamsParticipantsRole) IsKnown() bool {
	switch r {
	case CallAnalysisNewParamsParticipantsRoleAgent, CallAnalysisNewParamsParticipantsRoleCustomer:
		return true
	}
	return false
}
