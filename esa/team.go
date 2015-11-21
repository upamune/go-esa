package esa

import "net/url"

const (
	TeamURL = "/v1/teams"
)

type TeamService struct {
	client *Client
}

type TeamResponse struct {
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Privacy     string `json:"privacy"`
	URL         string `json:"url"`
}

type TeamsResponse struct {
	Teams      []TeamResponse `json:"teams"`
	PrevPage   interface{}    `json:"prev_page"`
	NextPage   interface{}    `json:"next_page"`
	TotalCount int            `json:"total_count"`
}

func (t *TeamService) GetTeams() (*TeamsResponse, error) {
	var teamsRes TeamsResponse
	res, err := t.client.get(TeamURL, url.Values{}, &teamsRes)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &teamsRes, nil
}

func (t *TeamService) GetTeam(teamName string) (*TeamResponse, error) {
	var teamRes TeamResponse
	teamURL := TeamURL + "/" + teamName
	res, err := t.client.get(teamURL, url.Values{}, &teamRes)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &teamRes, nil
}
