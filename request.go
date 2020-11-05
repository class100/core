package core

const (
	// Http常用头
	// HeaderAcceptLanguage 可接受的语言
	HeaderAcceptLanguage = "Accept-Language"
	// HeaderContentType 请求数据类型
	HeaderContentType = "Content-Type"
	// HeaderAuthorization 授权
	HeaderAuthorization = "Authorization"
	// HeaderContentDisposition
	HeaderContentDisposition = "Content-disposition"

	// Http方法集合
	// HttpMethodGet GET方法
	HttpMethodGet HttpMethod = "GET"
	// HttpMethodPost POST方法
	HttpMethodPost HttpMethod = "POST"
	// HttpMethodPut PUT方法
	HttpMethodPut HttpMethod = "PUT"
	// HttpMethodDelete DELETE方法
	HttpMethodDelete HttpMethod = "DELETE"
	// MethodPatch PATCH方法
	HttpMethodPatch HttpMethod = "PATCH"
	// MethodHead HEAD方法
	HttpMethodHead HttpMethod = "HEAD"
	// MethodOptions OPTIONS方法
	HttpMethodOptions HttpMethod = "OPTIONS"
)

type (
	// HttpMethod Http方法
	HttpMethod string
)
