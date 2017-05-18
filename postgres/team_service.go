package postgres

import scores "github.com/Kunde21/scoresServer"

var _ scores.TeamService = TeamService{}

type TeamService struct {
	client *Client
}

func (ts TeamService) Teams() (tm []scores.Team, err error) {
	panic("not implemented")
}

func (ts TeamService) TeamByID(id int64) (tm *scores.Team, err error) {
	panic("not implemented")
}

func (ts TeamService) AddTeam(team scores.Team) (tm *scores.Team, err error) {
	panic("not implemented")
}
