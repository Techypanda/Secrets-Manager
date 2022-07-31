package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

func main() {
	fmt.Println("AA")
	cluster := gocql.NewCluster("cassandra")
	cluster.Keyspace = "SecretsManager"
	_, err := cluster.CreateSession()
	if err != nil {
		panic(err.Error())
	}
}
