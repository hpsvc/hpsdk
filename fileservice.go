package hpsdk

import "github.com/hyperits/gosuite/httputil"

type FileServiceClient struct {
	Client
}

const (
	URI_ROOT = `/`
)

func NewFileServiceClient(accessKey, secret string) *FileServiceClient {
	client := NewClient(accessKey, secret)
	return &FileServiceClient{*client}
}

func (c *FileServiceClient) PostFile(id string, params map[string]interface{}) (, error) {

	opts := ClientRequestOptions{
		Method:      httputil.GET,
		URI:         URI_ROOT,
		AuthReqired: true,
	}

	c.DoRequest(opts)

	return nil, nil
}
