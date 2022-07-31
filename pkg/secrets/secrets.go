package secrets

type SecretContents struct {
	UserId          uint64
	RotationTime    int64
	SecretId        string
	EncryptedSecret string
}

type SecretsManagement interface {
	CreateSecret(userid uint64, secret string) (string, error)
	DeleteSecret(secretid string) error
	GetSecret(secretid string) (*SecretContents, error)
}
