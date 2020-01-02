package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

var options struct {
	users string
}

type Response struct {
	Login string 			`json:"login"`
	Avatar string 			`json:"avatar_url"`
	URL string				`json:"url"`
	Name string				`json:"name"`
	Company string			`json:"company"`
	Location string			`json:"location"`
	Bio string				`json:"bio"`
	Public_Repositories int	`json:"public_repos"`
	Followers int			`json:"followers"`
	Following int			`json:"following"`
}


func init(){
	
	const (
		usage = "Search Users (Eg. john,david)"
	)

	flag.StringVar(&options.users, "users", "", usage)
	flag.StringVar(&options.users, "u", "", "Search User (Eg. john)")
}

func GetUserFromGithub(name string){
	
	var responseBody Response

	resp, err := http.Get("https://api.github.com/users/" + name)
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
	
		fmt.Println()
		fmt.Println("Name: ", responseBody.Name)
		fmt.Println("Login: ", responseBody.Login)
		fmt.Println("Avatar: ", responseBody.Avatar)
		fmt.Println("URL: ", responseBody.URL)
		fmt.Println("Company: ", responseBody.Company)
		fmt.Println("Location: ", responseBody.Location)
		fmt.Println("Bio: ", responseBody.Bio)
		fmt.Println("Followers: ", responseBody.Followers)
		fmt.Println("Following: ", responseBody.Following)
		fmt.Println()
	} else {
		fmt.Printf("%s not found\n", name)
	}
}

func main()  {

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
     	fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var users *string
	users = &options.users
	user := *users
	user_result := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", user_result)

	for _, username := range user_result {
		if len(username) > 1 {
			go GetUserFromGithub(username)
		}
	}
}