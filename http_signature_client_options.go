package core

import (
	`time`
)

type (
	// Option 配置选项
	Option func(*HttpSignatureClientOptions)

	// HttpSignatureClientOptions 客户端配置
	HttpSignatureClientOptions struct {
		// Secret 应用授权
		Secret Secret
		// Algorithms 签名算法
		Algorithms []SignatureAlgorithm
		// Proxy 代理地址
		Proxy string
		// Timeout 超时时间
		Timeout time.Duration
		// Retry 重试机制
		Retry Retry
	}
)

func defaultOptions() HttpSignatureClientOptions {
	return HttpSignatureClientOptions{
		Algorithms: []SignatureAlgorithm{
			HmacWithSHA512,
		},
	}
}

// WithSecret 配置应用授权
func WithSecret(secret Secret) Option {
	return func(options *HttpSignatureClientOptions) {
		options.Secret = secret
	}
}

// WithAlgorithms 配置签名算法
func WithAlgorithms(algorithms ...SignatureAlgorithm) Option {
	return func(options *HttpSignatureClientOptions) {
		options.Algorithms = algorithms
	}
}

// WithProxy 配置代理
func WithProxy(proxy string) Option {
	return func(options *HttpSignatureClientOptions) {
		options.Proxy = proxy
	}
}

// WithRetry 配置重试
func WithRetry(retry Retry) Option {
	return func(options *HttpSignatureClientOptions) {
		options.Retry = retry
	}
}
