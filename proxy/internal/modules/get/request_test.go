package get

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"oss.indeed.com/go/modprox/pkg/clients/registry"
	"oss.indeed.com/go/modprox/pkg/coordinates"
	"oss.indeed.com/go/modprox/pkg/netservice"
	"oss.indeed.com/go/modprox/pkg/webutil"
	"oss.indeed.com/go/modprox/proxy/internal/modules/store"
)

const modsReply = ` {"serials": [{
	"id": 2,
	"source": "github.com/pkg/errors",
	"version": "v0.8.0"
}]}`

func Test_ModulesNeeded(t *testing.T) {
	index := store.NewIndexMock(t)
	defer index.MinimockFinish()

	ids := coordinates.RangeIDs{
		coordinates.RangeID{1, 3},
		coordinates.RangeID{6, 6},
		coordinates.RangeID{10, 20},
	}

	index.IDsMock.Return(ids, nil)

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(modsReply))
		}),
	)
	defer ts.Close()

	address, port := webutil.ParseURL(t, ts.URL)
	client := registry.NewClient(registry.Options{
		Timeout: 10 * time.Second,
		Instances: []netservice.Instance{{
			Address: address,
			Port:    port,
		}},
	})

	apiClient := NewRegistryAPI(client, index)

	serialModules, err := apiClient.ModulesNeeded(ids)
	require.NoError(t, err)

	require.Equal(t, []coordinates.SerialModule{
		{
			SerialID: 2,
			Module: coordinates.Module{
				Source:  "github.com/pkg/errors",
				Version: "v0.8.0",
			},
		},
	}, serialModules)
}
