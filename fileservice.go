package hpsdk

import (
	"bytes"
	"io"
	"mime/multipart"

	"github.com/hyperits/gosuite/httputil"
)

type FileServiceClient struct {
	Client
}

const (
	URI_ROOT = `/`
	URI_LIST = `/list`
	URI_INFO = `/info`
)

func NewFileServiceClient(accessKey, secret string, endpoint string, requestTimeout int) (*FileServiceClient, error) {
	client, err := NewClient(accessKey, secret, endpoint, requestTimeout)
	if err != nil {
		return nil, err
	}
	return &FileServiceClient{*client}, nil
}

// PostFile 上传文件
func (c *FileServiceClient) PostFile(filename string, reader io.Reader) (*httputil.Response, error) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, err
	}

	err = writer.WriteField("filename", filename)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return c.DoRequest(
		WithMethod(httputil.POST),
		WithURI(URI_ROOT),
		WithAuthReqired(true),
		WithHeaders(map[string]string{
			httputil.CONTENT_TYPE: writer.FormDataContentType(),
		}),
		WithBody(body),
	)
}

// GetFiles 获取文件列表
func (c *FileServiceClient) GetFiles(prefix string) (*httputil.Response, error) {
	return c.DoRequest(
		WithMethod(httputil.GET),
		WithURI(URI_LIST),
		WithAuthReqired(true),
		WithQueryParams(map[string]string{
			"prefix": prefix,
		}),
	)
}

// GetFile 获取文件
func (c *FileServiceClient) GetFile(filename string) (*httputil.Response, error) {
	return c.DoRequest(
		WithMethod(httputil.GET),
		WithURI(URI_ROOT),
		WithQueryParams(map[string]string{
			"filename": filename,
		}),
	)
}

// DeleteFile 删除文件
func (c *FileServiceClient) DeleteFile(filename string) (*httputil.Response, error) {
	return c.DoRequest(
		WithMethod(httputil.DELETE),
		WithURI(URI_ROOT),
		WithAuthReqired(true),
		WithQueryParams(map[string]string{
			"filename": filename,
		}),
	)
}

// InfoFile 获取文件信息
func (c *FileServiceClient) InfoFile(filename string) (*httputil.Response, error) {
	return c.DoRequest(
		WithMethod(httputil.GET),
		WithURI(URI_INFO),
		WithQueryParams(map[string]string{
			"filename": filename,
		}),
	)
}
