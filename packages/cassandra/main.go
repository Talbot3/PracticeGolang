package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var Session *gocqlx.Session

type Emp struct {
	id        string
	firstName string
	lastName  string
	age       int
}

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1:9042")
	cluster.ProtoVersion = 4
	cluster.CQLVersion = "3.4.5"

	// Wrap session on creation, gocqlx session embeds gocql.Session pointer.
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	Session = &session
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("cassandra init done")
}

func main() {
	fmt.Println("init done")
}
