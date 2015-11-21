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
# Initialization
client := esa.NewClient("access_token")

# Team API
client.Team.GetTeams()
#=> GET /v1/teams

client.Team.GetTeam("bar")
#=> GET /v1/teams/bar

client.Stats.GetTeamStats("bar")
#=> GET /v1/teams/bar/stats

# Post API
client.Post.GetTeamPosts("foo")
#=> GET /v1/teams/foo/posts

client.Post.GetTeamPost("foo", 1)
#=> GET /v1/teams/foobar/posts/1

var post esa.Post{}
client.Post.PostTeamPost("foobar", post)
#=> POST /v1/teams/foobar/posts

client.Post.PatchTeamPost("foobar", 1, post)
#=> PATCH /v1/teams/foobar/posts/1

client.Post.DeleteTeamPost("foobar", 1)
#=> DELETE /v1/teams/foobar/posts/1


# Comment API
client.Comment.GetTeamPostComments("foobar", 1)
#=> GET /v1/teams/foobar/posts/1/comments

var comment esa.Comment
client.Comment.PostTeamPostComment("foobar", 1, comment)
#=> POST /v1/teams/foobar/posts/1/comments

client.Comment.GetTeamComment("foobar", 123)
#=> GET /v1/teams/foobar/comments/123

client.Comment.PatchTeamComment("foobar", 123, comment)
#=> PATCH /v1/teams/foobar/comments/123

client.Comment.DeleteTeamComment("foobar", 123)
#=> DELETE /v1/teams/foobar/comments/123
```