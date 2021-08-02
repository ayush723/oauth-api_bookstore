package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)


 
func init(){
	//Connect to cassandra cluster:
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil{
		panic(err)
	}
	fmt.Println("cassandra connection successfully created")
	defer session.Close()
}