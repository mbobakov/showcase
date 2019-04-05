// Package clickhouse implements service.Metricer interface
package clickhouse

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/kshvakov/clickhouse" // swagger struct
	"github.com/mbobakov/showcase/storage"
	"github.com/pkg/errors"
)

// DB connection to clickhouse
type DB struct {
	sql   *sqlx.DB
	table string
}

// Init the connection and basic DB structure
func Init(dsn, table string) (*DB, error) {
	db, err := sqlx.Open("clickhouse", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't open a connection to the database")
	}
	err = migrate(context.Background(), table, db)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't migrate the database")
	}
	return &DB{sql: db, table: table}, nil

}

// Summary for the metrics
func (c *DB) Summary() (map[string]int64, error) {
	type resp struct {
		Name  string `sqlx:"name"`
		Count int64  `sqlx:"count"`
	}
	var res []resp
	query := fmt.Sprintf("SELECT name,count() as count FROM %s GROUP BY name", c.table)
	err := c.sql.Select(&res, query)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't select summary")
	}
	sum := make(map[string]int64, len(res))
	for _, r := range res {
		sum[r.Name] = r.Count
	}
	return sum, nil
}

//Store the datapoint for the metric
func (c *DB) Store(m *storage.MetricDTO) error {
	tx, err := c.sql.Beginx()
	if err != nil {
		return errors.Wrap(err, "Couldn't start a batch insertion")
	}
	query := fmt.Sprintf("INSERT INTO %s VALUES (?,?,?)", c.table)
	for t, v := range m.DataPoints {

		_, err = tx.Exec(query, t, m.Name, v)
		if err != nil {
			return errors.Wrap(err, "Couldn't add insert into the batch")
		}
	}
	return errors.Wrap(tx.Commit(), "Couldn't committhe batch")
}

func migrate(ctx context.Context, table string, db *sqlx.DB) error {
	createTableDDL := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s
		(
    	time UInt64,
    	name String,
    	value Int64
		) ENGINE = MergeTree()
			PARTITION BY toYYYYMM(toDateTime(time))
			ORDER BY (time,name)
	`, table)

	_, err := db.ExecContext(ctx, createTableDDL)
	return errors.Wrap(err, "Couldn't init a table")
}
