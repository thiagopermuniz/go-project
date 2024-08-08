package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDataService struct {
	mock.Mock
}

func (m *MockDataService) GetServiceData(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func TestDataHandler(t *testing.T) {
	mockService := new(MockDataService)
	mockService.On("GetServiceData", mock.Anything, "some-key").Return(`{"key":"some-key", "value":"some value"}`, nil)

	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		dataStr, err := mockService.GetServiceData(ctx, "some-key")
		if err != nil {
			http.Error(w, "could not get data", http.StatusInternalServerError)
			return
		}

		data := map[string]any{}
		err = json.Unmarshal([]byte(dataStr), &data)
		if err != nil {
			http.Error(w, "invalid data format", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(data)
	}

	req, err := http.NewRequest("GET", "/data", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/data", handler)
	router.ServeHTTP(rr, req)

	data := map[string]any{}
	err = json.Unmarshal(rr.Body.Bytes(), &data)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")
	assert.JSONEq(t, `{"key":"some-key", "value":"some value"}`, rr.Body.String(), "Response body differs")
	assert.Equal(t, "some-key", data["key"], "Response differs")
	assert.Equal(t, "some value", data["value"], "Response differs")
}
