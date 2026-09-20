package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// This is the test S11 L03 breaks on camera. Change the want string, push, and
// watch the workflow stop before the image is ever built — that is the whole
// point of gating the push on the test rather than beside it.
func TestBannerComesFromTheEnvironment(t *testing.T) {
	t.Setenv("BANNER", "storefront v1 — dev")
	if got, want := Banner(), "storefront v1 — dev"; got != want {
		t.Errorf("Banner() = %q, want %q", got, want)
	}
}

func TestBannerFallsBackWhenUnset(t *testing.T) {
	t.Setenv("BANNER", "")
	if got := Banner(); !strings.Contains(got, "no banner set") {
		t.Errorf("Banner() = %q, want the fallback", got)
	}
}

func TestHandlerServesTheBanner(t *testing.T) {
	t.Setenv("BANNER", "storefront v1 — prod")
	rec := httptest.NewRecorder()
	handler(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := strings.TrimSpace(rec.Body.String()); got != "storefront v1 — prod" {
		t.Errorf("body = %q, want %q", got, "storefront v1 — prod")
	}
}
