//go:generate kallax gen --output postgres/kallax.go

package scores

import "time"

// Stat is an enum type that shows the status of a Fixture
type Stat uint8

// Status flags
const (
	SCHED Stat = iota
	POSTP
	INPROG
	FINISHED
)

// Competition represents a single league or tourney.
type Competition struct {
	CreatedAt    *time.Time `db:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at"`
	CompID       int64      `db:"comp_id"`
	Caption      string     `db:"caption"`
	League       string     `db:"league"`
	CompYear     string     `db:"comp_year"`
	Matchday     int8       `db:"matchday"`
	NumMatchdays int8       `db:"num_matchdays"`
	NumTeams     int8       `db:"num_teams"`
	NumGames     int16      `db:"num_games"`
}

// Team represents a single club soccer team.
type Team struct {
	TeamID    int64  `db:"team_id"`
	TeamName  string `db:"team_name"`
	ShortName string `db:"short_name"`
	CrestURL  string `db:"crest_url"`
}

// Fixture represents a single scheduled match.
type Fixture struct {
	CreatedAt   *time.Time   `db:"created_at"`
	UpdatedAt   *time.Time   `db:"updated_at"`
	FixID       int64        `db:"fix_id"`
	CompID      *Competition `db:"comp_id"`
	StartTime   *time.Time   `db:"start_time"`
	EndTime     *time.Time   `db:"end_time"`
	Status      Stat         `db:"status"`
	Broadcast   string       `db:"broadcast"`
	HomeTeam    *Team        `db:"home_team"`
	AwayTeam    *Team        `db:"away_team"`
	HomeGoals   uint8        `db:"home_goals"`
	AwayGoals   uint8        `db:"away_goals"`
	HomeHTGoals uint8        `db:"home_ht_goals"`
	AwayHTGoals uint8        `db:"away_ht_goals"`
}

// Client is a backing store for persistent data
type Client interface {
	CompetitionService() CompetitionService
	TeamService() TeamService
	FixtureService() FixtureService
}

// CompetitionService handles storage/retrieval interactions on Competitions
type CompetitionService interface {
	Competitions() ([]Competition, error)
	CompetitionsBySeason(season string) ([]Competition, error)
	CompetitionByID(id int64) (*Competition, error)
	AddCompetition(comp Competition) (*Competition, error)
}

// TeamService handles storage/retrieval interactions on Teams
type TeamService interface {
	Teams() ([]Team, error)
	TeamByID(id int64) (*Team, error)
	AddTeam(tm Team) (*Team, error)
}

// FixtureService handles storage/retrieval interactions on Fixtures
type FixtureService interface {
	Fixtures() ([]Fixture, error)
	FixturesByComp(id int64) ([]Fixture, error)
	FixturesByTeam(id int64) ([]Fixture, error)
	FixtureByID(id int64) (*Fixture, error)
	AddFixture(fix Fixture) (*Fixture, error)
	UpdateFixture(fix Fixture) (*Fixture, error)
}
