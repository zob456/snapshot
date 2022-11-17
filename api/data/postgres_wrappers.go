package data

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// Non-exported helper funcs

func getRowsAffected(results sql.Result, targetNumRowsAffected int64) error {
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != targetNumRowsAffected {
		log.Printf(fmt.Sprintf("Rows affected: %v / %v", rowsAffected, targetNumRowsAffected))
		return errors.New("number of rows affected does not match the expected number of rows affected")
	}
	return nil
}

// Exported PostgreSQL methods

func PostgresUpdateColumnDataOneRow(db *sql.DB, query string, params ...interface{}) error {
	log.Printf(fmt.Sprintf("Query: %+v", query))
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	results, err := stmt.Exec(params...)
	err = getRowsAffected(results, 1)
	return nil
}

func PostgresScanOneRow(db *sql.DB, query string, params ...interface{}) (*sql.Row, error) {
	if len(params) < 1 {
		noParamsErr := errors.New("no params were passed")
		return nil, noParamsErr
	}
	log.Printf(fmt.Sprintf("Query: %+v", query))
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(params...)
	return row, nil
}