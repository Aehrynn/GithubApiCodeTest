package GithubApi

import (
	"GithubApiCodeTest/Main/v2/Structs"
	"GithubApiCodeTest/Main/v2/constants"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRepositories(organisationString string) ([]Structs.Repository, error) {
	url := "https://api.github.com/orgs/" + organisationString + "/repos"
	method := "GET"

	client := &http.Client{}
	req, err := constants.NewHttpRequest(method, url, nil)

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

	if res.StatusCode != http.StatusOK {
		errorMessage := Structs.GithubError{}
		decodeErr := json.NewDecoder(res.Body).Decode(&errorMessage)
		if decodeErr != nil {
			return []Structs.Repository{}, decodeErr
		}
		return []Structs.Repository{}, fmt.Errorf(errorMessage.Message)
	}

	repositories := []Structs.Repository{}
	decodeErr := json.NewDecoder(res.Body).Decode(&repositories)
	if decodeErr != nil {
		return []Structs.Repository{}, decodeErr
	}

	return repositories, nil
}

func GetTotalCommits(organisationString string, repositoryName string) (int, string, error) {
	url := "https://api.github.com/repos/" + organisationString + "/" + repositoryName + "/commits"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0, "", err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, "", err
	}
	defer res.Body.Close()

	repositories := []Structs.Commit{}
	decodeErr := json.NewDecoder(res.Body).Decode(&repositories)
	if decodeErr != nil {
		return 0, "", decodeErr
	}

	topContributer := ""
	//topContributerCommitNumber := 0
	commits := 0

	for _, _ = range repositories {
		commits = commits + 1
		//if value.Total >= topContributerCommitNumber {
		//	topContributer = value.Author.Login
		//	topContributerCommitNumber = value.Total
		//}
	}

	return commits, topContributer, nil
}

func GetNumberOfReleases(organisationString string, repositoryName string) ([]Structs.Releases, error) {
	url := "https://api.github.com/repos/" + organisationString + "/" + repositoryName + "/releases"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return []Structs.Releases{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []Structs.Releases{}, err
	}
	defer res.Body.Close()

	repositories := []Structs.Releases{}
	decodeErr := json.NewDecoder(res.Body).Decode(&repositories)
	if decodeErr != nil {
		return []Structs.Releases{}, decodeErr
	}

	return repositories, nil
}
