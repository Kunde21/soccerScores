package postgres

import (
	"fmt"
	"os"
	"testing"

	scores "github.com/Kunde21/scoresServer"
	"github.com/pkg/errors"
)

func init() {
	fmt.Println("TestDB:", os.Getenv("TEST_DB"))
}

func compareComps(a, b scores.Competition) error {
	switch {
	case a.CompID != b.CompID:
		return errors.Errorf("Different CompID\n%v\n%v\n", a, b)
	case a.Caption != b.Caption:
		return errors.Errorf("Different Caption\n%v\n%v\n", a, b)
	case a.League != b.League:
		return errors.Errorf("Different League\n%v\n%v\n", a, b)
	case a.CompYear != b.CompYear:
		return errors.Errorf("Different CompYear\n%v\n%v\n", a, b)
	case a.Matchday != b.Matchday:
		return errors.Errorf("Different Matchday\n%v\n%v\n", a, b)
	case a.NumMatchdays != b.NumMatchdays:
		return errors.Errorf("Different NumMatchdays\n%v\n%v\n", a, b)
	case a.NumTeams != b.NumTeams:
		return errors.Errorf("Different NumTeams\n%v\n%v\n", a, b)
	case a.NumGames != b.NumGames:
		return errors.Errorf("Different NumGames\n%v\n%v\n", a, b)
	}
	return nil
}

func testClient(t *testing.T) *Client {
	client := NewClient()
	var err error

	client.db, err = testDB()
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func TestAddCompetition(t *testing.T) {
	client := testClient(t)
	cs := client.CompetitionService()
	cmp, err := cs.AddCompetition(scores.Competition{
		Caption:      "Test",
		League:       "TstPro",
		CompYear:     "1900-1901",
		Matchday:     20,
		NumMatchdays: 25,
		NumTeams:     20,
		NumGames:     250,
	})
	if err != nil {
		t.Error(err)
	}
	if cmp.CompID == 0 {
		t.Fatalf("CompID not set on Insert:\n%+v\n", cmp)
	}
	cmpTst, err := cs.CompetitionByID(cmp.CompID)
	if err != nil {
		t.Error(err)
	}
	if err = compareComps(*cmp, *cmpTst); err != nil {
		t.Error(err)
	}

	dbCleanup(t, client.db)
}

var compsTest = []scores.Competition{
	{
		Caption:      "Test",
		League:       "TstPro",
		CompYear:     "1900-1901",
		Matchday:     20,
		NumMatchdays: 25,
		NumTeams:     20,
		NumGames:     250,
	},
	{
		Caption:      "Test2",
		League:       "TstPro",
		CompYear:     "1900-1901",
		Matchday:     15,
		NumMatchdays: 20,
		NumTeams:     20,
		NumGames:     250,
	},
	{
		Caption:      "Test3",
		League:       "TstAmtr",
		CompYear:     "1950-1951",
		Matchday:     20,
		NumMatchdays: 25,
		NumTeams:     20,
		NumGames:     250,
	},
}

func TestCompetitions(t *testing.T) {
	client := testClient(t)
	defer dbCleanup(t, client.db)
	cs := client.CompetitionService()
	var err error
	cmpTst := make([]scores.Competition, len(compsTest))

	for i := range compsTest {
		cmp, err := cs.AddCompetition(compsTest[i])
		if err != nil {
			t.Fatal("Insert", i, err)
		}
		cmpTst[i] = *cmp
	}

	cmpRes, err := cs.Competitions()
	if err != nil {
		t.Fatal("Competitions", err)
	}

find:
	for i := range cmpTst {
		for j := range cmpRes {
			if compareComps(cmpTst[i], cmpRes[j]) == nil {
				continue find
			}
		}
		t.Errorf("Competitions did not retrieve %d:\n %v", i, cmpTst[i])
	}
}

func TestCompetitionsBySeason(t *testing.T) {
	client := testClient(t)
	defer dbCleanup(t, client.db)
	cs := client.CompetitionService()
	var err error
	cmpTst := make([]scores.Competition, len(compsTest))

	for i := range compsTest {
		cmp, err := cs.AddCompetition(compsTest[i])
		if err != nil {
			t.Fatal("Insert", i, err)
		}
		cmpTst[i] = *cmp
	}

	cmpRes, err := cs.CompetitionsBySeason("1900-1901")
	if err != nil {
		t.Fatal("CompetitionsBySeason", err)
	}
find:
	for i := range cmpTst[:2] {
		for j := range cmpRes {
			if compareComps(cmpTst[i], cmpRes[j]) == nil {
				continue find
			}
		}
		t.Errorf("CompetitionsBySeason did not retrieve %d:\n %v", i, cmpTst[i])
	}
}
