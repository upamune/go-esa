package esa

import "net/url"

const (
// MembersURL esa API のメンバーのベ-スURL
	MembersURL = "/v1/teams"
)

// MembersService API docs: https://docs.esa.io/posts/102#6-0-0
type MembersService struct {
	client *Client
}

// Member メンバー情報
type Member struct {
	Email      string `json:"email"`
	Icon       string `json:"icon"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

// MembersResponse メンバー情報のレスポンス
type MembersResponse struct {
	Members []Member `json:"members"`
	NextPage   interface{} `json:"next_page"`
	PrevPage   interface{} `json:"prev_page"`
	TotalCount int         `json:"total_count"`
}

// GetTeamMembers チ-ム名を指定してメンバー情報を取得する
func (s *MembersService) Get(teamName string) (*MembersResponse, error) {
	var membersRes MembersResponse

	membersURL := MembersURL+ "/" + teamName + "/members"
	res, err := s.client.get(membersURL, url.Values{}, &membersRes)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &membersRes, nil
}
