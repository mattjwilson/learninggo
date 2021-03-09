package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	username, password, url := getParameters()
	getCurrentPassword(url)
	generateRequestHeader(username, password)
}

func getParameters() (string, string, string) {
	var username string
	var password string
	var url string

	if _, err := os.Stat("settings.mjw"); err == nil {
		// file exists
		file, err := os.Open("settings.mjw")
		if err != nil {
			fmt.Println("welp...something went sideways")
			fmt.Println(err)
			return "", "", ""
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		count := 0
		for scanner.Scan() {
			switch count {
			case 0:
				url = scanner.Text()
			case 1:
				username = scanner.Text()
			case 2:
				password = scanner.Text()
			default:
				fmt.Println("Unexpected entry in our file: ", scanner.Text())
			}
			count++
		}
	} else {
		// file does not exist
		file, err := os.Create("settings.mjw")
		defer file.Close()
		if err != nil {
			fmt.Println("Something went sideways")
		}

		fmt.Print("Username: ")
		fmt.Scan(&username)

		fmt.Print("Password: ")
		fmt.Scan(&password)

		fmt.Print("URL: ")
		fmt.Scan(&url)

		fmt.Fprintln(file, url)
		fmt.Fprintln(file, username)
		fmt.Fprintln(file, password)
	}
	fmt.Println(username, password, url)
	return username, password, url
}

func getCurrentPassword(sourceUrl string) {
	resp, err := http.Get(sourceUrl)
	if err != nil {
		fmt.Println("Something went wrong attempting to get the page.")
	}

	fmt.Println(resp)
}

// Creates the header set required in order to pull information from my router
func generateRequestHeader(username string, password string) []string {
	var returnVal []string

	return returnVal
}
