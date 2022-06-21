package Structs

import "time"

type Repository struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	HTMLURL          string        `json:"html_url"`
	Description      string        `json:"description"`
	Fork             bool          `json:"fork"`
	URL              string        `json:"url"`
	ForksURL         string        `json:"forks_url"`
	KeysURL          string        `json:"keys_url"`
	CollaboratorsURL string        `json:"collaborators_url"`
	TeamsURL         string        `json:"teams_url"`
	HooksURL         string        `json:"hooks_url"`
	IssueEventsURL   string        `json:"issue_events_url"`
	EventsURL        string        `json:"events_url"`
	AssigneesURL     string        `json:"assignees_url"`
	BranchesURL      string        `json:"branches_url"`
	TagsURL          string        `json:"tags_url"`
	BlobsURL         string        `json:"blobs_url"`
	GitTagsURL       string        `json:"git_tags_url"`
	GitRefsURL       string        `json:"git_refs_url"`
	TreesURL         string        `json:"trees_url"`
	StatusesURL      string        `json:"statuses_url"`
	LanguagesURL     string        `json:"languages_url"`
	StargazersURL    string        `json:"stargazers_url"`
	ContributorsURL  string        `json:"contributors_url"`
	SubscribersURL   string        `json:"subscribers_url"`
	SubscriptionURL  string        `json:"subscription_url"`
	CommitsURL       string        `json:"commits_url"`
	GitCommitsURL    string        `json:"git_commits_url"`
	CommentsURL      string        `json:"comments_url"`
	IssueCommentURL  string        `json:"issue_comment_url"`
	ContentsURL      string        `json:"contents_url"`
	CompareURL       string        `json:"compare_url"`
	MergesURL        string        `json:"merges_url"`
	ArchiveURL       string        `json:"archive_url"`
	DownloadsURL     string        `json:"downloads_url"`
	IssuesURL        string        `json:"issues_url"`
	PullsURL         string        `json:"pulls_url"`
	MilestonesURL    string        `json:"milestones_url"`
	NotificationsURL string        `json:"notifications_url"`
	LabelsURL        string        `json:"labels_url"`
	ReleasesURL      string        `json:"releases_url"`
	DeploymentsURL   string        `json:"deployments_url"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
	PushedAt         time.Time     `json:"pushed_at"`
	GitURL           string        `json:"git_url"`
	SSHURL           string        `json:"ssh_url"`
	CloneURL         string        `json:"clone_url"`
	SvnURL           string        `json:"svn_url"`
	Homepage         string        `json:"homepage"`
	Size             int           `json:"size"`
	StargazersCount  int           `json:"stargazers_count"`
	WatchersCount    int           `json:"watchers_count"`
	Language         interface{}   `json:"language"`
	HasIssues        bool          `json:"has_issues"`
	HasProjects      bool          `json:"has_projects"`
	HasDownloads     bool          `json:"has_downloads"`
	HasWiki          bool          `json:"has_wiki"`
	HasPages         bool          `json:"has_pages"`
	ForksCount       int           `json:"forks_count"`
	MirrorURL        interface{}   `json:"mirror_url"`
	Archived         bool          `json:"archived"`
	Disabled         bool          `json:"disabled"`
	OpenIssuesCount  int           `json:"open_issues_count"`
	License          interface{}   `json:"license"`
	AllowForking     bool          `json:"allow_forking"`
	IsTemplate       bool          `json:"is_template"`
	Topics           []interface{} `json:"topics"`
	Visibility       string        `json:"visibility"`
	Forks            int           `json:"forks"`
	OpenIssues       int           `json:"open_issues"`
	Watchers         int           `json:"watchers"`
	DefaultBranch    string        `json:"default_branch"`
	Permissions      struct {
		Admin    bool `json:"admin"`
		Maintain bool `json:"maintain"`
		Push     bool `json:"push"`
		Triage   bool `json:"triage"`
		Pull     bool `json:"pull"`
	} `json:"permissions"`
}

type Contributer struct {
	Total int `json:"total"`
	Weeks []struct {
		Weeks int `json:"w"`
		A     int `json:"a"`
		D     int `json:"d"`
		C     int `json:"c"`
	} `json:"weeks"`
	Author struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
}

type Releases struct {
	URL       string `json:"url"`
	AssetsURL string `json:"assets_url"`
	UploadURL string `json:"upload_url"`
	HTMLURL   string `json:"html_url"`
	ID        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeID          string        `json:"node_id"`
	TagName         string        `json:"tag_name"`
	TargetCommitish string        `json:"target_commitish"`
	Name            string        `json:"name"`
	Draft           bool          `json:"draft"`
	Prerelease      bool          `json:"prerelease"`
	CreatedAt       time.Time     `json:"created_at"`
	PublishedAt     time.Time     `json:"published_at"`
	Assets          []interface{} `json:"assets"`
	TarballURL      string        `json:"tarball_url"`
	ZipballURL      string        `json:"zipball_url"`
	Body            string        `json:"body"`
}

type Commit struct {
	Sha    string `json:"sha"`
	NodeID string `json:"node_id"`
	Commit struct {
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    struct {
			Sha string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
		URL          string `json:"url"`
		CommentCount int    `json:"comment_count"`
		Verification struct {
			Verified  bool   `json:"verified"`
			Reason    string `json:"reason"`
			Signature string `json:"signature"`
			Payload   string `json:"payload"`
		} `json:"verification"`
	} `json:"commit"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	CommentsURL string `json:"comments_url"`
	Author      struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	Committer struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"committer"`
	Parents []struct {
		Sha     string `json:"sha"`
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
	} `json:"parents"`
}
