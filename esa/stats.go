package esa


const (
	// StatsURL esa API の統計情報のベ-スURL
	StatsURL = "/v1/teams"
)

// StatsService API docs: https://docs.esa.io/posts/102#5-0-0
type StatsService struct {
	client *Client
}

// StatsResponse 統計情報のレスポンス
type StatsResponse struct {
	Comments           int `json:"comments"`
	DailyActiveUsers   int `json:"daily_active_users"`
	Members            int `json:"members"`
	MonthlyActiveUsers int `json:"monthly_active_users"`
	Posts              int `json:"posts"`
	Stars              int `json:"stars"`
	WeeklyActiveUsers  int `json:"weekly_active_users"`
}

// GetTeamStats チ-ム名を指定して統計情報を取得する
func (s *StatsService) Get(teamName string) (*StatsResponse, error) {
	var statsRes StatsResponse

	statsURL := StatsURL + "/" + teamName + "/stats"
	res, err := s.client.get(statsURL, nil, &statsRes)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &statsRes, nil
}
