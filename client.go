package jike_sdk_go

import (
	"context"
	"errors"
	"log"

	"github.com/go-resty/resty/v2"
)

const (
	apiBaseURL = "https://api.ruguoapp.com"
	apiVersion = "/1.0"

	uploadBaseURL = "https://upload.ruguoapp.com"
	uploadVersion = "/1.0"

	qiNiuBaseURL = "https://upload.qiniup.com"
)

var (
	ErrUnknown = errors.New("unknown error")
)

func WithDebug(debug bool) OptFunc {
	return func(o *options) {
		o.debug = debug
	}
}

type OptFunc func(*options)

type options struct {
	debug bool
}

func NewClient(opts ...OptFunc) *Client {
	o := new(options)
	for _, f := range opts {
		f(o)
	}

	c := &Client{
		opts: o,
		rc:   resty.New(),
	}

	if c.opts.debug {
		c.rc.EnableTrace().
			OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
				log.Printf("url: %s\n", request.URL)
				return nil
			}).
			OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
				log.Printf("header: %v\n", response.Header())
				log.Printf("cookies: %v\n", response.Cookies())
				log.Printf("body: %s\n", response.Body())
				return nil
			})
	}
	return c
}

type Client struct {
	opts *options
	rc   *resty.Client

	// status info
	xAccessToken  string
	xRefreshToken string
	xRequestID    string
}

func (c *Client) apiR(ctx context.Context) *resty.Request {
	r := c.rc.R().
		SetContext(ctx).
		SetHeader("Host", mockAPIHost).
		SetHeader("x-jike-access-token", c.xAccessToken).
		SetHeader("x-jike-refresh-token", c.xRefreshToken).
		SetHeaders(mockAPIHeaders)
	return r
}

func (c *Client) uploadTokenR(ctx context.Context) *resty.Request {
	r := c.rc.R().
		SetContext(ctx).
		SetHeader("Host", mockUploadTokenHost).
		SetHeader("x-jike-access-token", c.xAccessToken).
		SetHeader("x-jike-refresh-token", c.xRefreshToken).
		SetHeaders(mockAPIHeaders)
	return r
}

func (c *Client) uploadR(ctx context.Context, ut string) *resty.Request {
	r := c.rc.R().
		SetContext(ctx).
		SetHeaders(mockUploadHeaders).
		SetHeader("User-Agent", mockUploadUserAgent(ut))
	return r
}
