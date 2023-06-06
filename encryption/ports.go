package encryption

type KeyRepository interface {
	Get() ([]byte, error)
}
