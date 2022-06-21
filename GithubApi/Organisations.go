package GithubApi

import (
	"GithubApiCodeTest/Main/v2/Structs"
	"GithubApiCodeTest/Main/v2/config"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetOrganisation(oragnisationString string, configurations config.Configurations) (Structs.Organisation, error) {
	url := "https://api.github.com/orgs/" + oragnisationString
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return Structs.Organisation{}, err
	}

	if configurations.UseConfigCredentials {
		req.Header.Add("Authorization", "Bearer "+configurations.OAuthToken)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Structs.Organisation{}, err
	}
	defer res.Body.Close()

	organisation := Structs.Organisation{}
	decodeErr := json.NewDecoder(res.Body).Decode(&organisation)
	if decodeErr != nil {
		return Structs.Organisation{}, decodeErr
	}

	return organisation, nil
}
