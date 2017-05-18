package postgres

import (
	"os"

	scores "github.com/Kunde21/scoresServer"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/pkg/errors"
)

var _ scores.Client = &Client{}

type Client struct {
	competitionSvc CompetitionService
	teamSvc        TeamService
	fixtureSvc     FixtureService

	db *sqlx.DB
}

func NewClient() *Client {
	c := &Client{}
	c.competitionSvc.client = c
	c.teamSvc.client = c
	c.fixtureSvc.client = c
	return c
}

func (c *Client) Connect() error {
	db, err := sqlx.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_ADDR"))
	if err != nil {
		return err
	}

	c.db = db
	return nil
}

func (c *Client) Migrate() error {
	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver != "postgres" {
		return errors.New("Auto-migration is only supported on postgres databases.")
	}

	driver, err := postgres.WithInstance(c.db.DB, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "Instance creation")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		dbDriver, driver)
	if err != nil {
		return errors.Wrap(err, "Loading migrations")
	}

	err = m.Up()
	if err != nil {
		return errors.Wrap(err, "Migrating")
	}
	return nil
}

func (c *Client) CompetitionService() scores.CompetitionService {
	return c.competitionSvc
}

func (c *Client) TeamService() scores.TeamService {
	return c.teamSvc
}

func (c *Client) FixtureService() scores.FixtureService {
	return c.fixtureSvc
}
