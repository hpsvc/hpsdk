package hpsdk

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/hyperits/gosuite/httputil"
)

type Client struct {
	// 客户端ID
	AccessKey string `json:"accessKey"`
	// 客户端密钥
	Secret string `json:"secret"`
	// 接入点
	Endpoint string `json:"endpoint"`
	// HTTP 客户端
	httpClient *httputil.Client
	// 请求超时时间 单位秒
	RequestTimeout int `json:"requestTimeout"`
}

type ClientRequestOptions struct {
	// 请求方法
	Method string `json:"method"`
	// 请求 API URI
	URI string `json:"uri"`
	// 是否需要鉴权
	AuthReqired bool `json:"authReqired"`
	// 请求参数
	QueryParams map[string]string `json:"queryParams"`
	// 请求头
	Headers map[string]string `json:"headers"`
	// 请求体
	Body io.Reader `json:"body"`
}

type ClientRequestOption func(*ClientRequestOptions)

func WithMethod(method string) ClientRequestOption {
	return func(options *ClientRequestOptions) {
		options.Method = method
	}
}

func WithURI(uri string) ClientRequestOption {
	return func(options *ClientRequestOptions) {
		options.URI = uri
	}
}

func WithAuthReqired(auth bool) ClientRequestOption {
	return func(options *ClientRequestOptions) {
		options.AuthReqired = auth
	}
}

func WithQueryParams(params map[string]string) ClientRequestOption {
	return func(options *ClientRequestOptions) {
		options.QueryParams = params
	}
}

func WithHeaders(headers map[string]string) ClientRequestOption {
	return func(options *ClientRequestOptions) {
		options.Headers = headers
	}
}

func WithBody(body io.Reader) ClientRequestOption {
	return func(options *ClientRequestOptions) {
		options.Body = body
	}
}

func NewClientRequestOptions(options ...ClientRequestOption) *ClientRequestOptions {
	opts := &ClientRequestOptions{
		Method:      httputil.GET,
		URI:         "/",
		AuthReqired: false,
		QueryParams: make(map[string]string),
		Headers:     make(map[string]string),
		Body:        bytes.NewReader([]byte{}),
	}
	for _, option := range options {
		option(opts)
	}
	return opts
}

func NewClient(accessKey, secret string, endpoint string, requestTimeout int) (*Client, error) {
	_, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	return &Client{
		AccessKey:      accessKey,
		Secret:         secret,
		httpClient:     httputil.NewClient(),
		Endpoint:       endpoint,
		RequestTimeout: requestTimeout,
	}, nil
}

func (c *Client) DoRequest(options ...ClientRequestOption) (*httputil.Response, error) {
	opts := NewClientRequestOptions(options...)

	tm := time.Now().Unix()

	var args []string
	if opts.QueryParams != nil {
		for k, v := range opts.QueryParams {
			args = append(args, fmt.Sprintf("%s=%s", k, v))
		}
	}

	if opts.AuthReqired {
		sign := GenerateHMACSHA256Digest(c.AccessKey, c.Secret, tm)
		args = append(args, fmt.Sprintf("accessKey=%s", c.AccessKey))
		args = append(args, fmt.Sprintf("timestamp=%d", tm))
		args = append(args, fmt.Sprintf("sign=%s", sign))
	}

	url := ""
	if len(args) > 0 {
		url = fmt.Sprintf("%s%s?%s", c.Endpoint, opts.URI, strings.Join(args, "&"))
	} else {
		url = fmt.Sprintf("%s%s", c.Endpoint, opts.URI)
	}

	httpOpts := httputil.RequestOptions{
		Method:         opts.Method,
		URL:            url,
		Headers:        opts.Headers,
		Body:           opts.Body,
		RequestTimeout: time.Duration(c.RequestTimeout) * time.Second,
	}

	resp, err := c.httpClient.DoRequest(httpOpts)
	if err != nil {
		return nil, err
	}

	return c.httpClient.WrapHttpResponse(resp)
}
