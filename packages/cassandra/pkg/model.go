package pkg

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	gocql "github.com/gocql/gocql"
	"github.com/mitchellh/mapstructure"
)

type Signature struct {
	Signer    string `cql:"signer"`
	Comment   string `cql:"comment"`
	Timestamp int64  `cql:"timestamp"`
}

type Document struct {
	Company    string      `cql:"company"`
	Id         gocql.UUID  `cql:"id"`
	Name       string      `cql:"name"`
	Tags       []string    `cql:"tags"`
	Signatures []Signature `cql:"signatures"`
	Status     string      `cql:"status"`
	Timestamp  time.Time   `cql:"timestamp"`
}

var Session *gocql.Session

func initSession() {
	port := func(p string) int {
		i, err := strconv.Atoi(p)
		if err != nil {
			return 9042
		}

		return i
	}

	consistancy := func(c string) gocql.Consistency {
		gc, err := gocql.MustParseConsistency(c)
		if err != nil {
			return gocql.All
		}

		return gc
	}

	cluster := gocql.NewCluster(cassandraConfig.host)
	cluster.Port = port(cassandraConfig.port)
	cluster.Keyspace = cassandraConfig.keyspace
	cluster.Consistency = consistancy(cassandraConfig.consistancy)

	s, err := cluster.CreateSession()
	if err != nil {
		log.Printf("ERROR: fail create cassandra session, %s", err.Error())
		os.Exit(1)
	}
	Session = s
}

func clearSession() {
	Session.Close()
}

func createDocument(document *Document) error {
	q := `
		INSERT INTO documents (
		    company,
		    id,
		    name,
		    tags,
		    signatures,
				status,
				timestamp
		)
		VALUES (?, ?, ?, ?, ?, ?, ?)
    	`
	err := Session.Query(q,
		document.Company,
		document.Id,
		document.Name,
		document.Tags,
		document.Signatures,
		document.Status,
		document.Timestamp).Exec()
	if err != nil {
		log.Printf("ERROR: fail create document, %s", err.Error())
	}

	return err
}

func getDocument(company string, id gocql.UUID) (*Document, error) {
	toSignatures := func(i interface{}) []Signature {
		sigs := []Signature{}
		sig := Signature{}
		for _, s := range i.([]map[string]interface{}) {
			mapstructure.Decode(s, &sig)
			sigs = append(sigs, sig)
		}

		return sigs
	}

	m := map[string]interface{}{}
	q := `
		SELECT * FROM documents
			WHERE company = ? AND id = ?
		LIMIT 1
    	`
	itr := Session.Query(q, company, id).Consistency(gocql.One).Iter()
	for itr.MapScan(m) {
		document := &Document{}
		document.Company = m["company"].(string)
		document.Id = m["id"].(gocql.UUID)
		document.Name = m["name"].(string)
		document.Tags = m["tags"].([]string)
		document.Signatures = toSignatures(m["signatures"])
		document.Status = m["status"].(string)
		document.Timestamp = m["timestamp"].(time.Time)

		log.Printf("INFO: found document, %v", document)

		return document, nil
	}

	return nil, errors.New("document not found")
}

func updateDocument(company string, id gocql.UUID, name string, status string) error {
	q := `
        	UPDATE documents
		SET name = ?, status = ?
		WHERE company = ? AND id = ?
    	`
	err := Session.Query(q, name, status, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail update document, %s", err.Error())
		return err
	}

	return nil
}

func addTag(company string, id gocql.UUID, tag string) error {
	q := `
		UPDATE mystiko.documents
		SET tags = tags + ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []string{tag}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail add tag, %s", err.Error())
		return err
	}

	return nil
}

func removeTag(company string, id gocql.UUID, tag string) error {
	q := `
		UPDATE mystiko.documents
		SET tags = tags - ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []string{tag}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail remove tag, %s", err.Error())
		return err
	}

	return nil
}

func addSignature(company string, id gocql.UUID, signature Signature) error {
	q := `
		UPDATE mystiko.documents
		SET signatures = signatures + ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []Signature{signature}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail add signature, %s", err.Error())
		return err
	}

	return nil
}

func removeSignature(company string, id gocql.UUID, signature Signature) error {
	q := `
		UPDATE mystiko.documents
		SET signatures = signatures - ?
		WHERE company = ? AND id = ?;
	`

	err := Session.Query(q, []Signature{signature}, company, id).Exec()
	if err != nil {
		log.Printf("ERROR: fail remove signature, %s", err.Error())
		return err
	}

	return nil
}
