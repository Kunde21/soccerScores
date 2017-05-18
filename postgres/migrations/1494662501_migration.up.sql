BEGIN TRANSACTION;

CREATE TABLE competitions (
	comp_id serial PRIMARY KEY,
	created_at timestamptz,
	updated_at timestamptz,
	caption text,
	league text,
	comp_year text,
	matchday smallint,
	num_matchdays smallint,
	num_teams smallint,
	num_games smallint
);


CREATE TABLE teams (
	team_id serial PRIMARY KEY,
	team_name text,
	short_name text,
	crest_url text
);


CREATE TABLE fixtures (
	fix_id serial PRIMARY KEY,
	created_at timestamptz,
	updated_at timestamptz,
	comp_id bigint REFERENCES competitions(comp_id),
	start_time timestamptz,
	end_time timestamptz,
	status smallint,
	broadcast text,
	home_team bigint REFERENCES teams(team_id),
	away_team bigint REFERENCES teams(team_id),
	home_goals smallint,
	away_goals smallint,
	home_htgoals smallint,
	away_htgoals smallint
);


COMMIT;
