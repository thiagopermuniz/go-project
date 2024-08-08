package httpcli

import (
	"context"
	"net/http"
)

type CustomClient struct {
	client  *http.Client
	headers map[string]string
}
type CustomOption func(u *CustomClient)

func (u *CustomClient) initOpts(opts ...CustomOption) func() {
	for _, opt := range opts {
		opt(u)
	}
	return func() {
		u.headers = make(map[string]string)
	}
}

func NewCustomClient(cc *CustomClient, opts ...CustomOption) *CustomClient {
	c := &CustomClient{
		client:  &http.Client{},
		headers: make(map[string]string),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (u *CustomClient) Get(ctx context.Context, ep string, opts ...CustomOption) (resp *http.Response, err error) {
	cls := u.initOpts(opts...)
	defer cls()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	req = u.prepReq(req)
	return u.client.Get("/user")
}

func (u *CustomClient) prepReq(r *http.Request) *http.Request {
	for k, v := range u.headers {
		r.Header.Set(k, v)
	}

	return r

}
