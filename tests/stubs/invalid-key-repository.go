package stubs

type invalidKeyRepository struct{}

func NewInvalidKeyRepository() *invalidKeyRepository {
	return &invalidKeyRepository{}
}

func (f *invalidKeyRepository) Get() ([]byte, error) {
	return []byte(""), nil
}
