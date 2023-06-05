package repositories

import "os"

type privateKeyRepository struct{}

func NewPrivateKeyRepository() *privateKeyRepository {
	return &privateKeyRepository{}
}

func (r *privateKeyRepository) Get() ([]byte, error) {
	return os.ReadFile("keys/private.pem")
}
