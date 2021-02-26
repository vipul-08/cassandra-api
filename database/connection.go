package database

import (
	"fmt"
	"github.com/gocql/gocql"
	"time"
)

type DbConnection struct {
	Cluster *gocql.ClusterConfig
	Session *gocql.Session
}

var connection DbConnection

func ConnectDatabase() {
	connection.Cluster = gocql.NewCluster("127.0.0.1")
	connection.Cluster.Port = 9042
	connection.Cluster.ConnectTimeout = time.Second * 1000
	connection.Cluster.Consistency = gocql.Quorum
	connection.Cluster.Keyspace = "student_api"
	session, err := connection.Cluster.CreateSession()

	if err != nil {
		panic(err)
	}

	connection.Session = session
	fmt.Println("Connected to Database")
}

func GetDbConnection() DbConnection {
	return connection
}
