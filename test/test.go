package main

import (
	"context"
	"log"
	"test/test/queries"

	"github.com/edgedb/edgedb-go"
)

//go:generate edgeql-go
func main() {
	ctx := context.Background()

	edgeDBClient, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Creating test object")

	err = queries.CreateTestResult(ctx, edgeDBClient)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Executing correct test query")

	err = queries.TestQueryCorrected(ctx, edgeDBClient)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Executing generated test query")

	err = queries.TestQuery(ctx, edgeDBClient)
	if err != nil {
		log.Fatal(err)
	}
}
