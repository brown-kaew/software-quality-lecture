package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMakeHttpCall(t *testing.T) {

	t.Run("Response should match expect", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id": 1, "name": "Kaew", "info": "Gopher"}`))
		}))
		defer server.Close()
		want := &Response{
			ID:   1,
			Name: "Kaew",
			Info: "Gopher",
		}

		r, err := MakeHttpCall(server.URL)

		if !errors.Is(err, nil) {
			t.Errorf("expected %v, but got %v", nil, err)
		}

		if !reflect.DeepEqual(want, r) {
			t.Errorf("expected %v, but got %v", want, r)
		}
	})
}
