// File generated from our OpenAPI spec by Stainless.

package pagerules_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/internal/testutil"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/pagerules"
)

func TestPageruleNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Pagerules.New(context.TODO(), pagerules.PageruleNewParams{
		ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Actions: cloudflare.F([]pagerules.PageruleNewParamsAction{{
			Name: cloudflare.F(pagerules.PageruleNewParamsActionsNameForwardURL),
			Value: cloudflare.F(pagerules.PageruleNewParamsActionsValue{
				Type: cloudflare.F(pagerules.PageruleNewParamsActionsValueTypeTemporary),
				URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
			}),
		}}),
		Targets: cloudflare.F([]pagerules.PageruleNewParamsTarget{{
			Constraint: cloudflare.F(pagerules.PageruleNewParamsTargetsConstraint{
				Operator: cloudflare.F(pagerules.PageruleNewParamsTargetsConstraintOperatorMatches),
				Value:    cloudflare.F("*example.com/images/*"),
			}),
			Target: cloudflare.F(pagerules.PageruleNewParamsTargetsTargetURL),
		}}),
		Priority: cloudflare.F(int64(0)),
		Status:   cloudflare.F(pagerules.PageruleNewParamsStatusActive),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageruleUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Pagerules.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleUpdateParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Actions: cloudflare.F([]pagerules.PageruleUpdateParamsAction{{
				Name: cloudflare.F(pagerules.PageruleUpdateParamsActionsNameForwardURL),
				Value: cloudflare.F(pagerules.PageruleUpdateParamsActionsValue{
					Type: cloudflare.F(pagerules.PageruleUpdateParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}}),
			Targets: cloudflare.F([]pagerules.PageruleUpdateParamsTarget{{
				Constraint: cloudflare.F(pagerules.PageruleUpdateParamsTargetsConstraint{
					Operator: cloudflare.F(pagerules.PageruleUpdateParamsTargetsConstraintOperatorMatches),
					Value:    cloudflare.F("*example.com/images/*"),
				}),
				Target: cloudflare.F(pagerules.PageruleUpdateParamsTargetsTargetURL),
			}}),
			Priority: cloudflare.F(int64(0)),
			Status:   cloudflare.F(pagerules.PageruleUpdateParamsStatusActive),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageruleListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Pagerules.List(context.TODO(), pagerules.PageruleListParams{
		ZoneID:    cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		Direction: cloudflare.F(pagerules.PageruleListParamsDirectionDesc),
		Match:     cloudflare.F(pagerules.PageruleListParamsMatchAny),
		Order:     cloudflare.F(pagerules.PageruleListParamsOrderStatus),
		Status:    cloudflare.F(pagerules.PageruleListParamsStatusActive),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageruleDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Pagerules.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleDeleteParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageruleEditWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Pagerules.Edit(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleEditParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
			Actions: cloudflare.F([]pagerules.PageruleEditParamsAction{{
				Name: cloudflare.F(pagerules.PageruleEditParamsActionsNameForwardURL),
				Value: cloudflare.F(pagerules.PageruleEditParamsActionsValue{
					Type: cloudflare.F(pagerules.PageruleEditParamsActionsValueTypeTemporary),
					URL:  cloudflare.F("http://www.example.com/somewhere/$1/astring/$2/anotherstring/$3"),
				}),
			}}),
			Priority: cloudflare.F(int64(0)),
			Status:   cloudflare.F(pagerules.PageruleEditParamsStatusActive),
			Targets: cloudflare.F([]pagerules.PageruleEditParamsTarget{{
				Constraint: cloudflare.F(pagerules.PageruleEditParamsTargetsConstraint{
					Operator: cloudflare.F(pagerules.PageruleEditParamsTargetsConstraintOperatorMatches),
					Value:    cloudflare.F("*example.com/images/*"),
				}),
				Target: cloudflare.F(pagerules.PageruleEditParamsTargetsTargetURL),
			}}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageruleGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("user@example.com"),
	)
	_, err := client.Pagerules.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		pagerules.PageruleGetParams{
			ZoneID: cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
