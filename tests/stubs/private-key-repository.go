package stubs

type privateKeyRepository struct{}

func NewPrivateKeyRepository() *privateKeyRepository {
	return &privateKeyRepository{}
}

func (f *privateKeyRepository) Get() ([]byte, error) {
	return []byte("-----BEGIN RSA PRIVATE KEY-----\n" +
		"MIICXgIBAAKBgQDbKEW+yAQvf+MGT3BNkeXioNbUDpEqT8nGRsqnZdffxwRbsC0M\n" +
		"cTKVW2r45BCkxq7+f0IuA/3yQ07Z+DZSK3lrhS9gjvN2mU9grlrNbRoPYyoobb7Y\n" +
		"5frjIw4t7VQNSarh+sMvUW6kju4e+z15feEVjIqCtDp+OCV+nVx/uRRgBwIDAQAB\n" +
		"AoGBAJjlPdNeA92HoUFTxou3j+RORdJn30rfo2ubE+M5oCPU2uuEu8kEaD/fKYA+\n" +
		"zIXIv2l6KxejXVXZM3fKPUuSXQAlld314c1RRj8O5yB5uCPTKhHeNP17Ce9sQ/iG\n" +
		"y5HA1R1Pg4/8TfS9e/oBBjBUDvvGddlEehIm/HnUTiZq6uxhAkEA9TTXZEqvxqNQ\n" +
		"5x4bal9KrfPEgXWdbjmK+QZ98bZWK85FjyXGmdKmEUXwAN7c71fRFAJ4xNT+pEGa\n" +
		"Vs6yJddJTwJBAOTN5DhJ6SZ7MPOrj2R/L5BqgwHkw7zInt4wSY+ICF4tyQqUD7oB\n" +
		"jhwYUwIHK0w2sYslD2ue7hBDm8wL9PbQ38kCQQCQ1VNcQ9krZQ/GU7wxynNb6B5r\n" +
		"sUL59jcl/DWr7d1cJoxjjxaNiSTaEws+GGU7Yg122jbnXD/EWZOGLrR/MXh1AkEA\n" +
		"l7NbKWb3yqfLIQQSopMhOiMkkVDX5octritd2mhbA/ZfYUAh3TXCvyXxg1q1fmaU\n" +
		"+KitvLE3LlllJriWxj5e4QJADHYMKS1eeBgVKjUGbWU+xUczyz/T0sg4mbO03J0K\n" +
		"Duu0Mo6qPr3aI36wlfcseJsbRvkTNEvXEI5ZYJApJhKj0A==\n" +
		"-----END RSA PRIVATE KEY-----"), nil
}
