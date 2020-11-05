package core

import (
	`encoding/json`
	`net/http`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

// HttpSignatureClient 基于Http签名的客户端
type HttpSignatureClient struct {
	Options HttpSignatureClientOptions
}

// NewHttpSignatureClient 创建默认的Http签名的客户端
func NewHttpSignatureClient(options ...Option) (client *HttpSignatureClient, err error) {
	appliedOptions := defaultOptions()
	for _, apply := range options {
		apply(&appliedOptions)
	}

	client = &HttpSignatureClient{
		Options: appliedOptions,
	}

	return
}

func (hsc *HttpSignatureClient) RequestApi(
	url string,
	method HttpMethod,
	headers map[string]string,
	params interface{}, pathParams map[string]string,
	rsp interface{},
) (err error) {
	var (
		serverRsp          *resty.Response
		expectedStatusCode int
	)

	req := NewResty(hsc).SetResult(rsp)
	// 注入路径参数
	if 0 != len(pathParams) {
		req = req.SetPathParams(pathParams)
	}

	// 注入请求头
	if 0 != len(headers) {
		req.SetHeaders(headers)
	}

	switch method {
	case HttpMethodGet:
		expectedStatusCode = http.StatusOK

		if nil != params {
			req = req.SetQueryParams(params.(map[string]string))
		}
		serverRsp, err = req.Get(url)
	case HttpMethodPost:
		expectedStatusCode = http.StatusCreated

		if nil != params {
			req = req.SetBody(params)
		}
		serverRsp, err = req.Post(url)
	case HttpMethodPut:
		expectedStatusCode = http.StatusOK

		if nil != params {
			req = req.SetBody(params)
		}
		serverRsp, err = req.Put(url)
	case HttpMethodDelete:
		expectedStatusCode = http.StatusNoContent

		if nil != params {
			req = req.SetBody(params)
		}
		serverRsp, err = req.Delete(url)
	}
	if nil != err {
		return
	}

	if nil == serverRsp {
		err = gox.NewCodeError(gox.ErrorCode(serverRsp.StatusCode()), "无返回数据", RestyStringBody(serverRsp))

		return
	}

	// 检查状态码
	if expectedStatusCode != serverRsp.StatusCode() {
		err = gox.NewCodeError(gox.ErrorCode(serverRsp.StatusCode()), "请求服务器不符合预期", RestyStringBody(serverRsp))
	}

	return
}

func (hsc HttpSignatureClient) String() string {
	jsonBytes, _ := json.MarshalIndent(hsc, "", "    ")

	return string(jsonBytes)
}
