package hpsdk

import "github.com/hyperits/gosuite/httputil"

type Client struct {
	// 客户端ID
	AccessKey string `json:"access_key"`
	// 客户端密钥
	Secret string `json:"secret"`
	// 接入点
	Endpoint string `json:"endpoint"`

	httpClient *httputil.Client
}

type ClientRequestOptions struct {
	// 请求方法
	Method string `json:"method"`
	// 请求 API URI
	URI string `json:"uri"`
	// 是否需要鉴权
	AuthReqired bool `json:"auth_reqired"`
	// 请求参数
	Params map[string]interface{} `json:"params"`
	// 请求头
	Headers map[string]interface{} `json:"headers"`
	// 请求体
	Body interface{} `json:"body"`
}

type ClientRequestOption func(*ClientRequestOptions)

func WithMethod(method string) ClientRequestOption {
	return func(options *ClientRequestOptions) {
		options.Method = method
	}
}

func NewClient(accessKey, secret string) *Client {
	return &Client{
		AccessKey:  accessKey,
		Secret:     secret,
		httpClient: httputil.NewClient(),
	}
}

func (c *Client) DoRequest(options ClientRequestOptions) (*httputil.Response, error) {
	opts := httputil.RequestOptions{
		Method: options.Method,
	}

	resp, err := c.httpClient.DoRequest(opts)
	if err != nil {
		return nil, err
	}

	return c.httpClient.WrapHttpResponse(resp)
}
