package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}

	fmt.Println("========================================")
	fmt.Println("Cassandra connection succesfully created")
	fmt.Println("========================================")
}

// GetSession for connecting to Cassandra
func GetSession() *gocql.Session {
	return session
}
