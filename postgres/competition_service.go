package postgres

import (
	"time"

	scores "github.com/Kunde21/scoresServer"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var _ scores.CompetitionService = CompetitionService{}

type CompetitionService struct {
	client *Client
}

func (cs CompetitionService) Competitions() (cmp []scores.Competition, err error) {
	err = cs.client.db.Select(&cmp, "SELECT * FROM competitions;")
	if err != nil {
		return nil, errors.Wrap(err, "Query")
	}
	return cmp, nil
}

func (cs CompetitionService) CompetitionsBySeason(season string) (cmp []scores.Competition, err error) {
	err = cs.client.db.Select(&cmp, "SELECT * FROM competitions WHERE comp_year=$1", season)
	if err != nil {
		return nil, errors.Wrap(err, "Query")
	}
	return cmp, nil
}

func (cs CompetitionService) CompetitionByID(id int64) (cmp *scores.Competition, err error) {
	cmp = new(scores.Competition)
	err = cs.client.db.QueryRowx("SELECT * FROM competitions WHERE comp_id=$1;", id).StructScan(cmp)
	if err != nil {
		return cmp, errors.Wrap(err, "Query")
	}
	return cmp, err
}

func (cs CompetitionService) AddCompetition(comp scores.Competition) (cmp *scores.Competition, err error) {
	const compInsert = "INSERT INTO competitions VALUES (DEFAULT, :created_at, :updated_at, :caption, :league, :comp_year, :matchday, :num_matchdays, :num_teams, :num_games) RETURNING comp_id;"

	tme := time.Now().Truncate(time.Millisecond)
	comp.CreatedAt = &tme
	comp.UpdatedAt = &tme
	qry, vars, err := sqlx.Named(compInsert, &comp)
	if err != nil {
		return nil, errors.Wrap(err, "Prepare")
	}
	qry = sqlx.Rebind(sqlx.BindType("postgres"), qry)
	err = cs.client.db.QueryRowx(qry, vars...).Scan(&comp.CompID)
	if err != nil {
		return nil, errors.Wrap(err, "Insert")
	}
	return &comp, nil
}
