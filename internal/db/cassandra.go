package db

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"techytechster.com/secretsmanager/pkg/secrets"
)

type CassandraSecretsManager struct {
	session *gocql.Session
}

func (sm CassandraSecretsManager) CreateSecret(userid uint64, secret string) (string, error) {
	timestamp := time.Now().Unix()
	uuid := gocql.TimeUUID()
	return uuid.String(), sm.session.Query(`INSERT INTO secrets (user_id, rotation_time, secret_id, encrypted_secret) VALUES (?, ?, ?, ?)`,
		userid,
		timestamp,
		uuid,
		secret,
	).WithContext(context.TODO()).Exec()
}

func (sm CassandraSecretsManager) DeleteSecret(secretid string) error {
	err := sm.session.Query(`DELETE FROM secrets WHERE secret_id = ?`,
		secretid,
	).WithContext(context.TODO()).Exec()
	if err != nil {
		return fmt.Errorf("failed to delete requested secretid: %s, : %w", secretid, err)
	}
	return nil
}

func (sm CassandraSecretsManager) GetSecret(secretid string) (*secrets.SecretContents, error) {
	var userId uint64
	var rotationTime int64
	var encryptedSecret, secretId string
	err := sm.session.Query(`SELECT user_id, rotation_time, secret_id, encrypted_secret FROM secrets WHERE secret_id = ? LIMIT 1`,
		secretid,
	).WithContext(context.TODO()).Consistency(gocql.One).Scan(&userId, &rotationTime, &secretId, &encryptedSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to query cassandra for requested secretid: %s, : %w", secretid, err)
	}
	return &secrets.SecretContents{
		UserId:          userId,
		RotationTime:    rotationTime,
		SecretId:        secretId,
		EncryptedSecret: encryptedSecret,
	}, nil
}

func InitializeCassandra() (secrets.SecretsManagement, error) {
	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "SecretsManager"
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create a cassandra session for cassandrasecretsmanager: %w", err)
	}
	sm := CassandraSecretsManager{session}
	return &sm, nil
}
