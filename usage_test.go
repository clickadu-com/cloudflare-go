// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-cloudflare-api-key"),
		option.WithEmail("dev@cloudflare.com"),
	)
	zoneNewResponse, err := client.Zones.New(context.TODO(), cloudflare.ZoneNewParams{
		Account: cloudflare.F(cloudflare.ZoneNewParamsAccount{
			ID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		}),
		Name: cloudflare.F("example.com"),
		Type: cloudflare.F(cloudflare.ZoneNewParamsTypePartial),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", zoneNewResponse.Result.ID)
}
