package main

import (
	"GithubApiCodeTest/Main/v2/GithubApi"
	"GithubApiCodeTest/Main/v2/Structs"
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("GithubApiCodeTest")

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
				break
			}

			results := []Structs.GithubApiResults{}
			for _, repository := range respositories {
				totalCommits, _ := GithubApi.GetTotalCommits(anOrganisation.Name, repository.Name)
				releases, _ := GithubApi.GetNumberOfReleases(anOrganisation.Name, repository.Name)
				aResult := Structs.GithubApiResults{
					RepoName:         repository.Name,
					NumberOfCommits:  totalCommits,
					NumberOfReleases: len(releases),
				}
				results = append(results, aResult)
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
