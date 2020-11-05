package core

const (
	// 对称加密算法
	HmacWithSHA224     SignatureAlgorithm = "hmac-sha224"
	HmacWithSHA256     SignatureAlgorithm = "hmac-sha256"
	HmacWithSHA384     SignatureAlgorithm = "hmac-sha384"
	HmacWithSHA512     SignatureAlgorithm = "hmac-sha512"
	HmacWithRipemd160  SignatureAlgorithm = "hmac-ripemd160"
	HmacWithSHA3224    SignatureAlgorithm = "hmac-sha3-224"
	HmacWithSHA3256    SignatureAlgorithm = "hmac-sha3-256"
	HmacWithSHA3384    SignatureAlgorithm = "hmac-sha3-384"
	HmacWithSHA3512    SignatureAlgorithm = "hmac-sha3-512"
	HmacWithSHA512224  SignatureAlgorithm = "hmac-sha512-224"
	HmacWithSHA512256  SignatureAlgorithm = "hmac-sha512-256"
	HmacWithBlake2s256 SignatureAlgorithm = "hmac-blake2s-256"
	HmacWithBlake2b256 SignatureAlgorithm = "hmac-blake2b-256"
	HmacWithBlake2b384 SignatureAlgorithm = "hmac-blake2b-384"
	HmacWithBlake2b512 SignatureAlgorithm = "hmac-blake2b-512"
	Blake2sWith256     SignatureAlgorithm = "blake2s-256"
	Blake2bWith256     SignatureAlgorithm = "blake2b-256"
	Blake2bWith384     SignatureAlgorithm = "blake2b-384"
	Blake2bWith512     SignatureAlgorithm = "blake2b-512"
	// RAS非对称加密算法
	RsaWithSHA224    SignatureAlgorithm = "rsa-sha224"
	RsaWithSHA256    SignatureAlgorithm = "rsa-sha256"
	RsaWithSHA384    SignatureAlgorithm = "rsa-sha384"
	RsaWithSHA512    SignatureAlgorithm = "rsa-sha512"
	RsaWithRipemd160 SignatureAlgorithm = "rsa-ripemd160"
	// ECDSA非对称加密算法
	EcdsaWithSHA224    SignatureAlgorithm = "ecdsa-sha224"
	EcdsaWithSHA256    SignatureAlgorithm = "ecdsa-sha256"
	EcdsaWithSHA384    SignatureAlgorithm = "ecdsa-sha384"
	EcdsaWithSHA512    SignatureAlgorithm = "ecdsa-sha512"
	EcdsaWithRipemd160 SignatureAlgorithm = "ecdsa-ripemd160"
)

// SignatureAlgorithm 签名算法
type SignatureAlgorithm string
