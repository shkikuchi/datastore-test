package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/datastore"
)

const (
	PROJECT_ID  = "deep-bolt-382805"
	DATABASE_ID = "test-db"
	KIND        = "Students"
)

type Student struct {
	Name   string
	Height int
	Weight int
}

func main() {

	ctx := context.Background()

	// get client for given database.
	// if you use default db, this can be replaced with:
	// dsClient, err := datastore.NewClient(ctx, PROJECT_ID)
	dsClient, err := datastore.NewClientWithDatabase(ctx, PROJECT_ID, DATABASE_ID)
	if err != nil {
		log.Fatalf("could not get datastore client: %v", err)
	}
	defer dsClient.Close()

	// execute query.
	var students []Student
	q := datastore.NewQuery(KIND).
		FilterField("Height", "=", 170)
	keys, err := dsClient.GetAll(ctx, q, &students)
	if err != nil {
		log.Fatalf("query failed: %v", err)
	} else if len(keys) == 0 {
		log.Fatal("query failed: no such entity")
	}
	fmt.Printf("got entities: %+v\n", students)

}
