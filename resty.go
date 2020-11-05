package class100

import (
	`crypto`
	`crypto/tls`
	`net/http`
	`time`

	`github.com/go-fed/httpsig`
	`github.com/go-resty/resty/v2`
)

// NewResty Resty客户端
func NewResty(client *HttpSignatureClient) *resty.Request {
	return resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetPreRequestHook(func(c *resty.Client, req *http.Request) (err error) {
			req.Header.Add("date", time.Now().Format(time.RFC1123))

			privateKey := crypto.PrivateKey([]byte(client.Options.Secret.SecretKey))
			preferAlgorithms := []httpsig.Algorithm{httpsig.HMAC_SHA512}
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
		}).R()
}

// RestyStringBody 字符串形式的结果
func RestyStringBody(rsp *resty.Response) string {
	body := ""
	if nil != rsp {
		body = rsp.String()
	}

	return body
}
