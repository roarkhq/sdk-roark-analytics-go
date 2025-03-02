// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package roarkanalytics

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/roark-analytics-go/internal/apijson"
	"github.com/stainless-sdks/roark-analytics-go/internal/requestconfig"
	"github.com/stainless-sdks/roark-analytics-go/option"
)

// HealthService contains methods and other services that help with interacting
// with the roark API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	Options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r *HealthService) {
	r = &HealthService{}
	r.Options = opts
	return
}

// Returns the health status of the API and its dependencies
func (r *HealthService) Get(ctx context.Context, opts ...option.RequestOption) (res *HealthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HealthGetResponse struct {
	// Health check response payload
	Data HealthGetResponseData `json:"data,required"`
	JSON healthGetResponseJSON `json:"-"`
}

// healthGetResponseJSON contains the JSON metadata for the struct
// [HealthGetResponse]
type healthGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthGetResponseJSON) RawJSON() string {
	return r.raw
}

// Health check response payload
type HealthGetResponseData struct {
	Status    HealthGetResponseDataStatus `json:"status,required"`
	Timestamp string                      `json:"timestamp,required"`
	Version   string                      `json:"version,required"`
	JSON      healthGetResponseDataJSON   `json:"-"`
}

// healthGetResponseDataJSON contains the JSON metadata for the struct
// [HealthGetResponseData]
type healthGetResponseDataJSON struct {
	Status      apijson.Field
	Timestamp   apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type HealthGetResponseDataStatus string

const (
	HealthGetResponseDataStatusHealthy   HealthGetResponseDataStatus = "healthy"
	HealthGetResponseDataStatusDegraded  HealthGetResponseDataStatus = "degraded"
	HealthGetResponseDataStatusUnhealthy HealthGetResponseDataStatus = "unhealthy"
)

func (r HealthGetResponseDataStatus) IsKnown() bool {
	switch r {
	case HealthGetResponseDataStatusHealthy, HealthGetResponseDataStatusDegraded, HealthGetResponseDataStatusUnhealthy:
		return true
	}
	return false
}
