package esa

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"
)

const (
	// PostURL esa API のコメントのベ-スURL
	PostURL = "/v1/teams"
)

// PostService API docs: https://docs.esa.io/posts/102#6-0-0
type PostService struct {
	client *Client
}

// PostReq 記事のリクエスト
type PostReq struct {
	Post Post `json:"post"`
}

// Post 記事
type Post struct {
	BodyMd   string   `json:"body_md"`
	Category string   `json:"category"`
	Message  string   `json:"message"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags"`
	Wip      bool     `json:"wip"`
}

// PostResponse 記事のレスポンス
type PostResponse struct {
	BodyHTML      string `json:"body_html"`
	BodyMd        string `json:"body_md"`
	Category      string `json:"category"`
	CommentsCount int    `json:"comments_count"`
	CreatedAt     string `json:"created_at"`
	CreatedBy     struct {
		Icon       string `json:"icon"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
	} `json:"created_by"`
	DoneTasksCount  int      `json:"done_tasks_count"`
	FullName        string   `json:"full_name"`
	Kind            string   `json:"kind"`
	Message         string   `json:"message"`
	Name            string   `json:"name"`
	Number          int      `json:"number"`
	OverLapped      bool     `json:"overlapped"`
	RevisionNumber  int      `json:"revision_number"`
	Star            bool     `json:"star"`
	StargazersCount int      `json:"stargazers_count"`
	Tags            []string `json:"tags"`
	TasksCount      int      `json:"tasks_count"`
	UpdatedAt       string   `json:"updated_at"`
	UpdatedBy       struct {
		Icon       string `json:"icon"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
	} `json:"updated_by"`
	URL           string `json:"url"`
	Watch         bool   `json:"watch"`
	WatchersCount int    `json:"watchers_count"`
	Wip           bool   `json:"wip"`
}

// PostsResponse 複数記事のレスポンス
type PostsResponse struct {
	NextPage   interface{}    `json:"next_page"`
	Posts      []PostResponse `json:"posts"`
	PrevPage   interface{}    `json:"prev_page"`
	TotalCount int            `json:"total_count"`
}

// SharedPost 公開された記事
type SharedPost struct {
	HTML   string `json:"html"`
	Slides string `json:"slides"`
}

func createSearchQuery(query url.Values) string {
	var queries string
	for key, values := range query {
		queries += key + ":"
		for _, value := range values {
			queries += value + " "
		}
		queries += "+"
	}

	return queries
}

// GetTeamPosts チ-ム名とクエリを指定して記事を取得する
func (p *PostService) GetPosts(teamName string, query url.Values) (*PostsResponse, error) {
	var postsRes PostsResponse
	queries := createSearchQuery(query)

	searchQuery := url.Values{}
	searchQuery.Add("q", queries)
	searchQuery.Encode()

	postsURL := PostURL + "/" + teamName + "/posts"
	res, err := p.client.get(postsURL, searchQuery, &postsRes)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &postsRes, nil

}

// GetTeamPost チ-ム名と記事番号を指定して記事を取得する
func (p *PostService) GetPost(teamName string, postNumber int) (*PostResponse, error) {
	var postRes PostResponse

	postNumberStr := strconv.Itoa(postNumber)

	postURL := PostURL + "/" + teamName + "/posts" + "/" + postNumberStr
	res, err := p.client.get(postURL, url.Values{}, &postRes)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &postRes, nil
}

// PostTeamPost チ-ム名と記事を指定して記事を投稿する
func (p *PostService) Create(teamName string, post Post) (*PostResponse, error) {
	postURL := PostURL + "/" + teamName + "/posts"
	var postRes PostResponse
	var postReq PostReq
	postReq.Post = post
	var data []byte
	var err error
	if data, err = json.Marshal(postReq); err != nil {
		return nil, err
	}

	res, err := p.client.post(postURL, "application/json", bytes.NewReader(data), &postRes)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &postRes, nil
}

// PatchTeamPost チ-ム名と記事番号と記事を指定して記事を更新する
func (p *PostService) Update(teamName string, postNumber int, post Post) (*PostResponse, error) {
	var postRes PostResponse
	var postReq PostReq
	postReq.Post = post
	postNumberStr := strconv.Itoa(postNumber)
	postURL := PostURL + "/" + teamName + "/posts" + "/" + postNumberStr

	var data []byte
	var err error
	if data, err = json.Marshal(postReq); err != nil {
		return nil, err
	}

	res, err := p.client.patch(postURL, "application/json", bytes.NewReader(data), &postRes)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &postRes, nil
}

// DeleteTeamPost チ-ム名と記事番号を指定して記事を削除する
func (p *PostService) Delete(teamName string, postNumber int) error {
	postNumberStr := strconv.Itoa(postNumber)
	postURL := PostURL + "/" + teamName + "/posts" + "/" + postNumberStr

	res, err := p.client.delete(postURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// CreateSharing チ-ム名と記事を指定して記事を投稿する
func (p *PostService) CreateSharing(teamName string, postNumber int) (*SharedPost, error) {
	postNumberStr := strconv.Itoa(postNumber)
	postURL := PostURL + "/" + teamName + "/posts" + "/" + postNumberStr + "/sharing"
	var sharedRes SharedPost
	var postReq PostReq
	var data []byte
	var err error
	if data, err = json.Marshal(postReq); err != nil {
		return nil, err
	}

	res, err := p.client.post(postURL, "application/json", bytes.NewReader(data), &sharedRes)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &sharedRes, nil
}

// DeleteSharing チ-ム名と記事番号を指定して記事を削除する
func (p *PostService) DeleteSharing(teamName string, postNumber int) error {
	postNumberStr := strconv.Itoa(postNumber)
	postURL := PostURL + "/" + teamName + "/posts" + "/" + postNumberStr + "/sharing"

	res, err := p.client.delete(postURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
