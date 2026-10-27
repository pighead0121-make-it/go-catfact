package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
)

type APIServer struct {
	svc Service
}

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func makeHTTPFunc(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "requestID", rand.Intn(1000000))
		if err := fn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func NewAPIServer(svc Service) *APIServer {
	return &APIServer{
		svc: svc,
	}
}

func (s *APIServer) Start(listenAddr string) error {
	http.HandleFunc("/", makeHTTPFunc(s.handleGetCatFact))

	return http.ListenAndServe(listenAddr, nil)
}

func (s *APIServer) handleGetCatFact(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fact, err := s.svc.GetCatFact(ctx)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, fact)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
