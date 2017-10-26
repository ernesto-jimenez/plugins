// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// health client HTTP transport
//
// Command:
// $ goa gen goa.design/plugins/goakit/examples/client/fetcher/design

package client

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/http"
)

// Client lists the health service endpoint HTTP clients.
type Client struct {
	// Show Doer is the HTTP client used to make requests to the show endpoint.
	ShowDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the health service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ShowDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Show returns a endpoint that makes HTTP requests to the health service show
// server.
func (c *Client) Show() endpoint.Endpoint {
	var (
		decodeResponse = DecodeShowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildShowRequest(v)
		if err != nil {
			return nil, err
		}

		resp, err := c.ShowDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("health", "show", err)
		}
		return decodeResponse(resp)
	}
}