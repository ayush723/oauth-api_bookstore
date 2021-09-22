package cassandra

import (
	"os"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {

	host := os.Getenv("cassandra_host")
	if host == "" {
		host = "127.0.0.1"
	}

	//Connect to cassandra cluster:
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
