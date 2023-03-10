package fasthttp

import (
	"github.com/kgip/fasthttp"
	"os"
	"testing"
)

func TestMultipart(t *testing.T) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI("http://www.baidu.com/upload")
	req.Header.SetMethod(fasthttp.MethodPost)
	file, err := os.Open("test.txt")
	defer file.Close()
	if err != nil {
		t.Log(err)
	}
	multiPartReader := &fasthttp.MultipartReader{}
	multiPartReader.CreateFormFile("file", "test.txt", file)
	fileInfo, _ := os.Stat("test.txt")
	req.SetBodyStream(multiPartReader, int(fileInfo.Size()))
	resp := fasthttp.AcquireResponse()
	err = fasthttp.Do(req, resp)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(resp.Body()))
}
