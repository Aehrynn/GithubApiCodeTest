package tests

import (
	"GithubApiCodeTest/Main/v2/GithubApi"
	"GithubApiCodeTest/Main/v2/constants"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestWhenHttpClientError(t *testing.T) {
	constants.NewHttpRequest = func(method, url string, body io.Reader) (*http.Request, error) {
		return nil, fmt.Errorf("an error")
	}
	_, err := GithubApi.GetRepositories("testOrg")
	if err == nil {
		t.Errorf("expected an error response")
	}
	//if err != fmt.Errorf("an error") {
	//	t.Errorf("Not expected Error")
	//}
}
