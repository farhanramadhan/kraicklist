package serve

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"challenge.haraj.com.sa/kraicklist/internal/services"
)

type routeHandler struct {
	searchSvc services.SearhServiceInterface
}

func NewRouteHandler(searchSvc services.SearhServiceInterface) http.HandlerFunc {
	rh := &routeHandler{
		searchSvc: searchSvc,
	}

	return rh.Search
}

func (rh routeHandler) Search(w http.ResponseWriter, r *http.Request) {
	// fetch query string from query params
	q := r.URL.Query().Get("q")
	if len(q) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing search query in query params"))
		return
	}
	// search relevant records
	records, err := rh.searchSvc.Search(context.TODO(), q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// output success response
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.Encode(records)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf.Bytes())
}
