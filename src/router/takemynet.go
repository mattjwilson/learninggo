package router

import (
	"fmt"
	"net/http"
	"os"
	"bufio"
	"strings"
)

func main() {
	// toto
	fmt.Println("Hit")
}

func getParameters() (string, string, string) {
	if _, err := os.Stat("settings.mjw"); err == nil {
		// file exists
		var defaultSettings string
		var username string
		var password string
		var url string

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Would you like to use the discovered default settings?")
		fmt.Scan(&defaultSettings)
		//text, _ := reader.ReadString('\n')
		//text = strings.Replace(text, "\n", "", -1)

		fmt.Println(text)
		if defaultSettings == "y" {
			file, err := os.OpenFile("settings.mjw", os.O_RDWR|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("welp...something went sideways")
				fmt.Println(err)
				return
			}
			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
		}
	} else {
		// file does not exist
		file, err := os.Create("settings.mjw")
		defer file.Close()
		if(err != nil) {
			fmt.Println("Something went sideways")
		}

		fmt.Fprintln(file, url)
		fmt.Fprintln(file, username)
		fmt.Println(file, password)
	}
	return username, password, url
}

func getCurrentPassword(sourceUrl string) {
	resp, err := http.Get(sourceUrl)
	if(err != nil) {
		fmt.Println("Something went wrong attempting to get the page.")
	}
}

// Creates the header set required in order to pull information from my router
func generateRequestHeader() string[] {
	var returnVal string[]

	return returnVal
}
