package GithubApi

import (
	"GithubApiCodeTest/Main/v2/Structs"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRepositories(organisationString string) ([]Structs.Repository, error) {
	url := "https://api.github.com/orgs/" + organisationString + "/repos"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return []Structs.Repository{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []Structs.Repository{}, err
	}
	defer res.Body.Close()

	repositories := []Structs.Repository{}
	decodeErr := json.NewDecoder(res.Body).Decode(&repositories)
	if decodeErr != nil {
		return []Structs.Repository{}, decodeErr
	}

	return repositories, nil
}

func GetTotalCommits(organisationString string, repositoryName string) (int, error) {
	url := "https://api.github.com/repos/" + organisationString + "/" + repositoryName + "/stats/contributors"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer res.Body.Close()

	repositories := []Structs.Commits{}
	decodeErr := json.NewDecoder(res.Body).Decode(&repositories)
	if decodeErr != nil {
		return 0, decodeErr
	}

	return len(repositories), nil
}

func GetNumberOfReleases(organisationString string, repositoryName string) ([]Structs.Commits, error) {
	url := "https://api.github.com/repos/" + organisationString + "/" + repositoryName + "/releases"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return []Structs.Commits{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []Structs.Commits{}, err
	}
	defer res.Body.Close()

	repositories := []Structs.Commits{}
	decodeErr := json.NewDecoder(res.Body).Decode(&repositories)
	if decodeErr != nil {
		return []Structs.Commits{}, decodeErr
	}

	return repositories, nil
}
