// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package roarkanalytics_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/roarkhq/sdk-roark-analytics-go"
	"github.com/roarkhq/sdk-roark-analytics-go/internal/testutil"
	"github.com/roarkhq/sdk-roark-analytics-go/option"
)

func TestCallAnalysisNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := roarkanalytics.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.CallAnalysis.New(context.TODO(), roarkanalytics.CallAnalysisNewParams{
		CallDirection: roarkanalytics.F(roarkanalytics.CallAnalysisNewParamsCallDirectionInbound),
		InterfaceType: roarkanalytics.F(roarkanalytics.CallAnalysisNewParamsInterfaceTypePhone),
		Participants: roarkanalytics.F([]roarkanalytics.CallAnalysisNewParamsParticipant{{
			Role:        roarkanalytics.F(roarkanalytics.CallAnalysisNewParamsParticipantsRoleAgent),
			IsSimulated: roarkanalytics.F(true),
			Name:        roarkanalytics.F("Sales Agent"),
			PhoneNumber: roarkanalytics.F("+15551234567"),
			SpokeFirst:  roarkanalytics.F(true),
		}, {
			Role:        roarkanalytics.F(roarkanalytics.CallAnalysisNewParamsParticipantsRoleAgent),
			IsSimulated: roarkanalytics.F(true),
			Name:        roarkanalytics.F("John Doe"),
			PhoneNumber: roarkanalytics.F("+15557654321"),
			SpokeFirst:  roarkanalytics.F(false),
		}}),
		RecordingURL:       roarkanalytics.F("https://example.com/recording.wav"),
		StartedAt:          roarkanalytics.F("2025-03-02T09:10:31.785Z"),
		EndedReason:        roarkanalytics.F("endedReason"),
		IsTest:             roarkanalytics.F(false),
		StereoRecordingURL: roarkanalytics.F("https://example.com/recording_stereo.wav"),
	})
	if err != nil {
		var apierr *roarkanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCallAnalysisGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := roarkanalytics.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.CallAnalysis.Get(context.TODO(), "jobId")
	if err != nil {
		var apierr *roarkanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
