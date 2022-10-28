package qwhub_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/vikpe/go-qwhub"
	"github.com/vikpe/qw-hub-api/pkg/qtvscraper"
	"github.com/vikpe/qw-hub-api/pkg/twitch"
	"github.com/vikpe/serverstat/qserver/mvdsv"
)

func TestClient_MvdsvServers(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		hub := qwhub.NewClient()
		httpmock.ActivateNonDefault(hub.RestyClient.GetClient())
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://hubapi.quakeworld.nu/v2/servers/mvdsv", httpmock.NewStringResponder(http.StatusServiceUnavailable, ``))
		servers := hub.MvdsvServers()

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

			assert.Equal(t, servers, hub.MvdsvServers())
		})

		t.Run("with params", func(t *testing.T) {
			hub := qwhub.NewClient()
			httpmock.ActivateNonDefault(hub.RestyClient.GetClient())
			defer httpmock.DeactivateAndReset()
			httpmock.RegisterResponder("GET", "https://hubapi.quakeworld.nu/v2/servers/mvdsv?foo=1", httpmock.NewStringResponder(http.StatusOK, `[]`))

			hub.MvdsvServers(map[string]string{"foo": "1"})
			assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET https://hubapi.quakeworld.nu/v2/servers/mvdsv?foo=1"])
		})
	})
}

func TestClient_Streams(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		hub := qwhub.NewClient()
		httpmock.ActivateNonDefault(hub.RestyClient.GetClient())
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://hubapi.quakeworld.nu/v2/streams", httpmock.NewStringResponder(http.StatusServiceUnavailable, ``))
		streams := hub.Streams()

		assert.Equal(t, 1, httpmock.GetTotalCallCount())
		assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET https://hubapi.quakeworld.nu/v2/streams"])
		assert.Empty(t, streams)
	})

	t.Run("success", func(t *testing.T) {
		hub := qwhub.NewClient()
		httpmock.ActivateNonDefault(hub.RestyClient.GetClient())
		defer httpmock.DeactivateAndReset()

		streams := []twitch.Stream{
			{Title: "awesome stream 1"},
			{Title: "awesome stream 2"},
		}

		httpmock.RegisterResponder("GET", "https://hubapi.quakeworld.nu/v2/streams",
			func(req *http.Request) (*http.Response, error) {
				resp, _ := httpmock.NewJsonResponse(http.StatusOK, streams)
				return resp, nil
			},
		)

		assert.Equal(t, streams, hub.Streams())
	})
}

func TestClient_Demos(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		hub := qwhub.NewClient()
		httpmock.ActivateNonDefault(hub.RestyClient.GetClient())
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://hubapi.quakeworld.nu/v2/demos", httpmock.NewStringResponder(http.StatusServiceUnavailable, ``))
		streams := hub.Demos()

		assert.Equal(t, 1, httpmock.GetTotalCallCount())
		assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET https://hubapi.quakeworld.nu/v2/demos"])
		assert.Empty(t, streams)
	})

	t.Run("success", func(t *testing.T) {
		hub := qwhub.NewClient()
		httpmock.ActivateNonDefault(hub.RestyClient.GetClient())
		defer httpmock.DeactivateAndReset()

		demos := []qtvscraper.Demo{
			{
				QtvAddress:  "qw.foppa.dk:28000",
				Time:        time.Time{},
				Filename:    "2on2_foo_vs_bar[dm2]220101-1042",
				DownloadUrl: "http://qw.foppa.dk:28000/dl/demos/2on2_foo_vs_bar[dm2]220101-1042",
				QtvplayUrl:  "file:2on2_foo_vs_bar[dm2]220101-1042@qw.foppa.dk:28000",
			},
		}

		httpmock.RegisterResponder("GET", "https://hubapi.quakeworld.nu/v2/demos",
			func(req *http.Request) (*http.Response, error) {
				resp, _ := httpmock.NewJsonResponse(http.StatusOK, demos)
				return resp, nil
			},
		)

		assert.Equal(t, demos, hub.Demos())
	})

	t.Run("with params", func(t *testing.T) {
		hub := qwhub.NewClient()
		httpmock.ActivateNonDefault(hub.RestyClient.GetClient())
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "https://hubapi.quakeworld.nu/v2/demos?mode=2on2", httpmock.NewStringResponder(http.StatusOK, `[]`))

		hub.Demos(map[string]string{"mode": "2on2"})
		assert.Equal(t, 1, httpmock.GetCallCountInfo()["GET https://hubapi.quakeworld.nu/v2/demos?mode=2on2"])
	})
}
