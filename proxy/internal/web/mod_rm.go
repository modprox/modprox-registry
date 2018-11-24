package web

import (
	"fmt"
	"io"
	"net/http"

	"github.com/cactus/go-statsd-client/statsd"

	"github.com/modprox/mp/pkg/loggy"
	"github.com/modprox/mp/proxy/internal/modules/store"
)

type removeModule struct {
	index   store.Index
	store   store.ZipStore
	statter statsd.Statter
	log     loggy.Logger
}

func modRM(index store.Index, store store.ZipStore, statter statsd.Statter) http.Handler {
	return &removeModule{
		store:   store,
		statter: statter,
		index:   index,
		log:     loggy.New("mod-rm"),
	}
}

// e.g. POST http://localhost:9000/github.com/shoenig/toolkit/@v/v1.0.0.rm

func (h *removeModule) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mod, err := modInfoFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		_ = h.statter.Inc("mod-rm-bad-request", 1, 1)
		return
	}

	h.log.Infof("serving request for removal of %s", mod)

	if err := h.index.Remove(mod); err != nil {
		h.log.Errorf("failed to remove module from index %s: %v", mod, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// and remove from the store itself
	if err := h.store.DelZip(mod); err != nil {
		h.log.Errorf("failed to remove module from store %s: %v", mod, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("module %s removed", mod)
	w.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(w, msg); err != nil {
		h.log.Errorf("failed to write response: %v", err)
		return
	}
}
