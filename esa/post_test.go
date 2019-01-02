package esa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"reflect"
	"testing"
)

func TestPostGetPosts(t *testing.T) {
	type TestCase struct {
		in  string
		out PostsResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/post_posts.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Post.GetPosts("docs", url.Values{})
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %+v != %+v", res, testCase.out)
	}
}

func TestPostGetPost(t *testing.T) {
	type TestCase struct {
		in  string
		out PostResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/post_post.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Post.GetPost("docs", 1)
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %+v != %+v", res, testCase.out)
	}
}

func TestPostCreate(t *testing.T) {
	type TestCase struct {
		in  string
		out PostResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/post_post_request_response.json",
	}

	var post Post
	fileName := "../tests/stubs/post_post_request.json"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(data, &post); err != nil {
		log.Fatalln(err)
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Post.Create("docs", post)

	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %+v != %+v", res, testCase.out)
	}
}

func TestPostUpdate(t *testing.T) {
	type TestCase struct {
		in  string
		out PostResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/post_patch_response.json",
	}

	var post Post
	fileName := "../tests/stubs/post_patch_request.json"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(data, &post); err != nil {
		log.Fatalln(err)
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Post.Update("docs", 5, post)

	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %+v != %+v", res, testCase.out)
	}
}

func TestPostDelete(t *testing.T) {
	type TestCase struct {
		in  string
		out interface{}
	}

	testCase := TestCase{
		in: "../tests/stubs/post_delete.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	err := client.Post.Delete("docs", 5)

	if err != nil {
		t.Errorf("error Request %s\n", err)
	}
}

func TestCretaeSharing(t *testing.T) {
	type TestCase struct {
		in  string
		out SharedPost
	}

	testCase := TestCase{
		in: "../tests/stubs/post_sharing_create.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Post.CreateSharing("docs", 5)

	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s", res, testCase.out)
	}
}

func TestDeleteSharing(t *testing.T) {
	type TestCase struct {
		in  string
		out interface{}
	}

	testCase := TestCase{
		in: "../tests/stubs/post_sharing_delete.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	err := client.Post.DeleteSharing("docs", 5)

	if err != nil {
		t.Errorf("error Request %s\n", err)
	}
}

func Test_createSearchQuery(t *testing.T) {
	type TestCase struct {
		in  url.Values
		out string
	}

	testCases := []TestCase{
		{
			in: url.Values{
				"": []string{"esa"},
			},
			out: "esa",
		},
		{
			in: url.Values{
				"": []string{"esa", "docs"},
			},
			out: "esa docs",
		},
		{
			in: url.Values{
				"body": []string{"esa"},
			},
			out: "body:esa",
		},
		{
			in: url.Values{
				"body": []string{"esa", "docs"},
			},
			out: "body:esa body:docs",
		},
		{
			in: url.Values{
				"":     []string{"esa", "docs"},
				"body": []string{"start"},
			},
			out: "esa docs body:start",
		},
	}

	for _, ts := range testCases {
		searchQuery := createSearchQuery(ts.in)
		if ts.out != searchQuery {
			t.Errorf("error searchQuery [%s] != [%s]", searchQuery, ts.out)
		}
	}
}
