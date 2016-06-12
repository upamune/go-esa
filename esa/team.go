package esa

import (
)

const (
	// TeamURL esa API のチ-ムのベ-スURL
	TeamURL = "/v1/teams"
)

// TeamService API docs: https://docs.esa.io/posts/102#4-0-0
type TeamService struct {
	client *Client
}

// TeamResponse チ-ムのレスポンス
type TeamResponse struct {
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Privacy     string `json:"privacy"`
	URL         string `json:"url"`
}

// TeamsRespons 複数チ-ムのレスポンス
type TeamsResponse struct {
	Teams      []TeamResponse `json:"teams"`
	PrevPage   interface{}    `json:"prev_page"`
	NextPage   interface{}    `json:"next_page"`
	TotalCount int            `json:"total_count"`
}

// GetTeams チ-ムを取得する
func (t *TeamService) GetTeams() (*TeamsResponse, error) {
	var teamsRes TeamsResponse
	_, err := t.client.get(TeamURL, nil, &teamsRes)
	if err != nil {
		return nil, err
	}

	return &teamsRes, nil
}

// GetTeam チ-ム名を取得してチ-ムを取得する
func (t *TeamService) GetTeam(teamName string) (*TeamResponse, error) {
	var teamRes TeamResponse
	teamURL := TeamURL + "/" + teamName
	_, err := t.client.get(teamURL, nil, &teamRes)
	if err != nil {
		return nil, err
	}

	return &teamRes, nil
}
