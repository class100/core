package class100

// Secret 应用授权密钥
type Secret struct {
	// SecretId 标识客户端，相当于用户名
	SecretId string
	// SecretKey 标识授权码，相当于密码
	SecretKey string
}
