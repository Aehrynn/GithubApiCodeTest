package Structs

type GithubApiResults struct {
	RepoName         string `json:"repo_name"`
	NumberOfReleases int    `json:"number_of_releases"`
	NumberOfCommits  int    `json:"number_of_commits"`
	TopContributer   string `json:"top_contributer"`
}

type GithubError struct {
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}
