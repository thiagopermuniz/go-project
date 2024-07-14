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

func (m *MockDataService) GetData(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func TestDataHandler(t *testing.T) {
	mockService := new(MockDataService)
	mockService.On("GetData", mock.Anything, "some-key").Return("mocked data", nil)

	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		data, err := mockService.GetData(ctx, "some-key")
		if err != nil {
			http.Error(w, "could not get data", http.StatusInternalServerError)
			return
		}
		response := DataResponse{Data: data}
		json.NewEncoder(w).Encode(response)
	}

	req, err := http.NewRequest("GET", "/data", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/data", handler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Expected status OK")

	expected := `{"data":"mocked data"}`
	assert.JSONEq(t, expected, rr.Body.String(), "Response body differs")

}
