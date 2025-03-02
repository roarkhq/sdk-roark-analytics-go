// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package roarkanalytics_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/roark-analytics-go"
	"github.com/stainless-sdks/roark-analytics-go/internal/testutil"
	"github.com/stainless-sdks/roark-analytics-go/option"
)

func TestUsage(t *testing.T) {
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
	callAnalysis, err := client.CallAnalysis.New(context.TODO(), roarkanalytics.CallAnalysisNewParams{
		CallDirection: roarkanalytics.F(roarkanalytics.CallAnalysisNewParamsCallDirectionInbound),
		InterfaceType: roarkanalytics.F(roarkanalytics.CallAnalysisNewParamsInterfaceTypePhone),
		Participants: roarkanalytics.F([]roarkanalytics.CallAnalysisNewParamsParticipant{{
			Role: roarkanalytics.F(roarkanalytics.CallAnalysisNewParamsParticipantsRoleAgent),
		}, {
			Role: roarkanalytics.F(roarkanalytics.CallAnalysisNewParamsParticipantsRoleAgent),
		}}),
		RecordingURL: roarkanalytics.F("https://example.com/recording.wav"),
		StartedAt:    roarkanalytics.F("2025-03-02T09:10:31.785Z"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", callAnalysis.Data)
}
