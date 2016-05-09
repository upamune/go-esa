# go-esa

[![Circle CI](https://circleci.com/gh/upamune/go-esa/tree/master.svg?style=svg)](https://circleci.com/gh/upamune/go-esa/tree/master)
[![Coverage Status](https://coveralls.io/repos/upamune/go-esa/badge.svg?branch=master&service=github)](https://coveralls.io/github/upamune/go-esa?branch=master)
[![GoDoc](https://godoc.org/github.com/upamune/go-esa?status.svg)](https://godoc.org/github.com/upamune/go-esa)

esa API v1 client library, written in Golang

## Install

```
go get github.com/upamune/go-esa
```

## Usage

```go
// Initialization
client := esa.NewClient("access_token")

// Team API
client.Team.GetTeams()
// => GET /v1/teams

client.Team.GetTeam("bar")
// => GET /v1/teams/bar

// Stats API
client.Stats.Get("bar")
// => GET /v1/teams/bar/stats

// Post API
client.Post.GetPosts("foo")
// => GET /v1/teams/foo/posts

query := url.Values{}
query.Add("in", "help")
client.Post.GetPosts("foo", query)
// => GET /v1/teams/foo/posts?q=in%3Ahelp

client.Post.GetPost("foo", 1)
// => GET /v1/teams/foobar/posts/1

var post esa.Post
client.Post.Create("foobar", post)
// => POST /v1/teams/foobar/posts

client.Post.Update("foobar", 1, post)
// => PATCH /v1/teams/foobar/posts/1

client.Post.Delete("foobar", 1)
// => DELETE /v1/teams/foobar/posts/1

client.Post.CreateSharing("foobar", 1)
// => POST /v1/teams/foobar/posts/1/sharing

client.Post.DeleteSharing("foobar", 1)
// => DELETE /v1/teams/foobar/posts/1/sharing

// Comment API
client.Comment.GetComments("foobar", 1)
// => GET /v1/teams/foobar/posts/1/comments

client.Comment.GetComment("foobar", 123)
// => GET /v1/teams/foobar/comments/123

var comment esa.Comment
client.Comment.Create("foobar", 1, comment)
// => POST /v1/teams/foobar/posts/1/comments

client.Comment.Update("foobar", 123, comment)
// => PATCH /v1/teams/foobar/comments/123

client.Comment.Delete("foobar", 123)
// => DELETE /v1/teams/foobar/comments/123

// Members API
client.Members.Get("foo")
// => GET /v1/teams/foo/members

// Watch API
client.Post.Watchers(1)
// => GET /v1/teams/foobar/posts/1/watchers

client.Post.AddWatch(1)
// => POST /v1/teams/foobar/posts/1/watch

client.Post.DeleteWatch(1)
// => DELETE /v1/teams/foobar/posts/1/watch
```
