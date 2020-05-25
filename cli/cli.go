package cli

import (
	"encoding/json"
	"fmt"
	"github-cli/data"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

const (
	baseURL = "https://api.github.com/users/"
)

func GetUserFromGithub(name string, wg *sync.WaitGroup) {

	var responseBody data.Response

	resp, err := http.Get(baseURL + name)
	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		os.Exit(1)
	}

	if resp.StatusCode == 200 {
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Failed reading data from the response %s\n", err)
			os.Exit(1)
		}

		json.Unmarshal(responseData, &responseBody)

		fmt.Println("#######################################")
		fmt.Println("Name: ", responseBody.Name)
		fmt.Println("Login: ", responseBody.Login)
		fmt.Println("Avatar: ", responseBody.Avatar)
		fmt.Println("URL: ", responseBody.URL)
		fmt.Println("Company: ", responseBody.Company)
		fmt.Println("Location: ", responseBody.Location)
		fmt.Println("Bio: ", responseBody.Bio)
		fmt.Println("Followers: ", responseBody.Followers)
		fmt.Println("Following: ", responseBody.Following)
		fmt.Println("########################################")
	} else {
		fmt.Printf("%s not found\n", name)
	}
	wg.Done()
}
