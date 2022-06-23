package tests

import (
	"GithubApiCodeTest/Main/v2/GithubApi"
	"GithubApiCodeTest/Main/v2/config"
	"fmt"
	"github.com/jarcoal/httpmock"
	"strconv"
	"testing"
)

var testConfig = config.Configurations{}

//func TestWhenHttpClientError(t *testing.T) {
//	expectedError := "an error"
//	constants.NewHttpRequest = func(method, url string, body io.Reader) (*http.Request, error) {
//		return nil, fmt.Errorf(expectedError)
//	}
//	_, err := GithubApi.GetRepositories("testOrg", testConfig)
//
//	if err == nil {
//		t.Errorf("expected an error response")
//	}
//
//	if err.Error() != expectedError {
//		t.Errorf("expected an " + expectedError + " response")
//	}
//}

func TestWhenRepositoriesGetIsSuccessful(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.github.com/orgs/testOrg/repos",
		httpmock.NewStringResponder(200, `[
  {
    "id": 5745280,
    "node_id": "MDEwOlJlcG9zaXRvcnk1NzQ1Mjgw",
    "name": "publicrepo",
    "full_name": "testorg/publicrepo",
    "private": false,
    "owner": {
      "login": "testorg",
      "id": 2315026,
      "node_id": "MDEyOk9yZ2FuaXphdGlvbjIzMTUwMjY=",
      "avatar_url": "https://avatars.githubusercontent.com/u/2315026?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/testorg",
      "html_url": "https://github.com/testorg",
      "followers_url": "https://api.github.com/users/testorg/followers",
      "following_url": "https://api.github.com/users/testorg/following{/other_user}",
      "gists_url": "https://api.github.com/users/testorg/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/testorg/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/testorg/subscriptions",
      "organizations_url": "https://api.github.com/users/testorg/orgs",
      "repos_url": "https://api.github.com/users/testorg/repos",
      "events_url": "https://api.github.com/users/testorg/events{/privacy}",
      "received_events_url": "https://api.github.com/users/testorg/received_events",
      "type": "Organization",
      "site_admin": false
    },
    "html_url": "https://github.com/testorg/publicrepo",
    "description": null,
    "fork": false,
    "url": "https://api.github.com/repos/testorg/publicrepo",
    "forks_url": "https://api.github.com/repos/testorg/publicrepo/forks",
    "keys_url": "https://api.github.com/repos/testorg/publicrepo/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/testorg/publicrepo/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/testorg/publicrepo/teams",
    "hooks_url": "https://api.github.com/repos/testorg/publicrepo/hooks",
    "issue_events_url": "https://api.github.com/repos/testorg/publicrepo/issues/events{/number}",
    "events_url": "https://api.github.com/repos/testorg/publicrepo/events",
    "assignees_url": "https://api.github.com/repos/testorg/publicrepo/assignees{/user}",
    "branches_url": "https://api.github.com/repos/testorg/publicrepo/branches{/branch}",
    "tags_url": "https://api.github.com/repos/testorg/publicrepo/tags",
    "blobs_url": "https://api.github.com/repos/testorg/publicrepo/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/testorg/publicrepo/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/testorg/publicrepo/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/testorg/publicrepo/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/testorg/publicrepo/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/testorg/publicrepo/languages",
    "stargazers_url": "https://api.github.com/repos/testorg/publicrepo/stargazers",
    "contributors_url": "https://api.github.com/repos/testorg/publicrepo/contributors",
    "subscribers_url": "https://api.github.com/repos/testorg/publicrepo/subscribers",
    "subscription_url": "https://api.github.com/repos/testorg/publicrepo/subscription",
    "commits_url": "https://api.github.com/repos/testorg/publicrepo/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/testorg/publicrepo/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/testorg/publicrepo/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/testorg/publicrepo/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/testorg/publicrepo/contents/{+path}",
    "compare_url": "https://api.github.com/repos/testorg/publicrepo/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/testorg/publicrepo/merges",
    "archive_url": "https://api.github.com/repos/testorg/publicrepo/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/testorg/publicrepo/downloads",
    "issues_url": "https://api.github.com/repos/testorg/publicrepo/issues{/number}",
    "pulls_url": "https://api.github.com/repos/testorg/publicrepo/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/testorg/publicrepo/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/testorg/publicrepo/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/testorg/publicrepo/labels{/name}",
    "releases_url": "https://api.github.com/repos/testorg/publicrepo/releases{/id}",
    "deployments_url": "https://api.github.com/repos/testorg/publicrepo/deployments",
    "created_at": "2012-09-10T05:30:27Z",
    "updated_at": "2013-12-09T22:45:25Z",
    "pushed_at": "2012-09-11T03:26:38Z",
    "git_url": "git://github.com/testorg/publicrepo.git",
    "ssh_url": "git@github.com:testorg/publicrepo.git",
    "clone_url": "https://github.com/testorg/publicrepo.git",
    "svn_url": "https://github.com/testorg/publicrepo",
    "homepage": null,
    "size": 116,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": "PHP",
    "has_issues": true,
    "has_projects": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 0,
    "mirror_url": null,
    "archived": false,
    "disabled": false,
    "open_issues_count": 0,
    "license": null,
    "allow_forking": true,
    "is_template": false,
    "topics": [

    ],
    "visibility": "public",
    "forks": 0,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "master",
    "permissions": {
      "admin": false,
      "maintain": false,
      "push": false,
      "triage": false,
      "pull": true
    }
  },
  {
    "id": 5759417,
    "node_id": "MDEwOlJlcG9zaXRvcnk1NzU5NDE3",
    "name": "publicrepo2",
    "full_name": "testorg/publicrepo2",
    "private": false,
    "owner": {
      "login": "testorg",
      "id": 2315026,
      "node_id": "MDEyOk9yZ2FuaXphdGlvbjIzMTUwMjY=",
      "avatar_url": "https://avatars.githubusercontent.com/u/2315026?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/testorg",
      "html_url": "https://github.com/testorg",
      "followers_url": "https://api.github.com/users/testorg/followers",
      "following_url": "https://api.github.com/users/testorg/following{/other_user}",
      "gists_url": "https://api.github.com/users/testorg/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/testorg/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/testorg/subscriptions",
      "organizations_url": "https://api.github.com/users/testorg/orgs",
      "repos_url": "https://api.github.com/users/testorg/repos",
      "events_url": "https://api.github.com/users/testorg/events{/privacy}",
      "received_events_url": "https://api.github.com/users/testorg/received_events",
      "type": "Organization",
      "site_admin": false
    },
    "html_url": "https://github.com/testorg/publicrepo2",
    "description": null,
    "fork": false,
    "url": "https://api.github.com/repos/testorg/publicrepo2",
    "forks_url": "https://api.github.com/repos/testorg/publicrepo2/forks",
    "keys_url": "https://api.github.com/repos/testorg/publicrepo2/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/testorg/publicrepo2/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/testorg/publicrepo2/teams",
    "hooks_url": "https://api.github.com/repos/testorg/publicrepo2/hooks",
    "issue_events_url": "https://api.github.com/repos/testorg/publicrepo2/issues/events{/number}",
    "events_url": "https://api.github.com/repos/testorg/publicrepo2/events",
    "assignees_url": "https://api.github.com/repos/testorg/publicrepo2/assignees{/user}",
    "branches_url": "https://api.github.com/repos/testorg/publicrepo2/branches{/branch}",
    "tags_url": "https://api.github.com/repos/testorg/publicrepo2/tags",
    "blobs_url": "https://api.github.com/repos/testorg/publicrepo2/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/testorg/publicrepo2/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/testorg/publicrepo2/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/testorg/publicrepo2/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/testorg/publicrepo2/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/testorg/publicrepo2/languages",
    "stargazers_url": "https://api.github.com/repos/testorg/publicrepo2/stargazers",
    "contributors_url": "https://api.github.com/repos/testorg/publicrepo2/contributors",
    "subscribers_url": "https://api.github.com/repos/testorg/publicrepo2/subscribers",
    "subscription_url": "https://api.github.com/repos/testorg/publicrepo2/subscription",
    "commits_url": "https://api.github.com/repos/testorg/publicrepo2/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/testorg/publicrepo2/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/testorg/publicrepo2/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/testorg/publicrepo2/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/testorg/publicrepo2/contents/{+path}",
    "compare_url": "https://api.github.com/repos/testorg/publicrepo2/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/testorg/publicrepo2/merges",
    "archive_url": "https://api.github.com/repos/testorg/publicrepo2/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/testorg/publicrepo2/downloads",
    "issues_url": "https://api.github.com/repos/testorg/publicrepo2/issues{/number}",
    "pulls_url": "https://api.github.com/repos/testorg/publicrepo2/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/testorg/publicrepo2/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/testorg/publicrepo2/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/testorg/publicrepo2/labels{/name}",
    "releases_url": "https://api.github.com/repos/testorg/publicrepo2/releases{/id}",
    "deployments_url": "https://api.github.com/repos/testorg/publicrepo2/deployments",
    "created_at": "2012-09-11T03:28:27Z",
    "updated_at": "2013-10-01T15:42:25Z",
    "pushed_at": "2012-09-11T07:36:36Z",
    "git_url": "git://github.com/testorg/publicrepo2.git",
    "ssh_url": "git@github.com:testorg/publicrepo2.git",
    "clone_url": "https://github.com/testorg/publicrepo2.git",
    "svn_url": "https://github.com/testorg/publicrepo2",
    "homepage": null,
    "size": 116,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": null,
    "has_issues": true,
    "has_projects": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 0,
    "mirror_url": null,
    "archived": false,
    "disabled": false,
    "open_issues_count": 0,
    "license": null,
    "allow_forking": true,
    "is_template": false,
    "topics": [

    ],
    "visibility": "public",
    "forks": 0,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "master",
    "permissions": {
      "admin": false,
      "maintain": false,
      "push": false,
      "triage": false,
      "pull": true
    }
  }
]
`))
	repositories, err := GithubApi.GetRepositories("testOrg", testConfig)
	fmt.Println(err)
	if err != nil {
		t.Errorf("expected no error response")
	}

	if len(repositories) == 0 {
		t.Errorf("expected more than 0 repositories")
	}
}

func TestWhenReleasesGetIsSuccessful(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.github.com/repos/testOrg/testRepo/releases",
		httpmock.NewStringResponder(200, `[
  {
    "url": "https://api.github.com/repos/github/pong/releases/600374",
    "assets_url": "https://api.github.com/repos/github/pong/releases/600374/assets",
    "upload_url": "https://uploads.github.com/repos/github/pong/releases/600374/assets{?name,label}",
    "html_url": "https://github.com/github/pong/releases/tag/small.bin",
    "id": 600374,
    "author": {
      "login": "jnewland",
      "id": 47,
      "node_id": "MDQ6VXNlcjQ3",
      "avatar_url": "https://avatars.githubusercontent.com/u/47?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/jnewland",
      "html_url": "https://github.com/jnewland",
      "followers_url": "https://api.github.com/users/jnewland/followers",
      "following_url": "https://api.github.com/users/jnewland/following{/other_user}",
      "gists_url": "https://api.github.com/users/jnewland/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/jnewland/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/jnewland/subscriptions",
      "organizations_url": "https://api.github.com/users/jnewland/orgs",
      "repos_url": "https://api.github.com/users/jnewland/repos",
      "events_url": "https://api.github.com/users/jnewland/events{/privacy}",
      "received_events_url": "https://api.github.com/users/jnewland/received_events",
      "type": "User",
      "site_admin": false
    },
    "node_id": "MDc6UmVsZWFzZTYwMDM3NA==",
    "tag_name": "small.bin",
    "target_commitish": "master",
    "name": "small.bin",
    "draft": false,
    "prerelease": false,
    "created_at": "2011-02-20T01:36:54Z",
    "published_at": "2014-10-03T00:56:52Z",
    "assets": [
      {
        "url": "https://api.github.com/repos/github/pong/releases/assets/257265",
        "id": 257265,
        "node_id": "MDEyOlJlbGVhc2VBc3NldDI1NzI2NQ==",
        "name": "small.bin",
        "label": null,
        "uploader": {
          "login": "jnewland",
          "id": 47,
          "node_id": "MDQ6VXNlcjQ3",
          "avatar_url": "https://avatars.githubusercontent.com/u/47?v=4",
          "gravatar_id": "",
          "url": "https://api.github.com/users/jnewland",
          "html_url": "https://github.com/jnewland",
          "followers_url": "https://api.github.com/users/jnewland/followers",
          "following_url": "https://api.github.com/users/jnewland/following{/other_user}",
          "gists_url": "https://api.github.com/users/jnewland/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/jnewland/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/jnewland/subscriptions",
          "organizations_url": "https://api.github.com/users/jnewland/orgs",
          "repos_url": "https://api.github.com/users/jnewland/repos",
          "events_url": "https://api.github.com/users/jnewland/events{/privacy}",
          "received_events_url": "https://api.github.com/users/jnewland/received_events",
          "type": "User",
          "site_admin": false
        },
        "content_type": "application/macbinary",
        "state": "uploaded",
        "size": 10485760,
        "download_count": 139,
        "created_at": "2014-10-03T00:55:25Z",
        "updated_at": "2014-10-03T00:55:28Z",
        "browser_download_url": "https://github.com/github/pong/releases/download/small.bin/small.bin"
      }
    ],
    "tarball_url": "https://api.github.com/repos/github/pong/tarball/small.bin",
    "zipball_url": "https://api.github.com/repos/github/pong/zipball/small.bin",
    "body": "For release download speed testing\n"
  },
  {
    "url": "https://api.github.com/repos/github/pong/releases/371259",
    "assets_url": "https://api.github.com/repos/github/pong/releases/371259/assets",
    "upload_url": "https://uploads.github.com/repos/github/pong/releases/371259/assets{?name,label}",
    "html_url": "https://github.com/github/pong/releases/tag/large.bin",
    "id": 371259,
    "author": {
      "login": "jnewland",
      "id": 47,
      "node_id": "MDQ6VXNlcjQ3",
      "avatar_url": "https://avatars.githubusercontent.com/u/47?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/jnewland",
      "html_url": "https://github.com/jnewland",
      "followers_url": "https://api.github.com/users/jnewland/followers",
      "following_url": "https://api.github.com/users/jnewland/following{/other_user}",
      "gists_url": "https://api.github.com/users/jnewland/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/jnewland/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/jnewland/subscriptions",
      "organizations_url": "https://api.github.com/users/jnewland/orgs",
      "repos_url": "https://api.github.com/users/jnewland/repos",
      "events_url": "https://api.github.com/users/jnewland/events{/privacy}",
      "received_events_url": "https://api.github.com/users/jnewland/received_events",
      "type": "User",
      "site_admin": false
    },
    "node_id": "MDc6UmVsZWFzZTM3MTI1OQ==",
    "tag_name": "large.bin",
    "target_commitish": "master",
    "name": "large.bin",
    "draft": false,
    "prerelease": false,
    "created_at": "2011-02-20T01:36:54Z",
    "published_at": "2014-06-12T15:25:44Z",
    "assets": [
      {
        "url": "https://api.github.com/repos/github/pong/releases/assets/157415",
        "id": 157415,
        "node_id": "MDEyOlJlbGVhc2VBc3NldDE1NzQxNQ==",
        "name": "large.bin",
        "label": null,
        "uploader": {
          "login": "jnewland",
          "id": 47,
          "node_id": "MDQ6VXNlcjQ3",
          "avatar_url": "https://avatars.githubusercontent.com/u/47?v=4",
          "gravatar_id": "",
          "url": "https://api.github.com/users/jnewland",
          "html_url": "https://github.com/jnewland",
          "followers_url": "https://api.github.com/users/jnewland/followers",
          "following_url": "https://api.github.com/users/jnewland/following{/other_user}",
          "gists_url": "https://api.github.com/users/jnewland/gists{/gist_id}",
          "starred_url": "https://api.github.com/users/jnewland/starred{/owner}{/repo}",
          "subscriptions_url": "https://api.github.com/users/jnewland/subscriptions",
          "organizations_url": "https://api.github.com/users/jnewland/orgs",
          "repos_url": "https://api.github.com/users/jnewland/repos",
          "events_url": "https://api.github.com/users/jnewland/events{/privacy}",
          "received_events_url": "https://api.github.com/users/jnewland/received_events",
          "type": "User",
          "site_admin": false
        },
        "content_type": "application/macbinary",
        "state": "uploaded",
        "size": 545259520,
        "download_count": 24,
        "created_at": "2014-06-12T15:23:31Z",
        "updated_at": "2014-06-12T15:25:44Z",
        "browser_download_url": "https://github.com/github/pong/releases/download/large.bin/large.bin"
      }
    ],
    "tarball_url": "https://api.github.com/repos/github/pong/tarball/large.bin",
    "zipball_url": "https://api.github.com/repos/github/pong/zipball/large.bin",
    "body": "For download speed testing.\n"
  }
]

`))
	releases, err := GithubApi.GetNumberOfReleases("testOrg", "testRepo", testConfig)

	if err != nil {
		t.Errorf("expected no error response")
	}

	if len(releases) == 0 {
		t.Errorf("expected more than 0 repositories")
	}
}

func TestWhenCommitsGetIsSuccessful(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://api.github.com/repos/testOrg/testRepo/stats/contributors",
		httpmock.NewStringResponder(200, `[
  {
    "total": 1,
    "weeks": [
      {
        "w": 1297555200,
        "a": 3,
        "d": 0,
        "c": 1
      },
      {
        "w": 1298160000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1298764800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1299369600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1299974400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1300579200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1301184000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1301788800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1302393600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1302998400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1303603200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1304208000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1304812800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1305417600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1306022400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1306627200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1307232000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1307836800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1308441600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1309046400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1309651200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1310256000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1310860800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1311465600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1312070400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1312675200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1313280000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1313884800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1314489600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1315094400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1315699200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1316304000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1316908800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1317513600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1318118400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1318723200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1319328000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1319932800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1320537600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1321142400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1321747200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1322352000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1322956800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1323561600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1324166400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1324771200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1325376000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1325980800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1326585600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1327190400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1327795200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1328400000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1329004800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1329609600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1330214400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1330819200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1331424000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1332028800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1332633600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1333238400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1333843200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1334448000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1335052800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1335657600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1336262400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1336867200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1337472000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1338076800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1338681600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1339286400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1339891200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1340496000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1341100800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1341705600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1342310400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1342915200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1343520000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1344124800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1344729600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1345334400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1345939200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1346544000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1347148800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1347753600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1348358400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1348963200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1349568000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1350172800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1350777600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1351382400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1351987200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1352592000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1353196800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1353801600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1354406400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1355011200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1355616000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1356220800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1356825600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1357430400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1358035200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1358640000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1359244800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1359849600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1360454400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1361059200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1361664000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1362268800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1362873600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1363478400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1364083200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1364688000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1365292800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1365897600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1366502400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1367107200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1367712000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1368316800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1368921600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1369526400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1370131200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1370736000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1371340800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1371945600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1372550400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1373155200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1373760000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1374364800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1374969600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1375574400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1376179200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1376784000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1377388800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1377993600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1378598400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1379203200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1379808000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1380412800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1381017600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1381622400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1382227200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1382832000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1383436800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1384041600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1384646400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1385251200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1385856000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1386460800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1387065600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1387670400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1388275200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1388880000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1389484800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1390089600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1390694400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1391299200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1391904000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1392508800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1393113600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1393718400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1394323200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1394928000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1395532800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1396137600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1396742400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1397347200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1397952000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1398556800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1399161600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1399766400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1400371200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1400976000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1401580800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1402185600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1402790400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1403395200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1404000000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1404604800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1405209600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1405814400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1406419200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1407024000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1407628800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1408233600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1408838400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1409443200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1410048000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1410652800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1411257600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1411862400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1412467200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1413072000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1413676800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1414281600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1414886400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1415491200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1416096000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1416700800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1417305600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1417910400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1418515200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1419120000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1419724800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1420329600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1420934400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1421539200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1422144000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1422748800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1423353600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1423958400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1424563200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1425168000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1425772800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1426377600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1426982400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1427587200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1428192000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1428796800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1429401600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1430006400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1430611200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1431216000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1431820800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1432425600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1433030400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1433635200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1434240000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1434844800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1435449600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1436054400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1436659200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1437264000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1437868800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1438473600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1439078400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1439683200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1440288000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1440892800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1441497600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1442102400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1442707200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1443312000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1443916800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1444521600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1445126400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1445731200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1446336000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1446940800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1447545600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1448150400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1448755200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1449360000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1449964800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1450569600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1451174400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1451779200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1452384000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1452988800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1453593600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1454198400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1454803200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1455408000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1456012800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1456617600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1457222400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1457827200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1458432000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1459036800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1459641600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1460246400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1460851200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1461456000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1462060800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1462665600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1463270400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1463875200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1464480000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1465084800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1465689600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1466294400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1466899200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1467504000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1468108800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1468713600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1469318400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1469923200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1470528000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1471132800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1471737600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1472342400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1472947200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1473552000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1474156800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1474761600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1475366400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1475971200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1476576000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1477180800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1477785600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1478390400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1478995200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1479600000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1480204800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1480809600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1481414400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1482019200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1482624000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1483228800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1483833600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1484438400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1485043200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1485648000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1486252800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1486857600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1487462400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1488067200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1488672000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1489276800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1489881600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1490486400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1491091200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1491696000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1492300800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1492905600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1493510400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1494115200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1494720000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1495324800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1495929600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1496534400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1497139200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1497744000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1498348800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1498953600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1499558400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1500163200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1500768000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1501372800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1501977600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1502582400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1503187200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1503792000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1504396800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1505001600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1505606400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1506211200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1506816000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1507420800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1508025600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1508630400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1509235200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1509840000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1510444800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1511049600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1511654400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1512259200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1512864000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1513468800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1514073600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1514678400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1515283200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1515888000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1516492800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1517097600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1517702400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1518307200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1518912000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1519516800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1520121600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1520726400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1521331200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1521936000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1522540800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1523145600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1523750400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1524355200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1524960000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1525564800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1526169600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1526774400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1527379200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1527984000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1528588800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1529193600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1529798400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1530403200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1531008000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1531612800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1532217600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1532822400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1533427200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1534032000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1534636800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1535241600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1535846400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1536451200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1537056000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1537660800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1538265600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1538870400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1539475200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1540080000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1540684800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1541289600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1541894400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1542499200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1543104000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1543708800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1544313600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1544918400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1545523200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1546128000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1546732800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1547337600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1547942400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1548547200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1549152000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1549756800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1550361600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1550966400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1551571200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1552176000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1552780800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1553385600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1553990400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1554595200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1555200000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1555804800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1556409600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1557014400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1557619200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1558224000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1558828800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1559433600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1560038400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1560643200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1561248000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1561852800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1562457600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1563062400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1563667200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1564272000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1564876800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1565481600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1566086400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1566691200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1567296000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1567900800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1568505600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1569110400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1569715200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1570320000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1570924800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1571529600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1572134400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1572739200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1573344000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1573948800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1574553600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1575158400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1575763200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1576368000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1576972800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1577577600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1578182400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1578787200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1579392000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1579996800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1580601600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1581206400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1581811200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1582416000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1583020800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1583625600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1584230400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1584835200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1585440000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1586044800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1586649600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1587254400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1587859200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1588464000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1589068800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1589673600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1590278400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1590883200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1591488000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1592092800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1592697600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1593302400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1593907200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1594512000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1595116800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1595721600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1596326400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1596931200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1597536000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1598140800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1598745600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1599350400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1599955200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1600560000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1601164800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1601769600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1602374400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1602979200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1603584000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1604188800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1604793600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1605398400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1606003200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1606608000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1607212800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1607817600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1608422400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1609027200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1609632000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1610236800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1610841600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1611446400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1612051200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1612656000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1613260800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1613865600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1614470400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1615075200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1615680000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1616284800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1616889600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1617494400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1618099200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1618704000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1619308800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1619913600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1620518400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1621123200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1621728000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1622332800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1622937600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1623542400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1624147200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1624752000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1625356800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1625961600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1626566400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1627171200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1627776000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1628380800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1628985600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1629590400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1630195200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1630800000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1631404800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1632009600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1632614400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1633219200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1633824000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1634428800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1635033600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1635638400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1636243200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1636848000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1637452800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1638057600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1638662400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1639267200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1639872000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1640476800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1641081600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1641686400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1642291200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1642896000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1643500800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1644105600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1644710400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1645315200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1645920000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1646524800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1647129600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1647734400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1648339200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1648944000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1649548800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1650153600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1650758400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1651363200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1651968000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1652572800,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1653177600,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1653782400,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1654387200,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1654992000,
        "a": 0,
        "d": 0,
        "c": 0
      },
      {
        "w": 1655596800,
        "a": 0,
        "d": 0,
        "c": 0
      }
    ],
    "author": {
      "login": "rtomayko",
      "id": 404,
      "node_id": "MDQ6VXNlcjQwNA==",
      "avatar_url": "https://avatars.githubusercontent.com/u/404?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/rtomayko",
      "html_url": "https://github.com/rtomayko",
      "followers_url": "https://api.github.com/users/rtomayko/followers",
      "following_url": "https://api.github.com/users/rtomayko/following{/other_user}",
      "gists_url": "https://api.github.com/users/rtomayko/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/rtomayko/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/rtomayko/subscriptions",
      "organizations_url": "https://api.github.com/users/rtomayko/orgs",
      "repos_url": "https://api.github.com/users/rtomayko/repos",
      "events_url": "https://api.github.com/users/rtomayko/events{/privacy}",
      "received_events_url": "https://api.github.com/users/rtomayko/received_events",
      "type": "User",
      "site_admin": false
    }
  }
]
`))
	totalCommits, topContributer, err := GithubApi.GetTotalCommits("testOrg", "testRepo", testConfig)

	if err != nil {
		t.Errorf("expected no error response")
	}

	if totalCommits != 1 {
		t.Errorf("expected 1 commit not " + strconv.Itoa(totalCommits))
	}

	if topContributer != "rtomayko" {
		t.Errorf("expected rtomayko as top contributer not " + topContributer)
	}
}
