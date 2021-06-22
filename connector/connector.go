package connector

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	emacsconfigUrl = "postgres://weiss@localhost:5432/emacsconfig"
)

var (
	filenameFormat = "weiss_%s<%s.el"
)

func Connect() *pgxpool.Pool {
	conn, err := pgxpool.Connect(context.Background(), emacsconfigUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

