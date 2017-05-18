package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"github.com/pkg/errors"
)

func testDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", os.Getenv("TEST_DB"))
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "Instance creation")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return nil, errors.Wrap(err, "Loading migrations")
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, errors.Wrap(err, "Migrating")
	}
	// Closing migration and database here is necessary
	// in order to cleanup the test database.
	// Drop() will fail "unable to acquire lock" without this Close().
	err1, err2 := m.Close()
	if err1 != nil {
		log.Println(errors.Wrap(err1, "Close source"))
	}
	if err2 != nil {
		log.Println(errors.Wrap(err2, "Close db"))
	}

	db, err = sqlx.Connect("postgres", os.Getenv("TEST_DB"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func dbCleanup(t *testing.T, db *sqlx.DB) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		t.Error(errors.Wrap(err, "Instance creation"))
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		t.Error(errors.Wrap(err, "Loading migrations"))
	}

	err = m.Drop()
	if err != nil {
		t.Error(errors.Wrap(err, "Migrating"))
	}
	err1, err2 := m.Close()
	if err1 != nil {
		t.Error(errors.Wrap(err1, "Close source"))
	}
	if err2 != nil {
		t.Error(errors.Wrap(err2, "Close db"))
	}
}
