package database

import (
	"fmt"
	"os"

	"database/sql"
	_ "github.com/lib/pq"
)

func New() (*Postgres, func() error, error) {
	conn, err := sql.Open(
		"postgres",
		fmt.Sprintf("postgres://%s:%s@%v:%v/%v?sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DATABASE"),
		),
	)

	if err != nil {
		return nil, nil, fmt.Errorf("connecting to mysql: %w", err)
	}

	err = conn.Ping()

	if err != nil {
		return nil, nil, fmt.Errorf("ping: %w", err)
	}

	return &Postgres{client: conn}, conn.Close, nil
}
