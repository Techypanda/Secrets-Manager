package main

import (
	"techytechster.com/secretsmanager/internal/db"
)

func main() {
	cassandra, err := db.InitializeCassandra()
	if err != nil {
		panic(err)
	}
	uuid, err := cassandra.CreateSecret(1234, "acoolsecret")
	if err != nil {
		panic(err)
	}
	_, err = cassandra.GetSecret(uuid)
	if err != nil {
		panic(err)
	}
}
