package esa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

func TestCommentGetTeamPostComments(t *testing.T) {
	type TestCase struct {
		in  string
		out CommentsResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/comment_get_comments.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Comment.GetTeamPostComments("docs", 2)
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s", res, testCase.out)
	}
}

func TestCommentGetTeamComment(t *testing.T) {
	type TestCase struct {
		in  string
		out CommentResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/comment_get_comment.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Comment.GetTeamComment("docs", 13)
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s", res, testCase.out)
	}
}

func TestCommentPostTeamPostComment(t *testing.T) {
	type TestCase struct {
		in  string
		out CommentResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/comment_post_response.json",
	}

	var comment Comment
	fileName := "../tests/stubs/comment_post_request.json"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(data, &comment); err != nil {
		log.Fatalln(err)
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Comment.PostTeamPostComment("docs", 2, comment)

	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s", res, testCase.out)
	}
}

func TestCommentPatchTeamComment(t *testing.T) {
	type TestCase struct {
		in  string
		out CommentResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/comment_patch_response.json",
	}

	var comment Comment
	fileName := "../tests/stubs/comment_patch_request.json"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(data, &comment); err != nil {
		log.Fatalln(err)
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Comment.PatchTeamComment("docs", 22767, comment)

	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %s != %s", res, testCase.out)
	}
}

func TestCommentDeleteComment(t *testing.T) {
	type TestCase struct {
		in  string
		out interface{}
	}

	testCase := TestCase{
		in: "../tests/stubs/comment_delete.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	err := client.Comment.DeleteTeamComment("docs", 22767)

	if err != nil {
		t.Errorf("error Request %s\n", err)
	}
}
