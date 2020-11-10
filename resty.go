package core

import (
	`crypto`
	`crypto/tls`
	`net/http`
	`strings`
	`time`

	`github.com/go-fed/httpsig`
	`github.com/go-resty/resty/v2`
)

// NewResty Resty客户端
func NewResty(client *HttpSignatureClient) *resty.Request {
	rc := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetPreRequestHook(func(c *resty.Client, req *http.Request) (err error) {
			req.Header.Add("date", time.Now().Format(time.RFC1123))

			privateKey := crypto.PrivateKey([]byte(client.Options.Secret.SecretKey))

			preferAlgorithms := make([]httpsig.Algorithm, 0, len(client.Options.Algorithms))
			for _, algorithm := range client.Options.Algorithms {
				preferAlgorithms = append(preferAlgorithms, httpsig.Algorithm(algorithm))
			}
			digestAlgorithm := httpsig.DigestSha256
			headersToSign := []string{httpsig.RequestTarget, "date"}

			var signer httpsig.Signer
			if signer, _, err = httpsig.NewSigner(
				preferAlgorithms,
				digestAlgorithm,
				headersToSign,
				httpsig.Signature,
				time.Now().Add(time.Minute).Unix(),
			); nil != err {
				return
			}

			err = signer.SignRequest(privateKey, client.Options.Secret.SecretId, req, nil)

			return
		})
	if "" != strings.TrimSpace(client.Options.Proxy) {
		rc.SetProxy(client.Options.Proxy)
	}
	if 0 != client.Options.Timeout {
		rc.SetTimeout(client.Options.Timeout)
	}

	// 注入重试参数
	if 0 != client.Options.Retry.Count {
		rc.SetRetryCount(client.Options.Retry.Count)
	}
	if 0 != client.Options.Retry.WaitTime {
		rc.SetRetryWaitTime(client.Options.Retry.WaitTime)
	}
	if 0 != client.Options.Retry.MaxWaitTime {
		rc.SetRetryMaxWaitTime(client.Options.Retry.MaxWaitTime)
	}

	return rc.R()
}

// RestyStringBody 字符串形式的结果
func RestyStringBody(rsp *resty.Response) string {
	body := ""
	if nil != rsp {
		body = rsp.String()
	}

	return body
}
