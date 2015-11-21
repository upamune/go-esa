package esa
import "net/url"

const (
	StatsURL = "/v1/teams"
)

type StatsService struct {
	client *Client
}

type StatsResponse struct {
	Comments           int `json:"comments"`
	DailyActiveUsers   int `json:"daily_active_users"`
	Members            int `json:"members"`
	MonthlyActiveUsers int `json:"monthly_active_users"`
	Posts              int `json:"posts"`
	Stars              int `json:"stars"`
	WeeklyActiveUsers  int `json:"weekly_active_users"`
}


func (s *StatsService) GetTeamStats(teamName string) (*StatsResponse, error) {
	var statsRes StatsResponse

	statsURL := StatsURL + "/" + teamName + "/stats"
	res, err := s.client.get(statsURL, url.Values{}, &statsRes)
	if err != nil {
		return  nil, err
	}

	defer res.Body.Close()

	return &statsRes, nil
}

