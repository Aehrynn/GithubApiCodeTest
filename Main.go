package main

import (
	"GithubApiCodeTest/Main/v2/GithubApi"
	"GithubApiCodeTest/Main/v2/Structs"
	"GithubApiCodeTest/Main/v2/config"
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("GithubApiCodeTest")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	for {
		fmt.Print("Enter Github Organisation-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		anOrganisation, getOrganisationErr := GithubApi.GetOrganisation(text)
		if getOrganisationErr == nil {
			//fmt.Println(anOrganisation)
			respositories, err := GithubApi.GetRepositories(anOrganisation.Name)

			if err != nil {
				fmt.Println("Error getting repos")
				fmt.Println(err)
				break
			}

			results := []Structs.GithubApiResults{}
			for key, repository := range respositories {
				totalCommits, topContributer, _ := GithubApi.GetTotalCommits(anOrganisation.Name, repository.Name)
				releases, _ := GithubApi.GetNumberOfReleases(anOrganisation.Name, repository.Name)
				aResult := Structs.GithubApiResults{
					RepoName:         repository.Name,
					NumberOfCommits:  totalCommits,
					NumberOfReleases: len(releases),
					TopContributer:   topContributer,
				}
				results = append(results, aResult)
				if key >= configuration.MaxRepoSize {
					break
				}
			}

			if len(results) == 0 {
				fmt.Println("No Repositories found for that Organisation")
				break
			}

			headers := []string{
				"RepoName",
				"NumberOfReleases",
				"NumberOfCommits",
				"TopContributer",
			}
			buffer := new(bytes.Buffer)
			writer := csv.NewWriter(buffer)
			writer.Write(headers)

			for _, aResult := range results {

				row := []string{aResult.RepoName, strconv.Itoa(aResult.NumberOfReleases), strconv.Itoa(aResult.NumberOfCommits), aResult.TopContributer}

				writer.Write(row)
			}
			writer.Flush()

			csvString := buffer.String()
			fmt.Println(csvString)

			fileOutput, err := os.Create(anOrganisation.Name + "_Results.csv")

			fileOutput.WriteString(csvString)

			fileOutput.Sync()
			break
		}

		fmt.Print("Error getting organisation")
	}
}
