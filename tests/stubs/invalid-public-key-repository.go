package stubs

type invalidPublicKeyRepository struct{}

func NewInvalidPublicKeyRepository() *invalidPublicKeyRepository {
	return &invalidPublicKeyRepository{}
}

func (f *invalidPublicKeyRepository) Get() ([]byte, error) {
	return []byte(""), nil
}
