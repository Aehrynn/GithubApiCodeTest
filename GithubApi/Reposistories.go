package GithubApi

import (
	"GithubApiCodeTest/Main/v2/Structs"
	"GithubApiCodeTest/Main/v2/config"
	"GithubApiCodeTest/Main/v2/constants"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRepositories(organisationString string, configurations config.Configurations) ([]Structs.Repository, error) {
	url := "https://api.github.com/orgs/" + organisationString + "/repos"
	method := "GET"

	client := &http.Client{}
	req, err := constants.NewHttpRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return []Structs.Repository{}, err
	}

	if configurations.UseConfigCredentials {
		req.Header.Add("Authorization", "Bearer "+configurations.OAuthToken)
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

func GetTotalCommits(organisationString string, repositoryName string, configurations config.Configurations) (int, string, error) {
	//url := "https://api.github.com/repos/" + organisationString + "/" + repositoryName + "/commits?per_page=500"
	url := "https://api.github.com/repos/" + organisationString + "/" + repositoryName + "/stats/contributors"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0, "", err
	}
	if configurations.UseConfigCredentials {
		req.Header.Add("Authorization", "Bearer "+configurations.OAuthToken)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, "", err
	}
	defer res.Body.Close()

	//repositories := []Structs.Commit{}
	repositories := []Structs.Contributer{}
	decodeErr := json.NewDecoder(res.Body).Decode(&repositories)
	if decodeErr != nil {
		return 0, "", decodeErr
	}

	//topContributer, err := GetTopCommiter(organisationString, repositoryName, configurations)
	//if err != nil {
	//	fmt.Println(err)
	//	return 0, "", err
	//}

	topContributer := ""
	topContributerCommitNumber := 0
	totalCommits := 0

	for _, value := range repositories {
		totalCommits = totalCommits + value.Total
		if value.Total >= topContributerCommitNumber {
			topContributer = value.Author.Login
			topContributerCommitNumber = value.Total
		}
	}

	//return len(repositories), topContributer, nil
	return totalCommits, topContributer, nil
}

func GetTopCommiter(organisationString string, repositoryName string, configurations config.Configurations) (string, error) {
	url := "https://api.github.com/repos/" + organisationString + "/" + repositoryName + "/stats/contributors"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	if configurations.UseConfigCredentials {
		req.Header.Add("Authorization", "Bearer "+configurations.OAuthToken)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	repositories := []Structs.Contributer{}
	decodeErr := json.NewDecoder(res.Body).Decode(&repositories)
	if decodeErr != nil {
		return "", decodeErr
	}

	topContributer := ""
	topContributerCommitNumber := 0

	for _, value := range repositories {
		if value.Total >= topContributerCommitNumber {
			topContributer = value.Author.Login
			topContributerCommitNumber = value.Total
		}
	}

	return topContributer, nil
}

func GetNumberOfReleases(organisationString string, repositoryName string, configurations config.Configurations) ([]Structs.Releases, error) {
	url := "https://api.github.com/repos/" + organisationString + "/" + repositoryName + "/releases"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return []Structs.Releases{}, err
	}
	if configurations.UseConfigCredentials {
		req.Header.Add("Authorization", "Bearer "+configurations.OAuthToken)
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
