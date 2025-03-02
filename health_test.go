// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package roarkanalytics_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/roark-analytics-go"
	"github.com/stainless-sdks/roark-analytics-go/internal/testutil"
	"github.com/stainless-sdks/roark-analytics-go/option"
)

func TestHealthGet(t *testing.T) {
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
	_, err := client.Health.Get(context.TODO())
	if err != nil {
		var apierr *roarkanalytics.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
