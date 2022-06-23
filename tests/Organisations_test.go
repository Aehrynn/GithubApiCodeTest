package tests

import (
	"GithubApiCodeTest/Main/v2/GithubApi"
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestWhenOrganisationsGetIsSuccessful(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.github.com/orgs/testorg",
		httpmock.NewStringResponder(200, `{
  "login": "testorg",
  "id": 2315026,
  "node_id": "MDEyOk9yZ2FuaXphdGlvbjIzMTUwMjY=",
  "url": "https://api.github.com/orgs/testorg",
  "repos_url": "https://api.github.com/orgs/testorg/repos",
  "events_url": "https://api.github.com/orgs/testorg/events",
  "hooks_url": "https://api.github.com/orgs/testorg/hooks",
  "issues_url": "https://api.github.com/orgs/testorg/issues",
  "members_url": "https://api.github.com/orgs/testorg/members{/member}",
  "public_members_url": "https://api.github.com/orgs/testorg/public_members{/member}",
  "avatar_url": "https://avatars.githubusercontent.com/u/2315026?v=4",
  "description": null,
  "is_verified": false,
  "has_organization_projects": true,
  "has_repository_projects": true,
  "public_repos": 2,
  "public_gists": 0,
  "followers": 0,
  "following": 0,
  "html_url": "https://github.com/testorg",
  "created_at": "2012-09-10T05:15:35Z",
  "updated_at": "2016-02-27T03:27:25Z",
  "type": "Organization"
}
`))
	organisation, err := GithubApi.GetOrganisation("testorg", testConfig)

	if err != nil {
		t.Errorf("expected no error response")
	}

	if organisation.URL != "https://api.github.com/orgs/testorg" {
		t.Errorf("expected `https://api.github.com/orgs/testorg`, got `" + organisation.URL + "`")
	}

}
