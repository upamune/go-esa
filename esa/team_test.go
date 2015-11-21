package esa

import (
	"reflect"
	"testing"
)

func TestTeamGetTeams(t *testing.T) {
	type TestCase struct {
		in  string
		out TeamsResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/team_teams.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Team.GetTeams()
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s", res, testCase.out)
	}
}

func TestTeamGetTeam(t *testing.T) {
	type TestCase struct {
		in  string
		out TeamResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/team_team.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Team.GetTeam("docs")
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s", res, testCase.out)
	}
}
