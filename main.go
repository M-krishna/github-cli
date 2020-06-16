package main

import (
	"flag"
	"fmt"
	"github-cli/cli"
	"github-cli/data"
	"os"
	"strings"
	"sync"
)

func init() {
	const (
		usage = "Search Users (Eg. john) or (Eg. john,foo,bar)"
	)
	flag.StringVar(&data.Users, "u", "", usage)
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	userResult := strings.Split(data.Users, ",")
	fmt.Printf("Searching user(s): %s\n", userResult)

	var wg sync.WaitGroup
	for _, username := range userResult {
		if len(username) > 1 {
			wg.Add(1)
			go cli.GetUserFromGithub(username, &wg)
		}
	}
	wg.Wait()
}
