package repositories

import "os"

type publicKeyRepository struct{}

func NewPublicKeyRepository() *publicKeyRepository {
	return &publicKeyRepository{}
}

func (r *publicKeyRepository) Get() ([]byte, error) {
	return os.ReadFile("keys/public.pem")
}
