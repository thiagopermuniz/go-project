package repository

import (
	"fmt"
	"github.com/sony/gobreaker"
	"net/http"
	"time"
)

type Repository struct {
	cb        *gobreaker.CircuitBreaker
	failCount int
}

func NewRepository() *Repository {
	settings := gobreaker.Settings{
		Name:    "RepoCircuitBreaker",
		Timeout: 5 * time.Second, // shorten timeout for quicker testing
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 3
		},
	}
	cb := gobreaker.NewCircuitBreaker(settings)

	return &Repository{cb: cb}
}

func (r *Repository) fetchData() (string, error) {
	r.failCount++
	if r.failCount < 3 {
		return "", &httpError{statusCode: http.StatusInternalServerError, message: "internal server error"}
	} else if r.failCount == 3 {
		return "", &httpError{statusCode: http.StatusNoContent, message: "no content"}
	}
	return "data fetch succeeded", nil
}

type httpError struct {
	statusCode int
	message    string
}

func (e *httpError) Error() string {
	return e.message
}

func (r *Repository) GetData() (string, error) {
	result, err := r.cb.Execute(func() (interface{}, error) {
		fmt.Println("State:", r.cb.State().String())
		data, err := r.fetchData()
		if err != nil {
			if httpErr, ok := err.(*httpError); ok {
				switch httpErr.statusCode {
				case http.StatusInternalServerError:
					return nil, err
				case http.StatusNoContent:
					return "no content", nil
				}
			}
			return nil, err
		}
		return data, nil
	})

	if err != nil {
		return "", err
	}
	return result.(string), nil
}
