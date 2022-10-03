package qwhub_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/vikpe/go-qwhub"
	"github.com/vikpe/serverstat/qserver/mvdsv"
)

func TestClient_GetMvdsvServers(t *testing.T) {

	t.Run("error", func(t *testing.T) {
		hub := qwhub.NewClient()
		httpmock.ActivateNonDefault(hub.RestyClient.GetClient())
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://hubapi.quakeworld.nu/v2/servers/mvdsv", httpmock.NewStringResponder(http.StatusServiceUnavailable, ``))
		servers := hub.GetMvdsvServers()

		assert.Equal(t, 1, httpmock.GetTotalCallCount())
		assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET https://hubapi.quakeworld.nu/v2/servers/mvdsv"])
		assert.Empty(t, servers)
	})

	t.Run("success", func(t *testing.T) {
		t.Run("no params", func(t *testing.T) {
			hub := qwhub.NewClient()
			httpmock.ActivateNonDefault(hub.RestyClient.GetClient())
			defer httpmock.DeactivateAndReset()

			servers := []mvdsv.Mvdsv{
				{Address: "qw.foppa.dk:28501"},
				{Address: "qw.foppa.dk:28502"},
				{Address: "qw.foppa.dk:28503"},
			}

			httpmock.RegisterResponder("GET", "https://hubapi.quakeworld.nu/v2/servers/mvdsv",
				func(req *http.Request) (*http.Response, error) {
					resp, _ := httpmock.NewJsonResponse(http.StatusOK, servers)
					return resp, nil
				},
			)

			assert.Equal(t, servers, hub.GetMvdsvServers())
		})

		t.Run("with params", func(t *testing.T) {
			hub := qwhub.NewClient()
			httpmock.ActivateNonDefault(hub.RestyClient.GetClient())
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder("GET", "https://hubapi.quakeworld.nu/v2/servers/mvdsv?foo=1", httpmock.NewStringResponder(http.StatusOK, `[]`))

			hub.GetMvdsvServers(map[string]string{"foo": "1"})
			assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET https://hubapi.quakeworld.nu/v2/servers/mvdsv?foo=1"])
		})
	})
}