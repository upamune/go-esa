package esa
import (
	"net/http"
	"net/url"
	"encoding/json"
	"io"
	"io/ioutil"
)

const (
	defaultBaseURL = "https://api.esa.io"
)

type Client struct {
	client *http.Client
	apiKey string
	baseURL string
	Team *TeamService
	Status *StatsService
	Post *PostService
	Comment *CommentService
}

func NewClient(apikey string) *Client {
	c := &Client{}
	c.client = http.DefaultClient
	c.apiKey = apikey
	c.baseURL = defaultBaseURL
	c.Team = &TeamService{client: c}
	c.Status = &StatsService{client: c}
	c.Post = &PostService{client: c}
	c.Comment = &CommentService{client: c}

	return c
}

func (c *Client) createURL(esaURL string) string {
	return c.baseURL + esaURL + "?access_token=" + c.apiKey
}

func (c *Client) post(esaURL string, bodyType string, body io.Reader, v interface{}) ( resp *http.Response, err error) {
	res, err := c.client.Post(c.createURL(esaURL), bodyType, body)
	if err != nil {
		return nil, err
	}

	if err := responseUnmarshal(res.Body, v); err != nil{
		return nil, err
	}

	return res, nil
}

func (c *Client) patch(esaURL string, bodyType string, body io.Reader, v interface{}) ( resp *http.Response, err error) {
	path := c.createURL(esaURL)
	req, err := http.NewRequest("PATCH", path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", bodyType)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := responseUnmarshal(res.Body, v); err != nil{
		return nil, err
	}

	return res, nil
}

func (c *Client) delete(esaURL string) ( resp *http.Response, err error) {
	path := c.createURL(esaURL)
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) get(esaURL string, query url.Values, v interface{}) (resp *http.Response, err error) {
	path := c.createURL(esaURL)
	for key, value := range query {
		path += "&" + key + "=" + url.QueryEscape(value[0])
	}

	res, err := c.client.Get(path)
	if err != nil {
		return nil, err
	}

	if err := responseUnmarshal(res.Body, v); err != nil {
		return nil, err
	}

	return res, err
}

func responseUnmarshal(body io.ReadCloser, v interface{}) error {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
