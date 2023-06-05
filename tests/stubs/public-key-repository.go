package stubs

type publicKeyRepository struct{}

func NewPublicKeyRepository() *publicKeyRepository {
	return &publicKeyRepository{}
}

func (f *publicKeyRepository) Get() ([]byte, error) {
	return []byte("-----BEGIN PUBLIC KEY-----\n" +
		"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDbKEW+yAQvf+MGT3BNkeXioNbU\n" +
		"DpEqT8nGRsqnZdffxwRbsC0McTKVW2r45BCkxq7+f0IuA/3yQ07Z+DZSK3lrhS9g\n" +
		"jvN2mU9grlrNbRoPYyoobb7Y5frjIw4t7VQNSarh+sMvUW6kju4e+z15feEVjIqC\n" +
		"tDp+OCV+nVx/uRRgBwIDAQAB\n" +
		"-----END PUBLIC KEY-----"), nil
}
