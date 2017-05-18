package postgres

import scores "github.com/Kunde21/scoresServer"

var _ scores.FixtureService = FixtureService{}

type FixtureService struct {
	client *Client
}

func (fs FixtureService) Fixtures() (fx []scores.Fixture, err error) {
	panic("not implemented")
}

func (fs FixtureService) FixturesByComp(id int64) (fx []scores.Fixture, err error) {
	panic("not implemented")
}

func (fs FixtureService) FixturesByTeam(id int64) (fx []scores.Fixture, err error) {
	panic("not implemented")
}

func (fs FixtureService) FixtureByID(id int64) (fx *scores.Fixture, err error) {
	panic("not implemented")
}

func (fs FixtureService) AddFixture(fix scores.Fixture) (fx *scores.Fixture, err error) {
	panic("not implemented")
}

func (fs FixtureService) UpdateFixture(fix scores.Fixture) (fx *scores.Fixture, err error) {
	panic("not implemented")
}
