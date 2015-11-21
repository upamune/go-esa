package esa

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"
)

const (
	// CommentURL esa API のコメントのベ-スURL
	CommnetURL = "/v1/teams"
)

// CommentService API docs: https://docs.esa.io/posts/102#7-0-0
type CommentService struct {
	client *Client
}

// CommentResponse コメントのレスポンス
type CommentResponse struct {
	BodyHTML  string `json:"body_html"`
	BodyMd    string `json:"body_md"`
	CreatedAt string `json:"created_at"`
	CreatedBy struct {
		Icon       string `json:"icon"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
	} `json:"created_by"`
	ID        int    `json:"id"`
	UpdatedAt string `json:"updated_at"`
	URL       string `json:"url"`
}

// CommentsResponse 複数コメントのレスポンス
type CommentsResponse struct {
	Comments   []CommentResponse `json:"comments"`
	NextPage   interface{}       `json:"next_page"`
	PrevPage   interface{}       `json:"prev_page"`
	TotalCount int               `json:"total_count"`
}

// CommentReq コメントのリクエスト
type CommentReq struct {
	Comment Comment `json:"comment"`
}

// Comment コメント
type Comment struct {
	BodyMd string `json:"body_md"`
	User   string `json:"user"`
}

// GetTeamPostComments チ-ム名と記事番号を指定してコメントを取得する.
func (c *CommentService) GetComments(teamName string, postNumber int) (*CommentsResponse, error) {
	var commentsResponse CommentsResponse
	postNumberStr := strconv.Itoa(postNumber)
	commentURL := CommnetURL + "/" + teamName + "/posts" + "/" + postNumberStr + "/comments"

	res, err := c.client.get(commentURL, url.Values{}, &commentsResponse)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &commentsResponse, nil
}

// GetTeamComment チ-ム名とコメントIDを取得してコメントを取得する.
func (c *CommentService) GetComment(teamName string, commentID int) (*CommentResponse, error) {
	var commentResponse CommentResponse
	commentIDStr := strconv.Itoa(commentID)
	commentURL := CommnetURL + "/" + teamName + "/comments" + "/" + commentIDStr

	res, err := c.client.get(commentURL, url.Values{}, &commentResponse)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &commentResponse, nil
}

// PostTeamPostComment チ-ム名と記事番号とコメントを指定してコメントを投稿する
func (c *CommentService) Create(teamName string, postNumber int, comment Comment) (*CommentResponse, error) {
	postNumberStr := strconv.Itoa(postNumber)
	commentURL := CommnetURL + "/" + teamName + "/posts" + "/" + postNumberStr + "/comments"
	var commentResponse CommentResponse
	var commentReq CommentReq
	commentReq.Comment = comment

	var data []byte
	var err error
	if data, err = json.Marshal(commentReq); err != nil {
		return nil, err
	}

	res, err := c.client.post(commentURL, "application/json", bytes.NewReader(data), &commentResponse)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &commentResponse, nil
}

// PatchTeamComment チ-ム名とコメントIDとコメントを指定してコメントを更新する
func (c *CommentService) Update(teamName string, commentID int, comment Comment) (*CommentResponse, error) {
	commentIDStr := strconv.Itoa(commentID)
	commentURL := CommnetURL + "/" + teamName + "/comments" + "/" + commentIDStr
	var commentResponse CommentResponse
	var commentReq CommentReq
	commentReq.Comment = comment

	var data []byte
	var err error
	if data, err = json.Marshal(commentReq); err != nil {
		return nil, err
	}

	res, err := c.client.patch(commentURL, "application/json", bytes.NewReader(data), &commentResponse)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &commentResponse, nil
}

// DeleteTeamComment チ-ム名とコメントIDを指定してコメントを削除する
func (c *CommentService) Delete(teamName string, commentID int) error {
	commentIDStr := strconv.Itoa(commentID)
	commentURL := CommnetURL + "/" + teamName + "/comments" + "/" + commentIDStr

	res, err := c.client.delete(commentURL)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}
