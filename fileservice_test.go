package hpsdk_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/hpsvc/hpsdk"
	"github.com/hyperits/gosuite/converter"
)

// 测试 FileServiceClient 的 PostFile
func TestFileServiceClient_PostFile(t *testing.T) {

	filename := "resources/img.jpg"

	// 创建测试用的文件
	file, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	conf, err := hpsdk.LoadSdkConfig()
	if err != nil {
		t.Fatal(err)
	}

	// 创建测试用的文件服务
	cli, err := hpsdk.NewFileServiceClient(conf.AccessKey, conf.Secret, conf.EndPoint, conf.RequestTimeout)
	if err != nil {
		t.Fatal(err)
	}

	// 调用 PostFile 方法
	resp, err := cli.PostFile(filename, file)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("resp status code: %d", resp.StatusCode)
	}

	t.Logf("resp: %v", converter.ToJsonString(resp))
}

// 测试 FileServiceClient 的 GetFiles
func TestFileServiceClient_GetFiles(t *testing.T) {
	conf, err := hpsdk.LoadSdkConfig()
	if err != nil {
		t.Fatal(err)
	}

	// 创建测试用的文件服务
	cli, err := hpsdk.NewFileServiceClient(conf.AccessKey, conf.Secret, conf.EndPoint, conf.RequestTimeout)
	if err != nil {
		t.Fatal(err)
	}

	// 调用 GetFiles 方法
	resp, err := cli.GetFiles("")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("resp status code: %d", resp.StatusCode)
	}

	t.Logf("resp: %v", converter.ToJsonString(resp))
}

// 测试 FileServiceClient 的 GetFile
func TestFileServiceClient_GetFile(t *testing.T) {
	conf, err := hpsdk.LoadSdkConfig()
	if err != nil {
		t.Fatal(err)
	}

	filename := "resources/img.jpg"

	// 创建测试用的文件服务
	cli, err := hpsdk.NewFileServiceClient(conf.AccessKey, conf.Secret, conf.EndPoint, conf.RequestTimeout)
	if err != nil {
		t.Fatal(err)
	}
	// 调用 GetFile 方法
	resp, err := cli.GetFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("resp status code: %d", resp.StatusCode)
	}
	t.Logf("resp: %v", converter.ToJsonString(resp))
}

// 测试 FileServiceClient 的 InfoFile
func TestFileServiceClient_InfoFile(t *testing.T) {
	conf, err := hpsdk.LoadSdkConfig()
	if err != nil {
		t.Fatal(err)
	}

	filename := "resources/img.jpg"

	// 创建测试用的文件服务
	cli, err := hpsdk.NewFileServiceClient(conf.AccessKey, conf.Secret, conf.EndPoint, conf.RequestTimeout)
	if err != nil {
		t.Fatal(err)
	}
	// 调用 InfoFile 方法
	resp, err := cli.InfoFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("resp status code: %d", resp.StatusCode)
	}
	t.Logf("resp: %v", converter.ToJsonString(resp))
}

// 测试 FileServiceClient 的 DeleteFile
func TestFileServiceClient_DeleteFile(t *testing.T) {
	conf, err := hpsdk.LoadSdkConfig()
	if err != nil {
		t.Fatal(err)
	}

	filename := "resources/img.jpg"

	// 创建测试用的文件服务
	cli, err := hpsdk.NewFileServiceClient(conf.AccessKey, conf.Secret, conf.EndPoint, conf.RequestTimeout)
	if err != nil {
		t.Fatal(err)
	}
	// 调用 DeleteFile 方法
	resp, err := cli.DeleteFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("resp status code: %d", resp.StatusCode)
	}
	t.Logf("resp: %v", converter.ToJsonString(resp))
}
