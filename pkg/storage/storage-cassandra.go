package storage

import (
	"log"

	"github.com/gocql/gocql"
)

type Storage struct {
	db *gocql.Session
}

func SetupStorage() (*Storage, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "inspectors"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.NewSession()
	if err != nil {
		return &Storage{}, err
	}
	return &Storage{db: session}, nil
}

func (s *Storage) GetAllProductNames() ([]string, error) {
	var product string
	var products []string
	iter := s.db.Query(`SELECT mediaid from products`).Iter()
	for iter.Scan(&product) {
		products = append(products, product)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return products, nil
}
