package queries

import (
	"context"

	"github.com/edgedb/edgedb-go"
)

func CreateTestResult(ctx context.Context, client *edgedb.Client) error {
	_, err := createTest(ctx, client)
	return err
}

func TestQueryCorrected(ctx context.Context, client *edgedb.Client) error {
	_, err := testQueryCorrected(ctx, client)
	return err
}

func TestQuery(ctx context.Context, client *edgedb.Client) error {
	_, err := testQuery(ctx, client)
	return err
}
