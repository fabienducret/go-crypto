package decryption

type KeyRepository interface {
	Get() ([]byte, error)
}
