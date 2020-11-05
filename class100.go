package class100

import (
	`encoding/json`
	`net/http`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

type (
	// Client 客户端
	Client struct {
		// 授权
		AccessKey string `json:"accessKey"`
		SecretKey string `json:"secretKey"`
	}
)

func (c *Client) RequestApi(
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

	req := NewResty(c).SetResult(rsp)
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

func (c Client) String() string {
	jsonBytes, _ := json.MarshalIndent(c, "", "    ")

	return string(jsonBytes)
}
