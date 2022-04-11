package database

import (
	"database/sql"
	"fmt"
	"strings"
)

type Postgres struct {
	client *sql.DB
}

func (p Postgres) Insert(m map[string]string) error {
	values := make([]string, 0, len(m))
	for abbr, price := range m {
		values = append(values, fmt.Sprintf("('%s','%s',date)", abbr, price))
	}
	_, err := p.client.Exec(`DO $$
	DECLARE date timestamp with time zone := NOW();
	BEGIN
	INSERT INTO currencies VALUES ` + strings.Join(values, ",") + `;
	END $$;`)

	return err
}
