package data

var Users string

type Response struct {
	Login              string `json:"login"`
	Avatar             string `json:"avatar_url"`
	URL                string `json:"url"`
	Name               string `json:"name"`
	Company            string `json:"company"`
	Location           string `json:"location"`
	Bio                string `json:"bio"`
	PublicRepositories int    `json:"public_repos"`
	Followers          int    `json:"followers"`
	Following          int    `json:"following"`
}
