## GithubAPI Code Test

Get dependencies

`go get`

To build

`go build`

To Run (Linux)
`./Main`

Or `Main.exe` on Windows

To run unit tests and coverage

`go test ./... --cover -coverprofile=coverage.out -coverpkg=./... && go tool cover -html=coverage.out -o=coverage.html`

## Potential Enhancements

* Make number of total commits for all repository and not just origin / main branch
  * There is a discrepancy shown when seeing the total commits on the web ui versus the API
* Rate Control Limits
* Option for the kind of format of the output file
* Unit testing
* Handle pagination of Requests


## Usage Instructions

* Run the Program, you will enter the organizationame.
* Results will show on screen and a csv output will be written to disk, with format <OrganizationName>_Results<timstamp>.csv
* Default using non authentication will be limited by the github api per hour 

### To Run with Personal Access token

* Create a github account
* Create a Personal access token, the default settings are enough for using this program
  * https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token
* Update the config.yml file
  * replace the `<PAT_TOKEN>` at `OAuthToken` with the token, it will start with `ghp` 
  * change the value for `UseConfigCredentials` to `true`
* Rerun the program 

