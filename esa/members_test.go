package esa

import (
	"reflect"
	"testing"
)

func TestMembersGet(t *testing.T) {
	type TestCase struct {
		in  string
		out MembersResponse
	}

	testCase := TestCase{
		in: "../tests/stubs/members_get.json",
	}

	serve, client := Stub(testCase.in, &testCase.out)
	defer serve.Close()

	res, err := client.Members.Get("esa")
	if err != nil {
		t.Errorf("error Request %s\n", err)
	}

	if !reflect.DeepEqual(*res, testCase.out) {
		t.Errorf("error Response %+v != %+v", res, testCase.out)
	}
}
