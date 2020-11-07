package core

import (
	`encoding/json`
	`fmt`
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
	data interface{},
	rsp interface{},
	params ...gox.HttpParameter,
) (err error) {
	var (
		serverRsp          *resty.Response
		expectedStatusCode int
	)

	req := NewResty(hsc).SetResult(rsp)

	// 注入请求头和路径参数
	var (
		headers    = make(map[string]string)
		pathParams = make(map[string]string)
	)
	for _, param := range params {
		switch param.Type() {
		case gox.HttpParameterTypeHeader:
			headers[param.Key()] = param.Value()
		case gox.HttpParameterTypePathParameter:
			pathParams[param.Key()] = param.Value()
		}
	}
	if 0 != len(headers) {
		req.SetHeaders(headers)
	}
	if 0 != len(pathParams) {
		req.SetPathParams(pathParams)
	}

	switch method {
	case HttpMethodGet:
		expectedStatusCode = http.StatusOK

		if nil != data {
			var (
				flattenParams map[string]interface{}
				paramMap      = make(map[string]string)
			)

			if flattenParams, err = gox.StructToMap(data); nil != err {
				return
			}
			if flattenParams, err = gox.Flatten(flattenParams, "", gox.DotStyle); nil != err {
				return
			}

			for key, value := range flattenParams {
				paramMap[key] = fmt.Sprintf("%s", value)
			}
			req = req.SetQueryParams(paramMap)
		}
		serverRsp, err = req.Get(url)
	case HttpMethodPost:
		expectedStatusCode = http.StatusCreated

		if nil != data {
			req = req.SetBody(data)
		}
		serverRsp, err = req.Post(url)
	case HttpMethodPut:
		expectedStatusCode = http.StatusOK

		if nil != data {
			req = req.SetBody(data)
		}
		serverRsp, err = req.Put(url)
	case HttpMethodDelete:
		expectedStatusCode = http.StatusNoContent

		if nil != data {
			req = req.SetBody(data)
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
